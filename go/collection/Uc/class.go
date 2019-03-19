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

// class.go.

package uc

// Collection Class Methods.

import (
	"errors"
	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
)

// Adds a new empty Class to the Collection and its Database.
// ID of the new Class will be automatically generated.
// Returns the created Class.
func (this *Uc) AddClass(
	className string,
) (class.Class, error) {

	var aClass class.Class
	var classId uint
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Generate Class ID.
	classId, err = this.collection.GetNextFreeClassId()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Create a Class with a Name.
	aClass = class.New(classId, className)

	// Add a Class to the Collection.
	err = this.addCollectionClass(aClass)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Refresh.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Create a Class Table in the Database.
	err = this.storage.CreateTableClass(aClass)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Add the Class to the Classes Table in the Database.
	err = this.storage.AddClass(aClass)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	// Create a Properties Table for a Class in the Database.
	err = this.storage.CreateTableClassProperties(aClass)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClass, err
	}

	return aClass, nil
}

// Adds a Class to the Collection.
func (this *Uc) addCollectionClass(
	aClass class.Class,
) error {

	return this.collection.AddClass(aClass)
}
