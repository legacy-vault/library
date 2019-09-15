package DoubleBuffer

// A simple Double Buffer.
//
// It has two Accumulators. While one of them is receiving Items, an other
// Accumulator is being flushed. This Buffer may be used only by a single User,
// i.e. it does not support multi-threading.

import (
	"errors"
	"sync"
	"time"

	"github.com/legacy-vault/library/go/DoubleBuffer/Accumulator"
	"github.com/legacy-vault/library/go/DoubleBuffer/FlushTask"
	"github.com/legacy-vault/library/go/DoubleBuffer/Flusher"
	"github.com/legacy-vault/library/go/DoubleBuffer/Item"
)

const (
	MsgNotWorking             = "buffer is not working"
	MsgNotReceivingItems      = "buffer is not receiving items"
	MsgBadAccumulatorSelector = "accumulator selector is broken"
)

const (
	FlushTasksChannelSize  = 2
	FlushErrorsChannelSize = 1024
)

type DoubleBuffer struct {

	// Main Settings.
	capacity  int
	isWorking bool

	// Accumulator.
	accumulatorA       Accumulator.Accumulator
	accumulatorB       Accumulator.Accumulator
	currentAccumulator uint8

	// Flusher.
	flusherFunction Flusher.Flusher
	flushTasks      chan FlushTask.FlushTask
	flushErrors     chan error
	flusherWG       sync.WaitGroup

	// Receiver.
	isReceivingItems bool
	receiverWG       sync.WaitGroup
}

var (
	ErrNotWorking             = errors.New(MsgNotWorking)
	ErrNotReceivingItems      = errors.New(MsgNotReceivingItems)
	ErrBadAccumulatorSelector = errors.New(MsgBadAccumulatorSelector)
)

func NewDoubleBuffer(
	capacity int,
	flusherFunction Flusher.Flusher,
) (*DoubleBuffer, error) {

	var db *DoubleBuffer

	db = new(DoubleBuffer)
	db.initialize(capacity, flusherFunction)
	db.startWorking()
	db.startReceivingItems()

	return db, nil
}

func (this *DoubleBuffer) initialize(
	capacity int,
	flusherFunction func([]Item.Item,
	) error,
) {

	this.capacity = capacity

	this.accumulatorA.Initialize(capacity, flusherFunction)
	this.accumulatorB.Initialize(capacity, flusherFunction)
	this.currentAccumulator = Accumulator.AccumulatorA

	this.flusherFunction = flusherFunction
	this.flushTasks = make(chan FlushTask.FlushTask, FlushTasksChannelSize)
	this.flushErrors = make(chan error, FlushErrorsChannelSize)
	this.flusherWG.Add(1)
	go this.flushManager()
}

func (this DoubleBuffer) GetFlushErrorsChannel() chan error {
	return this.flushErrors
}

func (this *DoubleBuffer) startWorking() {
	this.isWorking = true
}

func (this *DoubleBuffer) startReceivingItems() {
	this.isReceivingItems = true
}

func (this *DoubleBuffer) flushManager() {

	var currentAccumulator *Accumulator.Accumulator
	var err error

	for flushTask := range this.flushTasks {

		switch flushTask.FlushedAccumulator {

		case Accumulator.AccumulatorA:
			currentAccumulator = &this.accumulatorA

		case Accumulator.AccumulatorB:
			currentAccumulator = &this.accumulatorB

		default:
			// This may be a hardware Fault or even a Virus Attack.
			this.flushErrors <- ErrBadAccumulatorSelector
			continue
		}

		// Do the Flushing.
		_, err = currentAccumulator.Flush()
		if err != nil {
			this.flushErrors <- err
		}

		err = currentAccumulator.StartReceivingItems()
		if err != nil {
			this.flushErrors <- err
		}
	}

	close(this.flushErrors)
	this.flusherWG.Done()
}

func (this *DoubleBuffer) receiveItemStrict(
	item Item.Item,
) error {

	var currentAccumulator *Accumulator.Accumulator
	var currentAccumulatorState uint8
	var err error
	var otherAccumulatorIdx uint8

	if !this.IsReceivingItems() {
		return ErrNotReceivingItems
	}

	// Control Flag for graceful Shutdown.
	this.receiverWG.Add(1)
	defer this.receiverWG.Done()

	switch this.currentAccumulator {

	case Accumulator.AccumulatorA:
		currentAccumulator = &this.accumulatorA
		otherAccumulatorIdx = Accumulator.AccumulatorB

	case Accumulator.AccumulatorB:
		currentAccumulator = &this.accumulatorB
		otherAccumulatorIdx = Accumulator.AccumulatorA

	default:
		// This may be a hardware Fault or even a Virus Attack.
		return ErrBadAccumulatorSelector
	}

	currentAccumulatorState, err = currentAccumulator.ReceiveItem(item)
	if err != nil {
		return err
	}
	if currentAccumulatorState == Accumulator.StateFull {

		// Send Flush Task.
		this.flushTasks <- FlushTask.FlushTask{
			FlushedAccumulator: this.currentAccumulator,
		}

		// Switch Accumulators.
		this.currentAccumulator = otherAccumulatorIdx
	}

	return nil
}

func (this *DoubleBuffer) ReceiveItem(
	item Item.Item,
) error {

	var err error

	err = this.receiveItemStrict(item)
	for err != nil {
		if err != Accumulator.ErrIsNotReceiving {
			return err
		}
		// Make a Switch between Routines to give Control to Flusher.
		time.Sleep(time.Microsecond)
		// Try again.
		err = this.receiveItemStrict(item)

	}

	return nil
}

func (this DoubleBuffer) IsReceivingItems() bool {
	return this.isReceivingItems
}

func (this *DoubleBuffer) Shutdown() error {

	var err error

	if !this.IsWorking() {
		return ErrNotWorking
	}

	// Stop receiving new Tasks and wait for the Receiver to finish.
	this.stopReceivingItems()
	this.receiverWG.Wait()

	// Flush all Accumulators with Force.
	err = this.forceFlushAllAccumulators()
	if err != nil {
		return err
	}

	// Wait for the Flusher to finish all Flush Tasks.
	close(this.flushTasks)
	this.flusherWG.Wait()

	this.stopWorking()

	return nil
}

func (this DoubleBuffer) IsWorking() bool {
	return this.isWorking
}

func (this *DoubleBuffer) stopReceivingItems() {
	this.isReceivingItems = false
}

func (this *DoubleBuffer) forceFlushAllAccumulators() error {

	var accumulators []Accumulator.Accumulator
	var accumulatorIndices []uint8

	accumulators = []Accumulator.Accumulator{
		this.accumulatorA,
		this.accumulatorB,
	}
	accumulatorIndices = []uint8{
		Accumulator.AccumulatorA,
		Accumulator.AccumulatorB,
	}

	for i, accumulator := range accumulators {
		// Check State.
		accState := accumulator.GetState()
		if (accState == Accumulator.StateFlushing) ||
			(accState == Accumulator.StateEmpty) {
			continue
		}
		// Send Flush Task.
		this.flushTasks <- FlushTask.FlushTask{
			FlushedAccumulator: accumulatorIndices[i],
		}
	}

	return nil
}

func (this *DoubleBuffer) stopWorking() {
	this.isWorking = false
}
