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

// Collection.go.

package collection

// Collection's Methods.

import (
	"errors"
	"fmt"
	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"math"
	"strings"
)

const ErrorReporter = "Collection"

const ErrFormatBadClassId = "Bad Class ID '%v'"
const ErrFormatDuplicateClassId = "Duplicate Class ID '%v " +
	"(Collection Name is '%v')'"
const ErrFormatMissingClassId = "Missing Class ID '%v' " +
	"(Collection Name is '%v')"
const ErrFormatMissingClassName = "Missing Class Name '%v' " +
	"(Collection Name is '%v')"
const ErrFormatDamagedCollectionClassId = "Damaged Class ID " +
	"(ID='%v', Index='%v') in Collection with Name '%v'"
const ErrNoFreeClassId = "No free Class IDs are left"
const ErrFormatClassNameUsed = "Class Name '%v' is used " +
	"(Collection Name is '%v')"
const ErrFormatClassNameConflict = "Class Update Name Conflict: " +
	"'%v' vs '%v' (Class Name is '%v')"

type Collection struct {

	// Parameters & Settings.
	name string

	// Sub-Items.
	// Associative Array of Classes. Key is ClassId.
	classes map[uint]class.Class
}

// Adds a Class to the Collection.
func (this *Collection) AddClass(
	aClass class.Class,
) error {

	var classId uint
	var className string
	var duplicateId bool
	var duplicateName bool
	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Classes List Initialization.
	if this.classes == nil {
		this.classes = make(map[uint]class.Class)
	}

	// Preparations.
	className = strings.ToLower(aClass.GetName())
	classId = aClass.GetId()

	// Is the Class ID empty?
	if class.IdIsEmpty(classId) {
		err = fmt.Errorf(
			ErrFormatBadClassId,
			classId,
		)
		return err
	}

	// Duplicate ID Check.
	duplicateId = this.ClassIdIsUsed(classId)
	if duplicateId {
		err = fmt.Errorf(
			ErrFormatDuplicateClassId,
			classId,
			this.name,
		)
		return err
	}

	// Check the Class Name.
	err = aClass.CheckName()
	if err != nil {
		return err
	}

	// Check for duplicate Name.
	duplicateName = this.ClassNameIsUsed(className)
	if duplicateName {
		err = fmt.Errorf(
			class.ErrFormatDuplicateClassName,
			className,
		)
		return err
	}

	// Verify the Class's Integrity.
	err = aClass.VerifyIntegrity()
	if err != nil {
		return err
	}

	// Set the Class.
	this.setClass(classId, aClass)

	return nil
}

// Checks whether the specified Class ID is used in the Collection.
func (this Collection) ClassIdIsUsed(
	classId uint,
) bool {

	var idExists bool

	// Duplicate ID Check.
	_, idExists = this.classes[classId]

	return idExists
}

// Checks whether the specified Class Name is used in the Collection.
func (this Collection) ClassNameIsUsed(
	className string,
) bool {

	var aClass class.Class
	var aClassName string

	// Preparations.
	className = strings.ToLower(className)

	// Search.
	for _, aClass = range this.classes {
		aClassName = strings.ToLower(aClass.GetName())
		if className == aClassName {
			return true
		}
	}

	return false
}

// Returns an existing Class of the Collection.
func (this Collection) getClass(
	classId uint,
) class.Class {

	return this.classes[classId]
}

// Returns a Class of the Collection by the Class ID.
func (this Collection) GetClassById(
	classId uint,
) (class.Class, error) {

	var err error
	var idExists bool
	var result class.Class

	// Search.
	result, idExists = this.classes[classId]
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingClassId,
			classId,
			this.name,
		)
		return result, err
	}

	return result, nil
}

// Returns a Class of the Collection by the Class Name.
// During the Search, Class Names are converted into a low Case.
func (this Collection) GetClassByName(
	className string,
) (class.Class, error) {

	var err error
	var currentClass class.Class
	var currentClassName string

	// Preparations.
	className = strings.ToLower(className)

	// Search.
	for _, currentClass = range this.classes {
		currentClassName = strings.ToLower(currentClass.GetName())
		if currentClassName == className {
			return currentClass, nil
		}
	}

	// No Luck.
	err = fmt.Errorf(
		ErrFormatMissingClassName,
		className,
		this.name,
	)
	return class.Class{}, err
}

// Returns the 'classes' Field.
func (this Collection) GetClasses() map[uint]class.Class {

	return this.classes
}

// Returns a Class ID of the Collection by the Class Name.
// During the Search, Class Names are converted into a low Case.
func (this Collection) GetClassIdByName(
	className string,
) (uint, error) {

	var err error
	var classId uint
	var currentClass class.Class
	var currentClassName string

	// Preparations.
	className = strings.ToLower(className)

	// Search.
	for classId, currentClass = range this.classes {
		currentClassName = strings.ToLower(currentClass.GetName())
		if currentClassName == className {
			return classId, nil
		}
	}

	// No Luck.
	err = fmt.Errorf(
		ErrFormatMissingClassName,
		className,
		this.name,
	)
	return class.IdNone, err
}

// Returns the 'name' Field.
func (this Collection) GetName() string {

	return this.name
}

// Returns the next free ID of a Class.
func (this Collection) GetNextFreeClassId() (uint, error) {

	var err error
	var id uint

	id = class.IdFirstAvailable
	if !this.ClassIdIsUsed(id) {
		return id, nil
	}
	for {
		id++
		if !this.ClassIdIsUsed(id) {
			return id, nil
		}
		if id == math.MaxUint64 {
			break
		}
	}

	err = errors.New(ErrNoFreeClassId)
	return class.IdNone, err
}

// Checks whether a new Class's Name is suitable.
// This Method is used when checking whether a Class Update is acceptable.
func (this Collection) newClassNameIsGood(
	existingClassId uint,
	newClass class.Class,
) (bool, error) {

	var err error
	var classNameOld string
	var classNameNew string
	var classOld class.Class

	// Prerations.
	classOld, err = this.GetClassById(existingClassId)
	if err != nil {
		return false, err
	}
	classNameOld = classOld.GetName()
	classNameOld = strings.ToLower(classNameOld)
	classNameNew = newClass.GetName()
	classNameNew = strings.ToLower(classNameNew)

	// Compare.
	if classNameNew == classNameOld {
		return true, nil
	}

	return false, nil
}

// Sets a Class in the Collection.
func (this *Collection) setClass(
	classId uint,
	aClass class.Class,
) {

	this.classes[classId] = aClass
}

// Updates a Class in the Collection.
// The ID must be a real existing ID.
func (this *Collection) UpdateClass(
	classId uint,
	aClass class.Class,
) error {

	var classOld class.Class
	var classNameOld string
	var err error
	var idExists bool
	var ok bool

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if classId != aClass.GetId() {
		err = fmt.Errorf(
			common.ErrFormatUpdateIdMismatch,
			classId,
			aClass.GetId(),
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Classes List Initialization.
	if this.classes == nil {
		this.classes = make(map[uint]class.Class)
	}

	// Is the Class ID empty?
	if class.IdIsEmpty(classId) {
		err = fmt.Errorf(
			ErrFormatBadClassId,
			classId,
		)
		return err
	}

	// ID Exists?
	idExists = this.ClassIdIsUsed(classId)
	if !idExists {
		err = fmt.Errorf(
			ErrFormatMissingClassId,
			classId,
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Name of the Class cannot be changed.
	ok, err = this.newClassNameIsGood(
		classId,
		aClass,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if !ok {
		// Prepare Data and create an Error.
		classOld = this.getClass(classId)
		classNameOld = classOld.GetName()
		err = fmt.Errorf(
			ErrFormatClassNameConflict,
			aClass.GetName(),
			classNameOld,
			this.name,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Verify the Class's Integrity.
	err = aClass.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Set the Class.
	this.setClass(classId, aClass)

	return nil
}

// Verifies the Integrity of a Collection.
func (this Collection) VerifyIntegrity() error {

	var err error

	// Verify.
	err = this.verifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Verifies the Integrity of a Collection.
func (this Collection) verifyIntegrity() error {

	var err error

	// Classes List Initialization.
	if this.classes == nil {
		this.classes = make(map[uint]class.Class)
	}

	// Verify Classes Array Index Matching.
	for aClassId, aClass := range this.classes {
		if aClassId != aClass.GetId() {
			err = fmt.Errorf(
				ErrFormatDamagedCollectionClassId,
				aClass.GetId(),
				aClassId,
				this.name,
			)
			return err
		}
	}

	// Verify the Integrity of Collections's Classes.
	for _, aClass := range this.classes {
		err = aClass.VerifyIntegrity()
		if err != nil {
			return err
		}
	}

	return nil
}

// Creates a new Collection.
func New(
	collectionName string,
) Collection {

	var result Collection

	// Set.
	result.name = collectionName

	return result
}
