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

// createDb.go.

package uc

// High-Level Methods to create Collection Data in its Database.

// These Methods create new Structures in the Database and are meant to be
// used only for the first Time Data Creation.

import (
	"errors"

	"xxx/Common"
	"xxx/Errorz"
	"xxx/Uc/Collection/Class"
	"xxx/Uc/Collection/Class/Object"
	objectProperty "xxx/Uc/Collection/Class/Object/Property"
	classProperty "xxx/Uc/Collection/Class/Property"
)

// Creates the Classes List in the Database:
//	* creates a Classes Table;
//	* adds each Class to the Classes Table;
//	* creates a Class Table for each Class.
func (this *Uc) createDbClasses() error {

	var err error

	// Create Classes Table in the Database.
	err = this.storage.CreateTableClasses()
	if err != nil {
		return err
	}

	// Create the Tables and update the Class List.
	for _, aClass := range this.collection.GetClasses() {

		// Add the Class to the Classes Table in the Database.
		err = this.storage.AddClass(aClass)
		if err != nil {
			return err
		}

		// Create a Class Table in the Database.
		err = this.storage.CreateTableClass(aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Creates Class Object Properties in the Database:
//	*	adds each Class Object Property (for each Class Object, for each Class)
//		to each Class Object Property Table.
func (this *Uc) createDbClassObjectProperties() error {

	var classes map[uint]class.Class
	var classObjects map[uint]object.Object
	var classObjectProperties map[uint]objectProperty.Property
	var err error

	classes = this.collection.GetClasses()
	for _, aClass := range classes {

		classObjects = aClass.GetObjects()
		for _, aClassObject := range classObjects {

			// Add all Class Object Properties into Database.
			classObjectProperties = aClassObject.GetProperties()
			for _, aClassObjectProperty := range classObjectProperties {

				// Add a Class Object Property into Database.
				err = this.storage.AddClassObjectProperty(
					aClass,
					aClassObject,
					aClassObjectProperty,
				)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Creates Class Objects in the Database:
//	* adds each Class Object (for each Class) to the Class Table.
func (this *Uc) createDbClassObjects() error {

	var classes map[uint]class.Class
	var classObjects map[uint]object.Object
	var err error

	classes = this.collection.GetClasses()
	for _, aClass := range classes {

		classObjects = aClass.GetObjects()
		for _, aClassObject := range classObjects {

			// Add a Class Object into Database.
			err = this.storage.AddClassObject(
				aClass,
				aClassObject,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates Class Properties List in the Database:
//	* creates a Class Properties Table for each Class;
//	* adds each Class Property (for each Class) to the Class Properties Table;
//	* creates a Class Property Table for each Class Property (for each Class).
func (this *Uc) createDbClassProperties() error {

	var classes map[uint]class.Class
	var classProperties map[uint]classProperty.Property
	var err error

	// Create the Tables and update the Class List.
	classes = this.collection.GetClasses()
	for _, aClass := range classes {

		// Create a Properties Table for a Class in the Database.
		err = this.storage.CreateTableClassProperties(aClass)
		if err != nil {
			return err
		}

		classProperties = aClass.GetProperties()
		for _, aClassProperty := range classProperties {

			// Add a Property to a Class in the Database.
			err = this.storage.AddClassProperty(
				aClass,
				aClassProperty,
			)
			if err != nil {
				return err
			}

			// Create a Property Table for a Class in the Database.
			err = this.storage.CreateTableClassProperty(
				aClass,
				aClassProperty,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates Collection's Data in its Database.
func (this *Uc) CreateDbCollection() error {

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
	if this.IsOpened() == true {
		err = errors.New(ErrOpened)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Create Everything in the Database.
	err = this.createDbCollection()
	if err != nil {
		return err
	}

	return nil
}

// Creates Collection's Data in its Database.
func (this *Uc) createDbCollection() error {

	var err error

	// Create Everything in the Database.
	err = this.createDbClasses()
	if err != nil {
		return err
	}
	err = this.createDbClassProperties()
	if err != nil {
		return err
	}
	err = this.createDbClassObjects()
	if err != nil {
		return err
	}
	err = this.createDbClassObjectProperties()
	if err != nil {
		return err
	}

	return nil
}
