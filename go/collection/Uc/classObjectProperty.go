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

// classObjectProperty.go.

package uc

// Collection Class Object Property Methods.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property"
)

// Adds a Property to the Collection Class Object.
func (this *Uc) AddClassObjectProperty(
	classId uint,
	objectId uint,
	propertyId uint,
	propertyValue interface{},
) (property.Property, error) {

	var aClass class.Class
	var aClassObject object.Object
	var aClassObjectProperty property.Property
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	// Preparations.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}
	aClassObject, err = aClass.GetObjectById(objectId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	// Generate Object ID if necessary.
	if property.IdIsEmpty(propertyId) {
		err = errors.New(ErrPropertyId)
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	// Create a Class Object Property.
	aClassObjectProperty = property.NewObjectProperty(
		propertyId,
		propertyValue,
	)

	// Add a Property to the Collection Class Object.
	err = this.addCollectionClassObjectProperty(
		classId,
		objectId,
		aClassObjectProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	// Refresh.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}
	aClassObject, err = aClass.GetObjectById(objectId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}
	aClassObjectProperty, err = aClassObject.GetPropertyById(propertyId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	// Add the Value to the Class Property Table in the Database.
	err = this.storage.AddClassObjectProperty(
		aClass,
		aClassObject,
		aClassObjectProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassObjectProperty, err
	}

	return aClassObjectProperty, nil
}

// Adds a Property to the Collection Class Object.
func (this *Uc) addCollectionClassObjectProperty(
	classId uint,
	objectId uint,
	aClassObjectProperty property.Property,
) error {

	var aClass class.Class
	var aClassObject object.Object

	var err error

	// Get the Class.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		return err
	}

	// Get the Object.
	aClassObject, err = aClass.GetObjectById(objectId)
	if err != nil {
		return err
	}

	// Add a Property to the Class Object.
	err = aClassObject.AddProperty(aClassObjectProperty)
	if err != nil {
		return err
	}

	// Save the Changes.
	err = aClass.UpdateObject(objectId, aClassObject)
	if err != nil {
		return err
	}
	err = this.collection.UpdateClass(classId, aClass)
	if err != nil {
		return err
	}

	return nil
}
