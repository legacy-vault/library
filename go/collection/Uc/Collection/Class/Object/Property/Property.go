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

// Collection Class Object Property.

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/legacy-vault/library/go/collection/Database/StringLiteral"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property/Value"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
)

const ErrorReporter = "ObjectProperty"

const ErrFormatTypeUnsupported = "Unsupported Type '%v'"
const ErrFormatTypeAssertion = "Type Assertion Error. " +
	"Interface='%v' Type='%v' ReflectKind='%v'"

const (
	TypeUint32 = "uint32"
	TypeString = "string"
)

const ErrFormatBadId = "Bad ID '%v'"

const IdNone = 0
const IdFirstAvailable = IdNone + 1

const StringTypeBorder = "'"

// Stores the Value and Property ID.
// Property ID is a Link to a Reference Property.
type Property struct {

	// Property ID is unique among Class Properties.
	// Id est: Properties of different Classes may have the same ID,
	// but Properties of the same Class must have unique IDs.
	id uint

	timeOfCreation time.Time
	timeOfUpdate   time.Time

	// Sub-Items.
	value value.Value
}

// Returns the 'id' Field.
func (this Property) GetId() uint {

	return this.id
}

// Returns the 'value' Field.
func (this Property) GetValue() value.Value {

	return this.value
}

// Returns the 'value' Field formatted for Usage in Database.
func (this Property) GetValueString(
	propertyKind kindsettings.KindSettings,
) (string, error) {

	var err error
	var ok bool
	var result string
	var resultUint32 uint32
	var sl stringliteral.StringLiteral

	// Analyze the Type.
	switch propertyKind.InternalType {

	case TypeUint32:
		resultUint32, ok = (this.value).(uint32)
		if !ok {
			err = fmt.Errorf(
				ErrFormatTypeAssertion,
				this.value,
				propertyKind.InternalType,
				reflect.TypeOf(this.value).Kind(),
			)
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}
		result = strconv.FormatUint(uint64(resultUint32), 10)

	case TypeString:
		result, ok = (this.value).(string)
		if !ok {
			err = fmt.Errorf(
				ErrFormatTypeAssertion,
				this.value,
				propertyKind.InternalType,
				reflect.TypeOf(this.value).Kind(),
			)
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}
		// Escape special Symbols and put into single Quotes.
		sl.Value = result
		result = sl.EscapedAndSingleQuoted()

	default:
		err = fmt.Errorf(
			ErrFormatTypeUnsupported,
			propertyKind.InternalType,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, nil
}

// Sets the Value of the Property.
func (this *Property) SetValue(
	value value.Value) {

	this.value = value
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

	return nil
}

// Checks whether the specified Object Property ID is empty.
func IdIsEmpty(
	objectPropertyId uint,
) bool {

	if objectPropertyId == IdNone {
		return true
	}

	return false
}

// Returns a Property to be used in the Object Property List.
func NewObjectProperty(
	propertyId uint,
	propertyValue value.Value,
) Property {

	var result Property

	result = Property{
		id:    propertyId,
		value: propertyValue,
	}

	return result
}
