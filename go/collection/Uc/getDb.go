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

// getDb.go.

package uc

// High-Level Methods to get Collection Data from its Database.

// These Methods create new Structures in the Collection and are meant to be
// used only for the first Time Data Creation.

import (
	"errors"
	"fmt"

	"xxx/Common"
	databaseCollection "xxx/Database/Collection"
	"xxx/Errorz"
	"xxx/Uc/Collection/Class"
	"xxx/Uc/Collection/Class/Object"
	objectProperty "xxx/Uc/Collection/Class/Object/Property"
	classProperty "xxx/Uc/Collection/Class/Property"
	"xxx/Uc/Collection/Class/Property/KindSettings"
)

// Gets the Classes List from a Database.
func (this *Uc) getDbClasses() error {

	var aClass class.Class
	var dbClasses []databaseCollection.DatabaseClass
	var err error

	// Select the Classes from a Database.
	dbClasses, err = this.storage.ReadClasses()
	if err != nil {
		return err
	}

	// Save the Classes into the Collection.
	for _, dbClass := range dbClasses {

		// Verify the Membership of the Entity read from the Database.
		// Membership is protected by ID Fields and SQL Foreign Keys.
		// Not needed here, as we are at the Root Element
		// which can not be verified.

		// Convert Class read from Database into a common Format.
		aClass = class.New(
			dbClass.Id,
			dbClass.Name,
		)

		// Add the Class into a Collection.
		err = this.collection.AddClass(aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Gets the Class Object Property Values from a Database.
func (this *Uc) getDbClassObjectProperties() error {

	var aClass class.Class
	var aClassObject object.Object
	var aClassObjectProperty objectProperty.Property
	var aClassProperty classProperty.Property
	var classes map[uint]class.Class
	var classId uint
	var classObjectId uint
	var classObjects map[uint]object.Object
	var classProperties map[uint]classProperty.Property
	var classPropertyId uint
	var dbClassObjectProperty databaseCollection.DatabaseClassObjectProperty
	var dbClassObjectPropertyIsSet bool
	var err error
	var ok bool

	// For each Class...
	classes = this.collection.GetClasses()
	for classId, aClass = range classes {

		// For each Class Object...
		classObjects = aClass.GetObjects()
		classProperties = aClass.GetProperties()
		for classObjectId, aClassObject = range classObjects {

			// For each Class Property...
			for classPropertyId, aClassProperty = range classProperties {

				// Select the Class Object Property Value from a Database.
				dbClassObjectProperty, err =
					this.storage.ReadClassObjectProperty(
						aClass,
						aClassProperty,
						aClassObject,
					)
				if err != nil {
					return err
				}

				// Check whether Property Value is set.
				// Property may be not set in the Database.
				dbClassObjectPropertyIsSet = dbClassObjectProperty.ValueIsSet
				if !dbClassObjectPropertyIsSet {
					continue
				}

				// Verify the Membership of the Entity read from the Database.
				// Membership is protected by ID Fields and SQL Foreign Keys.
				ok = this.membershipClassObjectProperty(
					aClass,
					aClassObject,
					aClassProperty,
					dbClassObjectProperty,
				)
				if !ok {
					err = fmt.Errorf(
						ErrFormatClassObjectPropertyMembership,
					)
					return err
				}

				// Get the Class Property Value Type from the Collection.
				// Notes.
				// This is needed due to the SQL Driver in Golang being not
				// competent. It reports all integer Values as 'INT'
				// indifferently whether it is a signed 'INT' or an unsigned
				// 'INT UNSIGNED'!
				dbClassObjectProperty.ValueDbTypeFromCollection =
					aClassProperty.GetKind()

				// Convert the raw Database Object Property Value
				// into something more usable.
				err = dbClassObjectProperty.ConvertValueDbToUsable()
				if err != nil {
					return err
				}

				// Convert Class Object Property read from Database into a
				// common Format.
				aClassObjectProperty = objectProperty.NewObjectProperty(
					classPropertyId,
					dbClassObjectProperty.Value,
				)

				// Add the Property into a Class Object.
				err = aClassObject.AddProperty(aClassObjectProperty)
				if err != nil {
					return err
				}

			}

			// Save the Class Object Changes.
			err = aClass.UpdateObject(classObjectId, aClassObject)
			if err != nil {
				return err
			}
		}

		// Save the Class Changes.
		err = this.collection.UpdateClass(classId, aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Gets the Class Objects List from a Database.
func (this *Uc) getDbClassObjects() error {

	var aClass class.Class
	var aClassObject object.Object
	var classes map[uint]class.Class
	var classId uint
	var dbClassObjects []databaseCollection.DatabaseClassObject
	var err error
	var ok bool

	// For each Class...
	classes = this.collection.GetClasses()
	for classId, aClass = range classes {

		// Select the Class Objects from a Database.
		dbClassObjects, err = this.storage.ReadClassObjects(aClass)
		if err != nil {
			return err
		}

		// Save the Objects into the Collection Class.
		for _, dbClassObject := range dbClassObjects {

			// Verify the Membership of the Entity read from the Database.
			// Membership is protected by ID Fields and SQL Foreign Keys.
			ok = this.membershipClassObject(
				aClass,
				dbClassObject,
			)
			if !ok {
				err = fmt.Errorf(
					ErrFormatClassObjectMembership,
				)
				return err
			}

			// Convert Class Object read from Database into a common Format.
			aClassObject = object.New(
				dbClassObject.Id,
			)

			// Add the Object into a Class.
			err = aClass.AddObject(aClassObject)
			if err != nil {
				return err
			}

			// Save the Class Changes.
			err = this.collection.UpdateClass(classId, aClass)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Gets the Class Properties List for all Classes from a Database.
func (this *Uc) getDbClassProperties() error {

	var aClass class.Class
	var aClassProperty classProperty.Property
	var aClassPropertyKind kindsettings.KindSettings
	var classId uint
	var classes map[uint]class.Class
	var dbClassProperties []databaseCollection.DatabaseClassProperty
	var err error
	var ok bool

	// For each Class...
	classes = this.collection.GetClasses()
	for classId, aClass = range classes {

		// Select the Class Properties for a Class from a Database.
		dbClassProperties, err = this.storage.ReadClassProperties(aClass)
		if err != nil {
			return err
		}

		// Save the Properties into the Collection Class.
		for _, dbClassProperty := range dbClassProperties {

			// Verify the Membership of the Entity read from the Database.
			// Membership is protected by ID Fields and SQL Foreign Keys.
			ok = this.membershipClassProperty(
				aClass,
				dbClassProperty,
			)
			if !ok {
				err = fmt.Errorf(
					ErrFormatClassPropertyMembership,
				)
				return err
			}

			// Convert Class Property read from Database into a common Format.
			aClassPropertyKind, err = kindsettings.NewWithDbType(
				dbClassProperty.DbType,
			)
			if err != nil {
				return err
			}
			aClassProperty, err = classProperty.NewReferenceProperty(
				dbClassProperty.Id,
				dbClassProperty.Name,
				dbClassProperty.Description,
				aClassPropertyKind,
			)
			if err != nil {
				return err
			}

			// Add the Property into a Class.
			err = aClass.AddProperty(aClassProperty)
			if err != nil {
				return err
			}

			// Save the Class Changes.
			err = this.collection.UpdateClass(classId, aClass)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Gets the Collection's Data from its Database.
func (this *Uc) GetDbCollection() error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check Status.
	if this.isInitialized() != true {
		err = errors.New(ErrNotInitialized)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Create Everything in the Database.
	err = this.getDbCollection()
	if err != nil {
		return err
	}

	return nil
}

// Gets the Collection's Data from its Database.
func (this *Uc) getDbCollection() error {

	var err error

	// Get Everything from the Database.
	err = this.getDbClasses()
	if err != nil {
		return err
	}
	err = this.getDbClassProperties()
	if err != nil {
		return err
	}
	err = this.getDbClassObjects()
	if err != nil {
		return err
	}
	err = this.getDbClassObjectProperties()
	if err != nil {
		return err
	}

	return nil
}
