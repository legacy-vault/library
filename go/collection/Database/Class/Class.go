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

// Database Class (Type).
// Has Nothing to do with Collection Class at all.
// Such a Mess is due to Golang forbidding to use the 'type' as a Type Name.
// The good old 'C'-Family Languages (C, C++, C#, Java) do not have this ugly
// Thing.

import (
	"errors"
	"fmt"
	"strings"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Database/Class/Value"
	"github.com/legacy-vault/library/go/collection/Errorz"
)

// Notes:
// The Word 'Class' is used due to the Restrictions of the 'Go' Programming
// Language, the Word 'Type' is reserved for the Language itself.

const ErrorReporter = "Class"

const ErrFormatClassValueAliasUnknown = "Unsupported Database Class: '%v'."
const ErrNoClass = "Class is not set"

type Class struct {
	value value.Value
}

// Gets an Alias of the internal Value of the Database Class (Type).
func (this Class) GetAlias() string {

	switch this.value {

	case value.Mysql:
		return value.MysqlAlias

	default:
		return value.NoneAlias
	}
}

// Gets an internal Value of the Database Class (Type).
func (this Class) GetValue() value.Value {

	return this.value
}

// Returns the 'configured' State.
func (this Class) IsConfigured() bool {

	if this.value != value.None {
		return true
	}

	return false
}

// Sets the Database Class (Type) internal Value by its Alias.
func (this *Class) Set(
	valueAlias string,
) error {

	var err error
	var valueAliasLC string

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	valueAliasLC = strings.ToLower(valueAlias)
	switch valueAliasLC {

	case value.MysqlAliasLc:
		this.value = value.Mysql

	default:
		this.value = value.None
		err = fmt.Errorf(
			ErrFormatClassValueAliasUnknown,
			valueAlias,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}
