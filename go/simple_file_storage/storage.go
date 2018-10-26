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

// storage.go.

// Simple File Storage.

package sfstorage

import (
	"io/ioutil"
	"path"
)

type Storage struct {
	basePath string
}

// Creates a new Simple File Storage.
// * [Storage Root Path] is the Path to the Folder, where all the Files are
// stored. Must be a valid Path.
func New(storageRootPath string) (*Storage, error) {

	var err error
	var rootPathExists bool
	var storage *Storage

	// Prepare Data.
	storage = new(Storage)
	storageRootPath = path.Clean(storageRootPath)

	// Check Folder's Existence.
	rootPathExists = FileExists(storageRootPath)

	// Create Folder if it does not exist.
	if !rootPathExists {
		err = CreateFolder(storageRootPath)
		if err != nil {
			return nil, err
		}
	}

	// Save Root Path.
	storage.basePath = storageRootPath

	return storage, nil
}

// Returns File's Contents.
// The File is specified by its relative (to Storage Root Folder) Path.
func (storage Storage) GetFileContents(relPath string) ([]byte, error) {

	var absPath string
	var contents []byte
	var err error

	// Prepare absolute Path.
	absPath = path.Join(storage.basePath, relPath)

	// Get File's Contents.
	contents, err = ioutil.ReadFile(absPath)
	if err != nil {
		return []byte{}, err
	}

	return contents, nil
}

// Returns the Root Path of the Storage.
func (storage Storage) GetRootPath() string {

	return storage.basePath
}
