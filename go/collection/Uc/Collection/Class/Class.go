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

// Class.go.

package class

// Collection Class.

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	objectProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	"github.com/legacy-vault/library/go/collection/Utf8"
)

const ErrorReporter = "Class"

const ErrFormatBadClassId = "Bad Class ID '%v'"
const ErrFormatBadClassName = "Bad Class Name '%v'"
const ErrFormatBadObjectId = "Bad Object ID '%v'"
const ErrFormatBadPropertyId = "Bad Property ID '%v'"
const ErrFormatDamagedExistingClassObjectId = "Damaged " +
	"existing Class Object ID " +
	"(ID='%v', Index='%v') in Class with ID '%v'"
const ErrFormatDamagedReplacementClassObjectId = "Damaged " +
	"Replacement Class Object ID " +
	"(ID='%v', Index='%v') in Class with ID '%v'"
const ErrFormatDamagedExistingClassPropertyId = "Damaged " +
	"existing Class Property ID " +
	"(ID='%v', Index='%v') in Class with ID '%v'"
const ErrFormatDuplicateClassId = "Duplicate Class ID '%v' " +
	"(Class Name is '%v')"
const ErrFormatDuplicateClassName = "Duplicate Class Name '%v'"
const ErrFormatDuplicateObjectId = "Duplicate Object ID '%v' " +
	"(Class Name is '%v')"
const ErrFormatDuplicatePropertyId = "Duplicate Property ID '%v' " +
	"(Class Name is '%v')"
const ErrFormatDuplicatePropertyName = "Duplicate Property Name '%v' " +
	"(Class Name is '%v')"
const ErrFormatMissingObjectId = "Missing Object ID '%v' " +
	"(Class Name is '%v')"
const ErrFormatUnregisteredObjectPropertyId = "Object Property ID '%v' is not " +
	"registered in the Class '%v'"
const ErrFormatMissingPropertyName = "Missing Property Name '%v' " +
	"(Class Name is '%v')"
const ErrFormatMissingPropertyId = "Missing Property ID '%v' " +
	"(Class ID is '%v')"
const ErrNoFreeObjectId = "No free Object IDs are left"
const ErrNoFreePropertyId = "No free Property IDs are left"
const ErrFormatDamagedPropertyId = "Damaged Property ID " +
	"(ID='%v', Index='%v') in Class with ID '%v'"
const ErrFormatDamagedObjectId = "Damaged Object ID " +
	"(ID='%v', Index='%v') in Class with ID '%v'"
const ErrFormatPropertyNameConflict = "Property Update Name Conflict: " +
	"'%v' vs '%v' (Class Name is '%v')"

const NameLengthMax = 255 // VARCHAR(255).

const ErrFormatBadId = "Bad ID '%v'"

const IdNone = 0
const IdFirstAvailable = IdNone + 1

// Notes:
//
//	1.	Class Name must contain only small ASCII Letters.
//	2.	Class Name is used in Database Table Names.

type Class struct {

	// Parameters & Settings.
	id             uint
	name           string
	timeOfCreation time.Time
	timeOfUpdate   time.Time

	// Reference for Class Properties (without Values).
	// Associative Array of Properties (Property Types). Key is PropertyId.
	// This Reference is used for Database Initialization.
	properties map[uint]property.Property

	// Sub-Items.
	// Associative Array of Class Objects. Key is ObjectId.
	objects map[uint]object.Object
}

// Adds an Object to the Class.
// If the ID is empty, a new one is generated.
func (this *Class) AddObject(
	anObject object.Object,
) error {

	var duplicateId bool
	var err error
	var objectId uint

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Objects List Initialization.
	if this.objects == nil {
		this.objects = make(map[uint]object.Object)
	}

	// Preparations.
	objectId = anObject.GetId()

	// Is an Object ID empty?
	if object.IdIsEmpty(objectId) {
		// Get the next free ID.
		objectId, err = this.GetNextFreeObjectId()
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
	}

	// Duplicate ID Check.
	duplicateId = this.ObjectIdIsUsed(objectId)
	if duplicateId {
		err = fmt.Errorf(
			ErrFormatDuplicateObjectId,
			objectId,
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify Object's Integrity.
	err = anObject.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Ensure that Object's Properties are registered in the Class.
	_, err = this.ObjectPropertiesAreRegistered(anObject)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Object.
	this.setObject(objectId, anObject)

	return nil
}

// Adds a Property to the Class.
// If the ID is empty, a new one is generated.
func (this *Class) AddProperty(
	aProperty classProperty.Property,
) error {

	var duplicateId bool
	var duplicateName bool
	var err error
	var propertyId uint
	var propertyName string

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Properties List Initialization.
	if this.properties == nil {
		this.properties = make(map[uint]classProperty.Property)
	}

	// Preparations.
	propertyName = aProperty.GetName()
	propertyId = aProperty.GetId()

	// Is a Property ID empty?
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
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Duplicate Name Check.
	duplicateName = this.PropertyNameIsUsed(propertyName)
	if duplicateName {
		err = fmt.Errorf(
			ErrFormatDuplicatePropertyName,
			propertyName,
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify Property's Integrity.
	err = aProperty.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Property.
	this.setProperty(propertyId, aProperty)

	return nil
}

// Checks Class Name.
func (this Class) CheckName() error {

	var err error

	// Check.
	err = this.checkName()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Checks Class Name.
func (this Class) checkName() error {

	var err error

	// Check Length.
	if len(this.name) > NameLengthMax {
		err = fmt.Errorf(
			ErrFormatBadClassName,
			this.name,
		)
		return err
	}

	// Check Symbols.
	if utf8.StringHasOnlyASCIILetterOrNumber(this.name) != true {
		err = fmt.Errorf(
			ErrFormatBadClassName,
			this.name,
		)
		return err
	}

	return nil
}

// Returns the 'id' Field.
func (this Class) GetId() uint {

	return this.id
}

// Returns the 'name' Field.
func (this Class) GetName() string {

	return this.name
}

// Returns the next free ID for an Object.
func (this Class) GetNextFreeObjectId() (uint, error) {

	var err error
	var id uint

	id = object.IdFirstAvailable
	if !this.ObjectIdIsUsed(id) {
		return id, nil
	}
	for {
		id++
		if !this.ObjectIdIsUsed(id) {
			return id, nil
		}
		if id == math.MaxUint64 {
			break
		}
	}

	err = errors.New(ErrNoFreeObjectId)
	err = errorz.Report(ErrorReporter, err)
	return object.IdNone, err
}

// Returns the next free ID for a Property.
func (this Class) GetNextFreePropertyId() (uint, error) {

	var err error
	var id uint

	id = classProperty.IdFirstAvailable
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
	return classProperty.IdNone, err
}

// Returns an existing Object of the Class.
func (this Class) getObject(
	objectId uint,
) object.Object {

	return this.objects[objectId]
}

// Returns an Object of the Class by the Object ID.
func (this Class) GetObjectById(
	objectId uint,
) (object.Object, error) {

	var err error
	var idExists bool
	var result object.Object

	// ID Exists?
	idExists = this.ObjectIdIsUsed(objectId)
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingObjectId,
			objectId,
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Get.
	result = this.getObject(objectId)

	return result, nil
}

// Returns the 'objects' Field.
func (this Class) GetObjects() map[uint]object.Object {

	return this.objects
}

// Returns the 'properties' Field.
func (this Class) GetProperties() map[uint]classProperty.Property {

	return this.properties
}

// Returns an existing Property of the Class.
func (this Class) getProperty(
	propertyId uint,
) classProperty.Property {

	return this.properties[propertyId]
}

// Returns a Property of the Class by the Property ID.
func (this Class) GetPropertyById(
	propertyId uint,
) (classProperty.Property, error) {

	var err error
	var idExists bool
	var result classProperty.Property

	// ID Exists?
	idExists = this.PropertyIdIsUsed(propertyId)
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingPropertyId,
			propertyId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Get.
	result = this.getProperty(propertyId)

	return result, nil
}

// Returns a Property of the Class by the Property Name.
// During the Search, Property Names are converted into a low Case.
func (this Class) GetPropertyByName(
	propertyName string,
) (classProperty.Property, error) {

	var err error
	var currentProperty classProperty.Property
	var currentPropertyName string

	// Preparations.
	propertyName = strings.ToLower(propertyName)

	// Search.
	for _, currentProperty = range this.properties {
		currentPropertyName = strings.ToLower(currentProperty.GetName())
		if currentPropertyName == propertyName {
			return currentProperty, nil
		}
	}

	// No Luck.
	err = fmt.Errorf(
		ErrFormatMissingPropertyName,
		propertyName,
		this.name,
	)
	err = errorz.Report(ErrorReporter, err)
	return classProperty.Property{}, err
}

// Returns a Property Name by its ID.
func (this Class) GetPropertyNameById(
	propertyId uint,
) (string, error) {

	var err error
	var idExists bool
	var property classProperty.Property
	var result string

	// Get a Property.
	property, idExists = this.properties[propertyId]
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingPropertyId,
			propertyId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	result = property.GetName()

	return result, nil
}

// Checks whether a new Property's Name is suitable.
// This Method is used when checking whether a Property Update is acceptable.
func (this Class) newPropertyNameIsGood(
	existingPropertyId uint,
	newProperty classProperty.Property,
) (bool, error) {

	var err error
	var propertyNameOld string
	var propertyNameNew string
	var propertyOld classProperty.Property

	// Prerations.
	propertyOld, err = this.GetPropertyById(existingPropertyId)
	if err != nil {
		return false, err
	}
	propertyNameOld = propertyOld.GetName()
	propertyNameOld = strings.ToLower(propertyNameOld)
	propertyNameNew = newProperty.GetName()
	propertyNameNew = strings.ToLower(propertyNameNew)

	// Compare.
	if propertyNameNew == propertyNameOld {
		return true, nil
	}

	return false, nil
}

// Checks whether the Object ID is already used in the Class.
func (this Class) ObjectIdIsUsed(
	objectId uint,
) bool {

	var idExists bool

	_, idExists = this.objects[objectId]

	return idExists
}

// Checks whether the Object's Properties are registered in the Class.
// Returns an Error if the Result is negative.
func (this Class) ObjectPropertiesAreRegistered(
	object object.Object,
) (bool, error) {

	var err error
	var objectProperties map[uint]objectProperty.Property
	var objectPropertyId uint

	// Check that Object's Properties are registered in the Class.
	objectProperties = object.GetProperties()
	for objectPropertyId, _ = range objectProperties {
		if !this.PropertyIdIsUsed(objectPropertyId) {
			err = fmt.Errorf(
				ErrFormatUnregisteredObjectPropertyId,
				objectPropertyId,
				this.name,
			)
			err = errorz.Report(ErrorReporter, err)
			return false, err
		}
	}

	return true, nil
}

// Checks whether the Property ID is already used in the Class.
func (this Class) PropertyIdIsUsed(
	propertyId uint,
) bool {

	var idExists bool

	_, idExists = this.properties[propertyId]

	return idExists
}

// Checks whether the Property Name is already used in the Class.
// During the Check, all Property Names are converted into a low Case Text.
func (this Class) PropertyNameIsUsed(
	propertyName string,
) bool {

	var currentPropertyName string
	var currentProperty classProperty.Property
	var propertyNameLC string

	propertyNameLC = strings.ToLower(propertyName)

	for _, currentProperty = range this.properties {
		currentPropertyName = strings.ToLower(currentProperty.GetName())
		if currentPropertyName == propertyNameLC {
			return true
		}
	}

	return false
}

// Sets an Object in the Class.
func (this *Class) setObject(
	objectId uint,
	obj object.Object,
) {

	this.objects[objectId] = obj
}

// Sets a Class Property in the Class.
func (this *Class) setProperty(
	propertyId uint,
	aClassProperty classProperty.Property,
) {

	this.properties[propertyId] = aClassProperty
}

// Updates an Object in the Class.
// The ID must be a real existing ID.
func (this *Class) UpdateObject(
	objectId uint,
	anObject object.Object,
) error {

	var err error
	var idExists bool

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if objectId != anObject.GetId() {
		err = fmt.Errorf(
			common.ErrFormatUpdateIdMismatch,
			objectId,
			anObject.GetId(),
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Objects List Initialization.
	if this.objects == nil {
		this.objects = make(map[uint]object.Object)
	}

	// Is Object ID empty?
	if object.IdIsEmpty(objectId) {
		err = fmt.Errorf(
			ErrFormatBadObjectId,
			objectId,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// ID Exists?
	idExists = this.ObjectIdIsUsed(objectId)
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingObjectId,
			objectId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify the Property's Integrity.
	err = anObject.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Ensure that Object's Properties are registered in the Class.
	_, err = this.ObjectPropertiesAreRegistered(anObject)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Property.
	this.setObject(objectId, anObject)

	return nil
}

// Updates a Property in the Class.
// The ID must be a real existing ID.
func (this *Class) UpdateProperty(
	propertyId uint,
	aProperty classProperty.Property,
) error {

	var err error
	var idExists bool
	var ok bool
	var propertyNameOld string
	var propertyOld classProperty.Property

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
		this.properties = make(map[uint]classProperty.Property)
	}

	// Is a Property ID empty?
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

	// Name of the Property cannot be changed.
	ok, err = this.newPropertyNameIsGood(
		propertyId,
		aProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if !ok {
		// Prepare Data and create an Error.
		propertyOld = this.getProperty(propertyId)
		propertyNameOld = propertyOld.GetName()
		err = fmt.Errorf(
			ErrFormatPropertyNameConflict,
			aProperty.GetName(),
			propertyNameOld,
			this.name,
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

// Verifies the Integrity of a Class.
func (this Class) VerifyIntegrity() error {

	var err error

	// Verify.
	err = this.verifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Verifies the Integrity of a Class.
func (this Class) verifyIntegrity() error {

	var err error

	// Lists Initialization.
	if this.properties == nil {
		this.properties = make(map[uint]classProperty.Property)
	}
	if this.objects == nil {
		this.objects = make(map[uint]object.Object)
	}

	// Check ID.
	if this.id == IdNone {
		err = fmt.Errorf(
			ErrFormatBadId,
			this.id,
		)
		return err
	}

	// Check Name.
	err = this.checkName()
	if err != nil {
		return err
	}

	// Verify Objects Array Index Matching.
	for anObjectId, anObject := range this.objects {
		if anObjectId != anObject.GetId() {
			err = fmt.Errorf(
				ErrFormatDamagedObjectId,
				anObject.GetId(),
				anObjectId,
				this.id,
			)
			return err
		}
	}

	// Verify Properties Array Index Matching.
	for aPropertyId, aProperty := range this.properties {
		if aPropertyId != aProperty.GetId() {
			err = fmt.Errorf(
				ErrFormatDamagedPropertyId,
				aProperty.GetId(),
				aPropertyId,
				this.id,
			)
			return err
		}
	}

	// Verify the Integrity of Objects.
	for _, anObject := range this.objects {
		err = anObject.VerifyIntegrity()
		if err != nil {
			return err
		}
	}

	// Verify the Integrity of Properties.
	for _, aProperty := range this.properties {
		err = aProperty.VerifyIntegrity()
		if err != nil {
			return err
		}
	}

	// Check that all Objects use only registered Class Properties.
	for _, anObject := range this.objects {
		_, err = this.ObjectPropertiesAreRegistered(anObject)
		if err != nil {
			return err
		}
	}

	return nil
}

// Checks whether the specified Class ID is empty.
func IdIsEmpty(
	classId uint,
) bool {

	if classId == IdNone {
		return true
	}

	return false
}

// Creates a new Class.
func New(
	id uint,
	name string,
) Class {

	return Class{
		id:   id,
		name: name,
	}
}
