package Accumulator

import (
	"errors"

	"github.com/legacy-vault/library/go/DoubleBuffer/Flusher"
	"github.com/legacy-vault/library/go/DoubleBuffer/Item"
)

const (
	MsgIsFull         = "Accumulator is full"
	MsgIsNotReceiving = "Accumulator is not receiving"
	MsgIsNotEmpty     = "Accumulator is not empty"
)

const (
	StateReceiving = 1
	StateFull      = 2
	StateFlushing  = 3
	StateEmpty     = 4
)

const (
	AccumulatorA = 1
	AccumulatorB = 2
)

type Accumulator struct {
	capacity        int
	items           []Item.Item
	state           uint8
	flusherFunction Flusher.Flusher
}

var (
	ErrIsFull         = errors.New(MsgIsFull)
	ErrIsNotReceiving = errors.New(MsgIsNotReceiving)
	ErrIsNotEmpty     = errors.New(MsgIsNotEmpty)
)

func (this *Accumulator) Initialize(
	capacity int,
	flusherFunction Flusher.Flusher,
) {

	this.capacity = capacity
	this.items = make([]Item.Item, 0, capacity)
	this.startReceivingItems()
	this.flusherFunction = flusherFunction
}

func (this *Accumulator) StartReceivingItems() error {

	if this.state != StateEmpty {
		return ErrIsNotEmpty
	}

	this.startReceivingItems()

	return nil
}

func (this *Accumulator) startReceivingItems() {
	this.state = StateReceiving
}

func (this Accumulator) GetState() uint8 {
	return this.state
}

func (this *Accumulator) ReceiveItem(
	item Item.Item,
) (uint8, error) {

	if this.state != StateReceiving {
		return this.state, ErrIsNotReceiving
	}
	if len(this.items) >= this.capacity {
		return this.state, ErrIsFull
	}

	this.items = append(this.items, item)
	if len(this.items) == this.capacity {
		this.state = StateFull
	}

	return this.state, nil
}

func (this *Accumulator) Flush() (uint8, error) {

	var err error
	var stateBeforeCall uint8

	stateBeforeCall = this.state

	this.state = StateFlushing
	err = this.flusherFunction(this.items)
	if err != nil {
		this.state = stateBeforeCall
		return this.state, err
	}

	this.items = make([]Item.Item, 0, this.capacity)
	this.state = StateEmpty

	return this.state, nil
}
