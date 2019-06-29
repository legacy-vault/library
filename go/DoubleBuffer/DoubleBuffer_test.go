//+build test

package DoubleBuffer

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/legacy-vault/library/go/DoubleBuffer/Flusher"
	"github.com/legacy-vault/library/go/DoubleBuffer/Item"
)

func Test_Simple(t *testing.T) {

	var db *DoubleBuffer
	var err error
	var flushedItems []interface{}
	var flusherFunc Flusher.Flusher
	var flushErrorChannel chan error

	// Create the Double Buffer Object.
	flushedItems = make([]interface{}, 0)
	flusherFunc = func(items []Item.Item) error {
		for _, item := range items {
			flushedItems = append(flushedItems, item)
		}
		return nil
	}
	db, err = NewDoubleBuffer(3, flusherFunc)
	if err != nil {
		t.FailNow()
	}
	flushErrorChannel = db.GetFlushErrorsChannel()

	err = db.ReceiveItem("First")
	if err != nil {
		t.FailNow()
	}
	err = db.ReceiveItem("Second")
	if err != nil {
		t.FailNow()
	}
	err = db.ReceiveItem("Third")
	if err != nil {
		t.FailNow()
	}
	err = db.ReceiveItem("4-th")
	if err != nil {
		t.FailNow()
	}
	err = db.ReceiveItem("5-th")
	if err != nil {
		t.FailNow()
	}
	err = db.ReceiveItem("6-th")
	if err != nil {
		t.FailNow()
	}

	// Wait for the Flusher to process its Tasks.
	time.Sleep(time.Millisecond)

	err = db.ReceiveItem("7-th")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	// Check the Errors Channel.
	go func() {
		for err := range flushErrorChannel {
			log.Println(err)
			t.Fail()
		}
	}()

	// Stop the Buffer.
	err = db.Shutdown()
	if err != nil {
		t.FailNow()
	}

	fmt.Println("flushedItems:", flushedItems)
	if len(flushedItems) != 7 {
		t.Log("Flushed Items Count")
		t.FailNow()
	}
	if (flushedItems[0] != "First") ||
		(flushedItems[1] != "Second") ||
		(flushedItems[2] != "Third") ||
		(flushedItems[3] != "4-th") ||
		(flushedItems[4] != "5-th") ||
		(flushedItems[5] != "6-th") ||
		(flushedItems[6] != "7-th") {
		t.Log("Flushed Items Values")
		t.FailNow()
	}
}

func Test_Stress_A(t *testing.T) {

	const BufferAccumulatorCapacity = 1000 * 1000

	var db *DoubleBuffer
	var err error
	var flushedItems []interface{}
	var flushedItem interface{}
	var flusherFunc Flusher.Flusher
	var flushErrorChannel chan error
	var i uint32
	var idx int
	var iMax uint32

	// Create the Double Buffer Object.
	flushedItems = make([]interface{}, 0)
	flusherFunc = func(items []Item.Item) error {
		for _, item := range items {
			flushedItems = append(flushedItems, item)
		}
		return nil
	}
	db, err = NewDoubleBuffer(BufferAccumulatorCapacity, flusherFunc)
	if err != nil {
		t.FailNow()
	}
	flushErrorChannel = db.GetFlushErrorsChannel()

	// Check the Errors Channel.
	go func() {
		for err := range flushErrorChannel {
			log.Println(err)
			t.Fail()
		}
	}()

	// Send Data.
	iMax = BufferAccumulatorCapacity * 50
	for i = 1; i <= iMax; i++ {
		err = db.ReceiveItem(i)
		if err != nil {
			t.Log("i=", i, err)
			t.FailNow()
		}
	}

	// Stop the Buffer.
	err = db.Shutdown()
	if err != nil {
		t.FailNow()
	}

	// Check the Results.
	//fmt.Println("flushedItems:", flushedItems)
	if len(flushedItems) != int(iMax) {
		t.Log("Flushed Items Count")
		t.FailNow()
	}
	for idx, flushedItem = range flushedItems {
		if flushedItem != uint32(idx+1) {
			log.Printf(
				"Flushed Items Value Mismatch: '%v' vs '%v'.",
				idx+1,
				flushedItem,
			)
			t.FailNow()
		}
	}
}
