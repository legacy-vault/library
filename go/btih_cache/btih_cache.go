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

// btih_cache.go.

// BTIH Cache.

// This Package provides a cached Access to the File Storage.
// BTIH is used as a Key in Cache List.
// File Contents are used as a Value in Cache List.
// Files are stored in a linear Order and their Names follow this Format:
// 		"<BTIH>.<Extension>".

package btih_cache

import (
	"errors"
	"log"

	"github.com/legacy-vault/library/go/fixed_size_bubble_cache"
	"github.com/legacy-vault/library/go/simple_file_storage"
)

const FileExtension = ".torrent"

type BTIHCache struct {
	cache    fsbcache.Cache
	storage  sfstorage.Storage
	isActive bool
}

var (
	ErrCache         = errors.New("Cache Error")
	ErrStorage       = errors.New("Storage Error")
	ErrTypeAssertion = errors.New("Type Assertion Error")
)

// Creates a new BTIH Cache.
// * [Capacity] is the maximum Size of Cache. New Records added when the Cache
// is at its maximum Size, will remove the Tail Record. Tail Record is a Record
// with the oldest Access Time.
// * [RecordTTL] is the Period of Time, after which the Record is considered
// outdated. If the requested Record is outdated, it is removed from the Cache.
// Measured in Seconds.
// * [Storage Root Folder] is the Root Folder where all Files are stored.
func New(
	capacity uint64,
	recordTTL int64,
	storageRootFolder string,
) (*BTIHCache, error) {

	var cache *fsbcache.Cache
	var err error
	var result BTIHCache
	var storage *sfstorage.Storage

	// Prepare a Simple File Storage.
	storage, err = sfstorage.New(storageRootFolder)
	if err != nil {
		return nil, err
	}

	// Prepare a Fixed Size Bubble Cache.
	cache = fsbcache.New(capacity, recordTTL)

	// Save Results.
	result.cache = *cache
	result.storage = *storage
	result.isActive = true

	return &result, nil
}

// Gets File's Contents by the BTIH.
func (bc *BTIHCache) GetFileByBTIH(btih string) []byte {

	var err error
	var ifc interface{}
	var fileContents []byte
	var fileIsCached bool
	var filePath string
	var ok bool

	// Check State.
	if bc.isActive == false {
		return []byte{}
	}

	// Check the Cache.
	ifc, fileIsCached = bc.cache.GetRecordDataByUID(btih)
	if fileIsCached {

		// Get File's Contents from Cache.
		fileContents, ok = ifc.([]byte)
		if !ok {
			log.Println(ErrTypeAssertion, ifc)
			return []byte{}
		}

		return fileContents
	}

	// File is not cached.

	// Get File from Storage.
	filePath = btih + FileExtension
	fileContents, err = bc.storage.GetFileContents(filePath)
	if err != nil {
		log.Println(ErrStorage, err)
		return []byte{}
	}

	// Add File into Cache.
	err = bc.cache.AddRecord(btih, fileContents)
	if err != nil {
		log.Println(ErrCache, err)
		return []byte{}
	}

	// All Clear.
	return fileContents
}

// Checks whether the BTIH specified is cached.
func (bc *BTIHCache) IsCached(btih string) bool {

	var isCached bool

	isCached = bc.cache.RecordUIDIsActive(btih)
	if !isCached {
		return false
	}

	return true
}

// Stops the BTIH Cache Object.
// Returns 'true' on Success.
// After this Method has been successfully applied,
// no further Cache Usage is possible!
func (bc *BTIHCache) Stop() bool {

	var ok bool

	// Clear Cache.
	ok = bc.cache.Clear()
	if !ok {
		return false
	}

	bc.isActive = false

	return true
}
