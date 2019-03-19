// ============================================================================
//
// Copyright © 2019 by McArcher.
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
// ============================================================================
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2019-03-19.
// Web Site Address is an Address in the global Computer Internet Network.
//
// ============================================================================

// open.go.

package uc

// User Collection Methods for Initialization.

import (
	"errors"

	"xxx/Common"
	"xxx/Errorz"
	"xxx/Uc/Collection"
	"xxx/Uc/Collection/Class"
	"xxx/Uc/Storage/Settings"
)

// Configures the User's Collection.
func (this *Uc) configure(
	settings settings.Settings,
) error {

	var err error

	// Apply the Storage Settings.
	err = this.storage.Configure(settings)
	if err != nil {
		return err
	}

	return nil
}

// Connects to the User's Collection's Database.
func (this *Uc) connect() error {

	var err error

	// Connect.
	err = this.storage.Connect()
	if err != nil {
		return err
	}

	return nil
}

// Configures the User's Collection and connects to its Database.
func (this *Uc) init(
	settings settings.Settings,
) error {

	var err error

	// Apply the Settings.
	err = this.configure(settings)
	if err != nil {
		return err
	}

	// Connect.
	err = this.connect()
	if err != nil {
		return err
	}

	// Load SQL Command Templates.
	err = this.storage.LoadSqlCommandTemplates()
	if err != nil {
		return err
	}

	// Set Status.
	this.initializationIsDone = true
	if this.isInitialized() != true {
		err = errors.New(ErrNotInitialized)
		return err
	}

	return nil
}

// Returns the 'initializationIsDone' Status Flag.
func (this Uc) isInitialized() bool {

	return this.initializationIsDone
}

// Returns the 'isOpen' Status Flag.
func (this Uc) IsOpened() bool {

	return this.isOpen
}

// Opens a Collection from the Database.
func (this *Uc) Open() error {

	var classes map[uint]class.Class
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.isInitialized() != true {
		err = errors.New(ErrNotInitialized)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.IsOpened() == true {
		err = errors.New(ErrOpened)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Get the Collection's Data from its Database.
	err = this.GetDbCollection()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Prepare Parameters.
	classes = this.collection.GetClasses()

	// Check all the Tables of the Collection in the Database.
	err = this.storage.CheckDbCollection(classes)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set Status.
	this.isOpen = true

	return nil
}

// Creates a new Collection,
// Applies the Settings,
// Connects to the Database.
func New(
	collectionName string,
	settings settings.Settings,
) (*Uc, error) {

	var err error
	var result *Uc

	result = new(Uc)

	// Crate the empty Collection.
	result.collection = collection.New(collectionName)

	// Apply the Settings and connect to the Database.
	err = result.init(settings)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return nil, err
	}

	return result, nil
}
