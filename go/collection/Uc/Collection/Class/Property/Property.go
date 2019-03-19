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

// Property.go.

package property

// Collection Class Property.

import (
	"fmt"
	"time"

	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/Kind"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
	"github.com/legacy-vault/library/go/collection/Utf8"
)

// Notes:
//
//	1.	Property Name must contain only small ASCII Letters.
//	2.	Property Name is used in Database Table Names.

const ErrorReporter = "ClassProperty"

const ErrFormatBadName = "Bad Name: '%v'"

const NameLengthMax = 255 // VARCHAR(255).
const ErrFormatBadId = "Bad ID '%v'"
const ErrFormatBadDescription = "Bad Description '%v'"

const IdNone = 0
const IdFirstAvailable = IdNone + 1

// Acts as a Reference for a Property.
// Id est, it has all the Fields except the Value.
// Value is stored in an Object's Property.
type Property struct {

	// Parameters & Settings.
	description string

	// Property ID is unique among Class Properties.
	// Id est: Properties of different Classes may have the same ID,
	// but Properties of the same Class must have unique IDs.
	id uint

	// Property Type.
	kind kind.Kind

	// Property Name is unique among Class Properties.
	// Id est: Properties of different Classes may have the same Name,
	// but Properties of the same Class must have unique Names.
	// Names are converted to a low Case when checked for Uniqueness.
	name string

	timeOfCreation time.Time
	timeOfUpdate   time.Time
}

// Checks Class Name.
func (this Property) CheckName() error {

	var err error

	err = this.checkName()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Checks Class Name.
func (this Property) checkName() error {

	var err error

	// Check Length.
	if len(this.name) > NameLengthMax {
		err = fmt.Errorf(
			ErrFormatBadName,
			this.name,
		)
		return err
	}

	// Check Symbols.
	if utf8.StringHasOnlyASCIILetterOrNumber(this.name) != true {
		err = fmt.Errorf(
			ErrFormatBadName,
			this.name,
		)
		return err
	}

	return nil
}

// Returns the 'description' Field.
func (this Property) GetDescription() string {

	return this.description
}

// Returns the 'id' Field.
func (this Property) GetId() uint {

	return this.id
}

// Returns the Property Type (Kind).
func (this Property) GetKind() kindsettings.KindSettings {

	return this.getKind()
}

// Returns the Property Type (Kind).
func (this Property) getKind() kindsettings.KindSettings {

	return this.kind.Get()
}

// Returns the 'name' Field.
func (this Property) GetName() string {

	return this.name
}

// Returns the Property Type (Kind).
func (this Property) GetType() kindsettings.KindSettings {

	return this.getKind()
}

// Verifies the Integrity of a Property.
func (this Property) VerifyIntegrity() error {

	var err error

	// Check ID.
	if this.id == IdNone {
		err = fmt.Errorf(
			ErrFormatBadId,
			this.id,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check Description.
	if len(this.description) == 0 {
		err = fmt.Errorf(
			ErrFormatBadDescription,
			this.description,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check Kind (Type).
	err = this.kind.VerifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check Name.
	err = this.checkName()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Checks whether the specified Class Property ID is empty.
func IdIsEmpty(
	classPropertyId uint,
) bool {

	if classPropertyId == IdNone {
		return true
	}

	return false
}

// Returns a Reference Property to be used in the Class Property List.
// A Reference Property has no Value.
func NewReferenceProperty(
	propertyId uint,
	propertyName string,
	propertyDescription string,
	propertyKindSettings kindsettings.KindSettings,
) (Property, error) {

	var err error
	var result Property
	var resultKind kind.Kind

	err = resultKind.Set(propertyKindSettings)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	result = Property{
		description: propertyDescription,
		id:          propertyId,
		kind:        resultKind,
		name:        propertyName,
	}

	return result, nil
}
