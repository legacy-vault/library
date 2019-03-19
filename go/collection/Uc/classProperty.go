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

// classProperty.go.

package uc

// Collection Class Property Methods.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
)

// Adds a new Property to the Collection Class.
// If PropertyId is set to Zero, it will be generated automatically.
// Returns the created Class Property.
func (this *Uc) AddClassProperty(
	classId uint,
	propertyId uint,
	propertyName string,
	propertyDescription string,
	propertyKindSettings kindsettings.KindSettings,
) (property.Property, error) {

	var aClass class.Class
	var aClassProperty property.Property
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Preparations.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Generate Property ID if necessary.
	if property.IdIsEmpty(propertyId) {
		propertyId, err = aClass.GetNextFreePropertyId()
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return aClassProperty, err
		}
	}

	// Check the Property Kind Settings.
	if !propertyKindSettings.IsValid() {
		err = errors.New(ErrPropertyKindSettings)
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Create a Class Property.
	aClassProperty, err = classProperty.NewReferenceProperty(
		propertyId,
		propertyName,
		propertyDescription,
		propertyKindSettings,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Add a Class Property to the Collection.
	err = this.addCollectionClassProperty(
		classId,
		aClassProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Refresh.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}
	aClassProperty, err = aClass.GetPropertyById(propertyId)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Create a Class Property Table in the Database.
	err = this.storage.CreateTableClassProperty(aClass, aClassProperty)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	// Add the Class Property to the Class Properties Table in the Database.
	err = this.storage.AddClassProperty(aClass, aClassProperty)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return aClassProperty, err
	}

	return aClassProperty, nil
}

// Adds a new Property to the Collection Class.
func (this *Uc) addCollectionClassProperty(
	classId uint,
	aClassProperty property.Property,
) error {

	var aClass class.Class
	var err error

	// Get the Class.
	aClass, err = this.collection.GetClassById(classId)
	if err != nil {
		return err
	}

	// Add a Property to the Class.
	err = aClass.AddProperty(aClassProperty)
	if err != nil {
		return err
	}

	// Save the Class.
	err = this.collection.UpdateClass(classId, aClass)
	if err != nil {
		return err
	}

	return nil
}
