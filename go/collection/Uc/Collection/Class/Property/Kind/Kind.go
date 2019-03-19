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

// Kind.go.

package kind

// Collection Class Object Property Kind.

import (
	"fmt"

	"github.com/legacy-vault/library/go/collection/Database/StringLiteral"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
)

const ErrorReporter = "Kind"

const ErrFormatTypeDb = "Bad Type for Database: %v"
const ErrFormatTypeInternal = "Bad internal Type: %v"

type Kind struct {

	// Raw Type used by the Database.
	// E.g.: 'VARCHAR(255)'.
	dbType string

	// Alias of an internal Type used by this Program.
	// E.g.: 'time-1' or 'uint32'.
	internalType string
}

// Returns the Type and its Selector.
func (this Kind) Get() kindsettings.KindSettings {

	return kindsettings.KindSettings{
		DbType:       this.dbType,
		InternalType: this.internalType,
	}
}

// Sets the Kind.
func (this *Kind) set(
	aKind Kind,
) {
	// Set.
	this.dbType = aKind.dbType
	this.internalType = aKind.internalType
}

// Sets the Kind.
func (this *Kind) Set(
	setter kindsettings.KindSettings,
) error {

	var err error
	var aKind Kind
	var sl stringliteral.StringLiteral

	// Preparation.
	aKind = Kind{
		dbType:       setter.DbType,
		internalType: setter.InternalType,
	}

	// Check the Kind.
	err = aKind.verifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Filter out the unwanted Symbols.
	sl.Value = setter.DbType
	this.dbType = sl.Escaped()

	// Set.
	this.set(aKind)

	return nil
}

// Verifies the Integrity of a Kind (Type).
func (this Kind) VerifyIntegrity() error {

	var err error

	// Verify Integrity.
	err = this.verifyIntegrity()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Verifies the Integrity of a Kind (Type).
func (this Kind) verifyIntegrity() error {

	var err error

	// Check the Values.
	if len(this.dbType) == 0 {
		err = fmt.Errorf(
			ErrFormatTypeDb,
			this.dbType,
		)
		return err
	}
	if len(this.internalType) == 0 {
		err = fmt.Errorf(
			ErrFormatTypeInternal,
			this.internalType,
		)
		return err
	}

	return nil
}
