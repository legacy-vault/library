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

// KindSettings.go.

package kindsettings

import (
	"fmt"
	"strings"

	"xxx/Errorz"
)

// Collection Class Object Property Kind Setter.

const ErrorReporter = "KindSettings"

// Database Type & Type Prefix.
const DbTypePrefixVarchar = "VARCHAR("
const DbTypeChar = "CHAR"
const DbTypeText = "TEXT"
const DbTypeTextTiny = "TINYTEXT"
const DbTypeTextMedium = "MEDIUMTEXT"
const DbTypeTextLong = "LONGTEXT"
const DbTypeDate = "DATE"
const DbTypeTime = "TIME"
const DbTypeDatetime = "DATETIME"
const DbTypeUnsigned = "UNSIGNED"
const DbTypeDecimal = "DECIMAL"
const DbTypeNumeric = "NUMERIC"
const DbTypeFloat = "FLOAT"
const DbTypeReal = "REAL"
const DbTypeDouble = "DOUBLE"
const DbTypeInteger = "INTEGER"
const DbTypeIntTiny = "TINYINT"
const DbTypeIntSmall = "SMALLINT"
const DbTypeIntMedium = "MEDIUMINT"
const DbTypeInt = "INT"
const DbTypeIntBig = "BIGINT"

// Internal Type.
const InternalTypeString = "string"
const InternalTypeTime = "time" // -> Time Format A.
const InternalTypeInt = "int"
const InternalTypeInt8 = "int8"
const InternalTypeInt16 = "int16"
const InternalTypeInt32 = "int32"
const InternalTypeInt64 = "int64"
const InternalTypeUint = "uint"
const InternalTypeUint8 = "uint8"
const InternalTypeUint16 = "uint16"
const InternalTypeUint32 = "uint32"
const InternalTypeUint64 = "uint64"
const InternalTypeFloat = "float" // -> float64.

const ErrFormatTypeUnsupported = "Unsupported Type '%v'"

const TimeFormatA = "2006-01-02 15:04:05"

type KindSettings struct {
	// Raw Type used by the Database.
	// E.g.: 'VARCHAR(255)'.
	DbType string

	// Alias of an internal Type used by this Program.
	// E.g.: 'time-1' or 'uint32'.
	InternalType string
}

// Checks the KindSettings Object.
func (this KindSettings) IsValid() bool {

	// Check the internal Type.
	switch this.InternalType {

	case InternalTypeString:
	case InternalTypeTime:
	case InternalTypeInt:
	case InternalTypeInt8:
	case InternalTypeInt16:
	case InternalTypeInt32:
	case InternalTypeInt64:
	case InternalTypeUint:
	case InternalTypeUint8:
	case InternalTypeUint16:
	case InternalTypeUint32:
	case InternalTypeUint64:
	case InternalTypeFloat:

	default:
		return false
	}

	return true
}

// Creates a KindSettings Object using the DbType Field.
func NewWithDbType(
	dbType string,
) (KindSettings, error) {

	var err error
	var result KindSettings

	dbType = strings.ToUpper(dbType)
	result.DbType = dbType

	// Textual Types?
	if strings.HasPrefix(dbType, DbTypePrefixVarchar) {
		result.InternalType = InternalTypeString
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeChar) {
		result.InternalType = InternalTypeString
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeText) {
		result.InternalType = InternalTypeString
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeTextTiny) {
		result.InternalType = InternalTypeString
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeTextMedium) {
		result.InternalType = InternalTypeString
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeTextLong) {
		result.InternalType = InternalTypeString
		return result, nil
	}

	// Date-Time Types?
	if strings.HasPrefix(dbType, DbTypeDate) {
		result.InternalType = InternalTypeTime
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeTime) {
		result.InternalType = InternalTypeTime
		return result, nil
	}
	if strings.HasPrefix(dbType, DbTypeDatetime) {
		result.InternalType = InternalTypeTime
		return result, nil
	}

	// Numeric Types?
	if strings.Contains(dbType, DbTypeUnsigned) {

		// Unsigned Number.
		if strings.HasPrefix(dbType, DbTypeInteger) {
			result.InternalType = InternalTypeUint
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntTiny) {
			result.InternalType = InternalTypeUint8
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntSmall) {
			result.InternalType = InternalTypeUint16
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntMedium) {
			result.InternalType = InternalTypeUint32
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeInt) {
			result.InternalType = InternalTypeUint32
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntBig) {
			result.InternalType = InternalTypeUint64
			return result, nil
		}

	} else {

		// Signed Number.
		if strings.HasPrefix(dbType, DbTypeInteger) {
			result.InternalType = InternalTypeInt
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntTiny) {
			result.InternalType = InternalTypeInt8
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntSmall) {
			result.InternalType = InternalTypeInt16
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntMedium) {
			result.InternalType = InternalTypeInt32
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeInt) {
			result.InternalType = InternalTypeInt32
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeIntBig) {
			result.InternalType = InternalTypeInt64
			return result, nil
		}

		// Floating Point Types.
		if strings.HasPrefix(dbType, DbTypeDecimal) {
			result.InternalType = InternalTypeFloat
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeNumeric) {
			result.InternalType = InternalTypeFloat
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeFloat) {
			result.InternalType = InternalTypeFloat
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeReal) {
			result.InternalType = InternalTypeFloat
			return result, nil
		}
		if strings.HasPrefix(dbType, DbTypeDouble) {
			result.InternalType = InternalTypeFloat
			return result, nil
		}
	}

	// Nothing has been found.
	err = fmt.Errorf(
		ErrFormatTypeUnsupported,
		dbType,
	)
	err = errorz.Report(ErrorReporter, err)
	return result, err
}
