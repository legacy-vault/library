//============================================================================//
//
// Copyright © 2018 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-26.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// cache.go.

// Fixed Size Bubble Cache :: Cache Object Functions.

// Description of the Package is available in a separate 'ReadMe' File.

package fsbcache

import (
	"errors"
	"reflect"
	"time"
)

type Cache struct {
	head *Record
	tail *Record

	size     uint64
	capacity uint64

	recordByUID map[RecordUID]*Record
	recordTTL   int64 // Seconds
}

var (
	ErrDataIsEmpty = errors.New("'Data' Field is not set")
	ErrUIDIsEmpty  = errors.New("'UID' Field is not set")
)

// Creates a new List.
// * [Capacity] is the maximum Size of Cache. New Records added when the Cache
// is at its maximum Size, will remove the Tail Record. Tail Record is a Record
// with the oldest Access Time.
// * [RecordTTL] is the Period of Time, after which the Record is considered
// outdated. If the requested Record is outdated, it is removed from the Cache.
// Measured in Seconds.
func New(
	capacity uint64,
	recordTTL int64,
) *Cache {

	var cache *Cache

	if capacity == 0 {
		capacity++
	}

	cache = new(Cache)
	cache.initialize(capacity, recordTTL)

	return cache
}

// Checks the Record's Parameters and adds a Record to the Cache.
func (cache *Cache) AddRecord(uid RecordUID, data interface{}) error {

	var uidType reflect.Kind

	// Check 'Data' Field.
	if data == nil {
		return ErrDataIsEmpty
	}

	// Check 'UID' Field.
	uidType = reflect.TypeOf(uid).Kind()
	switch uidType {

	case reflect.String:
		if len(uid) == 0 {
			return ErrUIDIsEmpty
		}
	}

	cache.addARecord(&Record{UID: uid, Data: data})

	return nil
}

// Deletes all Records from the Cache.
// Returns 'true' on Success.
func (cache *Cache) Clear() bool {

	var i uint64
	var cacheIsIntegral bool
	var size uint64

	// Before deleting the Records, we must ensure that Cache is not broken.
	// Broken Cache Deletion would have caused us a lot of Memory Leaks!
	cacheIsIntegral = cache.isIntegral()
	if !cacheIsIntegral {
		return false
	}

	// Prepare Data.
	size = cache.size

	// Delete all Items from Tail to Head.
	for i = 1; i <= size; i++ {
		cache.deleteARecord(cache.tail)
	}

	return true
}

// Deletes from the Cache a Record specified by its UID if it exists in Cache.
func (cache *Cache) DeleteRecordByUID(uid RecordUID) {

	var record *Record
	var recordExists bool

	// Find the Record.
	record, recordExists = cache.recordByUID[uid]
	if !recordExists {
		return
	}

	// Delete the Record.
	cache.deleteARecord(record)

	return
}

// Enlists Values of all Records of the Cache.
func (cache *Cache) EnlistAllRecordValues() []interface{} {

	var i uint64
	var record *Record
	var size uint64
	var values []interface{}

	size = cache.size
	values = make([]interface{}, size)
	if size == 0 {
		return values
	}

	// Get the first Item.
	record = cache.head
	values[0] = record.Data

	// Get all other Items.
	for i = 1; i < size; i++ {
		record = record.nextItem
		values[i] = record.Data
	}

	return values
}

// Enlists all Records of the Cache.
func (cache *Cache) EnlistAllRecords() []*Record {

	var i uint64
	var record *Record
	var records []*Record
	var size uint64

	size = cache.size
	records = make([]*Record, size)
	if size == 0 {
		return records
	}

	// Get the first Item.
	record = cache.head
	records[0] = record

	// Get all other Items.
	for i = 1; i < size; i++ {
		record = record.nextItem
		records[i] = record
	}

	return records
}

// Gets Record's Data by its UID.
// Second returned Parameter is 'false' when one the following is true:
// 	-	The UID does not exist,
//	-	The Record exists but is outdated.
func (cache *Cache) GetRecordDataByUID(uid RecordUID) (interface{}, bool) {

	var data interface{}
	var nowTimeStamp int64
	var record *Record
	var recordLATMax int64
	var uidExists bool

	record, uidExists = cache.recordByUID[uid]
	if !uidExists {
		return nil, false
	}

	// Check the TTL. Is the Record Outdated?
	recordLATMax = record.lastAccessTime + cache.recordTTL
	nowTimeStamp = time.Now().Unix()
	if nowTimeStamp > recordLATMax {
		// Record is outdated and must be deleted.
		cache.deleteARecord(record)

		return nil, false
	}

	data = record.Data

	return data, true
}

// Returns the 'RecordTTL' Parameter of the Cache.
func (cache *Cache) GetRecordTTL() int64 {

	return cache.recordTTL
}

// Checks whether the specified Record's UID exists in the Cache.
func (cache *Cache) RecordUIDExists(uid RecordUID) bool {

	var uidExists bool

	_, uidExists = cache.recordByUID[uid]

	return uidExists
}

// Changes the 'RecordTTL' Parameter of the Cache.
func (cache *Cache) SetRecordTTL(recordTTL int64) {

	cache.recordTTL = recordTTL

	return
}

// Adds a Record to the Cache.
// If the Record with a specified UID already exists in the Cache,
// then that existing Record is moved to the Head Position.
// If the Record with a specified UID does not exist in the Cache,
// then a new Head is inserted into the Cache.
// If the Cache has is at its maximum Size, then the last Record (the Tail
// Record) is removed from the Cache.
// At the End, the Head Record is updated with the Data from the Function's
// Argument 'Data' Field.
func (cache *Cache) addARecord(r *Record) {

	var deletedUid RecordUID
	var existingRecord *Record
	var left *Record
	var newTail *Record
	var oldHead *Record
	var right *Record
	var uid RecordUID
	var uidExists bool

	uid = r.UID
	existingRecord, uidExists = cache.recordByUID[uid]
	// Record already exists in the Cache?
	if uidExists {

		// Cache's Size remains the same.
		// Move Item to the Head Position...

		if existingRecord == cache.head {
			// 1. The Record is already the Head.

			// Update Record's Data.
			cache.head.UpdateDataAndLAT(r.Data)

			return
		}

		// 1. Unlink the Record from its current Position.
		if existingRecord == cache.tail {

			// Unlink the Tail Record.
			newTail = cache.tail.previousItem
			newTail.nextItem = nil
			cache.tail = newTail
			existingRecord.nextItem = nil
			existingRecord.previousItem = nil

		} else {

			// Unlink the Record normally.
			left = existingRecord.previousItem
			right = existingRecord.nextItem
			left.nextItem = right
			right.previousItem = left
			existingRecord.nextItem = nil
			existingRecord.previousItem = nil
		}

		// 2. Link the Record as a new Head.
		oldHead = cache.head
		existingRecord.nextItem = oldHead
		oldHead.previousItem = existingRecord
		cache.head = existingRecord

		// Update Record's Data.
		cache.head.UpdateDataAndLAT(r.Data)

		return
	}

	// The Record does not exist in Cache.
	// Cache Size must be changed if it is not at its maximum Value.

	// 1. Remove the Tail Record if the Cache size is at its Maximum Value.
	if cache.size == cache.capacity {

		// Delete the Tail Record from the Fast Access Register.
		deletedUid = cache.tail.UID
		delete(cache.recordByUID, deletedUid)

		// Unlink the Tail Record.
		newTail = cache.tail.previousItem
		cache.tail.previousItem = nil
		newTail.nextItem = nil
		cache.tail = newTail
		cache.size--
	}

	// 2. Insert (Link) the Record as a new Head.
	if cache.size == 0 {

		// Insertion into an empty Cache.
		r.previousItem = nil
		r.nextItem = nil
		cache.head = r
		cache.tail = r
		cache.size++

	} else {

		// Insertion into a non-empty Cache.
		oldHead = cache.head
		r.nextItem = oldHead
		oldHead.previousItem = r
		cache.head = r
		cache.size++
	}

	// Add the Record to the Fast Access Register.
	cache.recordByUID[uid] = r

	// Update Record's Last Access Time.
	cache.head.UpdateLAT()

	return
}

// Deletes the Record from the Cache's List, Unlinks the Record.
// Also modifies the Fast Access Register.
func (cache *Cache) deleteARecord(r *Record) {

	var deletedUid RecordUID
	var left *Record
	var right *Record

	// Delete the Record from the Fast Access Register.
	deletedUid = r.UID
	delete(cache.recordByUID, deletedUid)

	// Change Cache Size.
	cache.size--

	// Remove the Record from the Double Link List.
	if cache.size == 0 {

		// New Cache Size is Zero => It must be empty.
		cache.head = nil
		cache.tail = nil
		return
	}

	// New Cache Size is >0 => It had 2 or more Records before Removal.

	if r == cache.head {

		// Removed Record is the Head Record.
		cache.head = r.nextItem
		cache.head.previousItem = nil
		r.nextItem = nil
		return

	} else if r == cache.tail {

		// Removed Record is the Tail Record.
		cache.tail = r.previousItem
		cache.tail.nextItem = nil
		r.previousItem = nil
		return

	} else {

		// Removed Record is removed normally.
		left = r.previousItem
		right = r.nextItem
		left.nextItem = right
		right.previousItem = left
		r.nextItem = nil
		r.previousItem = nil
		return
	}
}

// Initializes the Cache.
func (cache *Cache) initialize(
	c uint64,
	recordTTL int64,
) {

	cache.head = nil
	cache.tail = nil

	cache.capacity = c
	cache.size = 0

	cache.recordTTL = recordTTL
	cache.recordByUID = make(map[RecordUID]*Record)

	return
}

// Checks the Integrity of the Cache.
// This is a Self-Check Function intended to find Anomalies.
// This Function is not intended to be used in an ordinary Case.
// Returns 'true' if the Cache is in a good Shape.
func (cache *Cache) isIntegral() bool {

	var cursor *Record
	var cursorNextItem *Record
	var cursorPreviousItem *Record
	var head *Record
	var i uint64
	var record *Record
	var size uint64
	var sizeAnomaly bool
	var tail *Record

	// Check Fast Access Register.
	for _, record = range cache.recordByUID {
		if record == nil {
			return false
		}
	}

	// Prepare Data.
	head = cache.head
	tail = cache.tail
	size = cache.size

	// Capacity Check.
	if size > cache.capacity {
		return false
	}

	// Empty List?
	if size == 0 {
		if head != nil {
			return false
		}
		if tail != nil {
			return false
		}
		return true
	}

	// Single-Item List?
	if size == 1 {
		if head == nil {
			return false
		}
		if tail == nil {
			return false
		}
		if head != tail {
			return false
		}
		if head.previousItem != nil {
			return false
		}
		if tail.nextItem != nil {
			return false
		}
		return true
	}

	// List has two or more Items.

	// Check Head Corner.
	if head.previousItem != nil {
		return false
	}

	// Try to inspect all Items from Head to Tail.
	// This checks Connectivity by the 'next' Pointer.
	cursor = head
	cursorNextItem = cursor.nextItem
	i = 1
	for cursorNextItem != nil {
		cursor = cursorNextItem
		cursorNextItem = cursor.nextItem
		i++
		// Defence against Self-Loop Anomaly.
		if i > size {
			sizeAnomaly = true
			break
		}
	}
	if sizeAnomaly {
		// Size Anomaly can happen if we either have a Self-Loop in the Chain
		// or the Corner Item for some Reason is not the End of the Chain.
		return false
	}
	if i != size {
		return false
	}
	// We have stopped the Search at the first Break in the Chain.
	// Are we really there where we should be?
	if cursor != tail {
		// We have found a broken Connection.
		return false
	}

	// Check Tail Corner.
	if tail.nextItem != nil {
		return false
	}

	// Now, try to inspect all Items in a reversed Order.
	// This checks Connectivity by the 'previous' Pointer.
	cursor = tail
	cursorPreviousItem = cursor.previousItem
	i = 1
	for cursorPreviousItem != nil {
		cursor = cursorPreviousItem
		cursorPreviousItem = cursor.previousItem
		i++
		// Defence against Self-Loop Anomaly.
		if i > size {
			sizeAnomaly = true
			break
		}
	}
	if sizeAnomaly {
		// Size Anomaly can happen if we either have a Self-Loop in the Chain
		// or the Corner Item for some Reason is not the End of the Chain.
		return false
	}
	if i != size {
		return false
	}
	// We have stopped the Search at the first Break in the Chain.
	// Are we really there where we should be?
	if cursor != head {
		// We have found a broken Connection.
		return false
	}

	// All Clear.
	return true
}
