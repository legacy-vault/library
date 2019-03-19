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

// Object.go.

package object

// Collection Class Object.

import (
	"errors"
	"fmt"
	"math"
	"time"

	"xxx/Common"
	"xxx/Errorz"
	"xxx/Uc/Collection/Class/Object/Property"
)

const ErrorReporter = "Object"

const ErrFormatBadPropertyId = "Bad Property ID '%v'"
const ErrFormatDamagedObjectPropertyId = "Damaged Object Property ID " +
	"(ID='%v', Index='%v') in Object with ID '%v'"
const ErrFormatDuplicatePropertyId = "Duplicate Property ID '%v' (Object ID is '%v')"
const ErrFormatDuplicatePropertyName = "Duplicate Property Name '%v' (Object ID is '%v')"
const ErrFormatMissingPropertyId = "Missing Property ID '%v' (Object ID is '%v')"

const ErrFormatBadId = "Bad ID '%v'"
const ErrNoFreePropertyId = "No free Property IDs are left"

const IdNone = 0
const IdFirstAvailable = IdNone + 1

type Object struct {

	// Parameters & Settings.
	id             uint
	timeOfCreation time.Time
	timeOfUpdate   time.Time

	// Sub-Items.
	// Associative Array of Properties (Property Types). Key is PropertyId.
	properties map[uint]property.Property
}

// Adds a Property to an Object.
// If the ID is empty, a new one is generated.
func (this *Object) AddProperty(
	aProperty property.Property,
) error {

	var duplicateId bool
	var err error
	var propertyId uint

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Properties List Initialization.
	if this.properties == nil {
		this.properties = make(map[uint]property.Property)
	}

	// Preparations.
	propertyId = aProperty.GetId()

	// Is an Object Property ID empty?
	if property.IdIsEmpty(propertyId) {
		// Get the next free ID.
		propertyId, err = this.GetNextFreePropertyId()
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
	}

	// Duplicate ID Check.
	duplicateId = this.PropertyIdIsUsed(propertyId)
	if duplicateId {
		err = fmt.Errorf(
			ErrFormatDuplicatePropertyId,
			propertyId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify the Property's Integrity.
	err = aProperty.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Property.
	this.setProperty(propertyId, aProperty)

	return nil
}

// Returns the 'id' Field.
func (this Object) GetId() uint {

	return this.id
}

// Returns the next free ID for a Property.
func (this Object) GetNextFreePropertyId() (uint, error) {

	var err error
	var id uint

	id = property.IdFirstAvailable
	if !this.PropertyIdIsUsed(id) {
		return id, nil
	}
	for {
		id++
		if !this.PropertyIdIsUsed(id) {
			return id, nil
		}
		if id == math.MaxUint64 {
			break
		}
	}

	err = errors.New(ErrNoFreePropertyId)
	err = errorz.Report(ErrorReporter, err)
	return property.IdNone, err
}

// Returns the 'properties' Field.
func (this Object) GetProperties() map[uint]property.Property {

	return this.properties
}

// Returns a Property of the Object by the Property ID.
func (this Object) GetPropertyById(
	propertyId uint,
) (property.Property, error) {

	var err error
	var idExists bool
	var result property.Property

	// Duplicate ID Check.
	result, idExists = this.properties[propertyId]
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingPropertyId,
			propertyId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, nil
}

// Checks whether the specified Property ID is used in the Object.
func (this Object) PropertyIdIsUsed(
	propertyId uint,
) bool {

	var idExists bool

	// Duplicate ID Check.
	_, idExists = this.properties[propertyId]

	return idExists
}

// Sets a Property in the Object.
func (this *Object) setProperty(
	propertyId uint,
	aProperty property.Property,
) {

	this.properties[propertyId] = aProperty
}

// Updates a Property in an Object.
// The ID must be a real existing ID.
func (this *Object) UpdateProperty(
	propertyId uint,
	aProperty property.Property,
) error {

	var err error
	var idExists bool

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if propertyId != aProperty.GetId() {
		err = fmt.Errorf(
			common.ErrFormatUpdateIdMismatch,
			propertyId,
			aProperty.GetId(),
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Properties List Initialization.
	if this.properties == nil {
		this.properties = make(map[uint]property.Property)
	}

	// Is an Object Property ID empty?
	if property.IdIsEmpty(propertyId) {
		err = fmt.Errorf(
			ErrFormatBadPropertyId,
			propertyId,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// ID Exists?
	idExists = this.PropertyIdIsUsed(propertyId)
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingPropertyId,
			propertyId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify the Property's Integrity.
	err = aProperty.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Property.
	this.setProperty(propertyId, aProperty)

	return nil
}

// Verifies the Integrity of an Object.
func (this Object) VerifyIntegrity() error {

	var err error

	// Verify.
	err = this.verifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Verifies the Integrity of an Object.
func (this Object) verifyIntegrity() error {

	var err error

	// Properties List Initialization.
	if this.properties == nil {
		this.properties = make(map[uint]property.Property)
	}

	// Check ID.
	if this.id == IdNone {
		err = fmt.Errorf(
			ErrFormatBadId,
			this.id,
		)
		return err
	}

	// Verify Properties Array Index Matching.
	for aPropertyId, aProperty := range this.properties {
		if aPropertyId != aProperty.GetId() {
			err = fmt.Errorf(
				ErrFormatDamagedObjectPropertyId,
				aProperty.GetId(),
				aPropertyId,
				this.id,
			)
			return err
		}
	}

	// Verify the Integrity of Object's Properties.
	for _, aProperty := range this.properties {
		err = aProperty.VerifyIntegrity()
		if err != nil {
			return err
		}
	}

	return nil
}

// Checks whether the specified Object ID is empty.
func IdIsEmpty(
	objectId uint,
) bool {

	if objectId == IdNone {
		return true
	}

	return false
}

// Returns a new Object.
func New(
	objectId uint,
) Object {

	var result Object

	result.id = objectId

	return result
}
