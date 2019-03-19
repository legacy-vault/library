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

// classObject.go.

package uc

// Collection Class Object Methods.

import (
	"errors"
	"xxx/Common"
	"xxx/Errorz"
	"xxx/Uc/Collection/Class"
	"xxx/Uc/Collection/Class/Object"
)

// Adds a new Object to the Collection Class.
// If objectId is set to Zero, it will be generated automatically.
// Returns the created Class Object.
func (this *Uc) AddClassObject(
	classId uint,
	objectId uint,
) (object.Object, error) {

	var aClass class.Class
	var aClassObject object.Object
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}

	// Preparations.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}

	// Generate Object ID if necessary.
	if object.IdIsEmpty(objectId) {
		objectId, err = aClass.GetNextFreeObjectId()
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return aClassObject, err
		}
	}

	// Create a Class Object.
	aClassObject = object.New(objectId)

	// Add a Class Object to the Collection.
	err = this.addCollectionClassObject(
		classId,
		aClassObject,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}

	// Refresh.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}
	aClassObject, err = aClass.GetObjectById(objectId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}

	// Add the Class Object to the Class Table in the Database.
	err = this.storage.AddClassObject(aClass, aClassObject)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObject, err
	}

	return aClassObject, nil
}

// Adds an new Object to a Collection Class.
func (this *Uc) addCollectionClassObject(
	classId uint,
	aClassObject object.Object,
) error {

	var aClass class.Class
	var err error

	// Get the Class.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		return err
	}

	// Add an Object to the Class.
	err = aClass.AddObject(aClassObject)
	if err != nil {
		return err
	}

	// Set the Class.
	err = this.collection.UpdateClass(classId, aClass)
	if err != nil {
		return err
	}

	return nil
}
