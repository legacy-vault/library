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

// Database Collection.

// Collection Types and Structures used in the Database.

import (
	"fmt"
	"strconv"
	"time"

	"xxx/Common"
	"xxx/Errorz"
	"xxx/Uc/Collection/Class/Property/KindSettings"
)

const ErrorReporter = "Collection"

const ErrFormatTypeAssertion = "Error during Type Casting to '%v'. " +
	"Value='%v'"
const ErrFormatTypeUnsupported = "Unsupported Value Type: '%v'"

type DatabaseClass struct {
	Id   uint
	Name string
}

type DatabaseClassObject struct {
	ClassId uint
	Id      uint
}

type DatabaseClassObjectProperty struct {
	PropertyId                uint
	ObjectId                  uint
	ClassId                   uint
	Value                     interface{}
	ValueIsSet                bool
	ValueDbTypeFromDriver     string
	ValueDbTypeFromCollection kindsettings.KindSettings
}

type DatabaseClassProperty struct {
	ClassId     uint
	DbType      string
	Description string
	Id          uint
	Name        string
}

// Converts the raw Database Object Property Value into something more usable.
func (this *DatabaseClassObjectProperty) ConvertValueDbToUsable() error {

	var err error
	var tFloat64 float64
	var tInt64 int64
	var tTime time.Time
	var tUint64 uint64
	var valueBytes []byte
	var valueString string

	// Convert an Interface of Bytes to Bytes.
	valueBytes, err = common.SqlResponseInterfaceToBytes(this.Value)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	valueString = string(valueBytes)

	switch this.ValueDbTypeFromCollection.InternalType {

	case kindsettings.InternalTypeString:

		this.Value = valueString

	case kindsettings.InternalTypeInt:

		tInt64, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = int(tInt64)

	case kindsettings.InternalTypeInt64:

		tInt64, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = tInt64

	case kindsettings.InternalTypeInt32:

		tInt64, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = int32(tInt64)

	case kindsettings.InternalTypeInt16:

		tInt64, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = int16(tInt64)

	case kindsettings.InternalTypeInt8:

		tInt64, err = strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = int8(tInt64)

	case kindsettings.InternalTypeUint:

		tUint64, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = uint(tUint64)

	case kindsettings.InternalTypeUint64:

		tUint64, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = tUint64

	case kindsettings.InternalTypeUint32:

		tUint64, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = uint32(tUint64)

	case kindsettings.InternalTypeUint16:

		tUint64, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = uint16(tUint64)

	case kindsettings.InternalTypeUint8:

		tUint64, err = strconv.ParseUint(valueString, 10, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = uint8(tUint64)

	case kindsettings.InternalTypeFloat:

		tFloat64, err = strconv.ParseFloat(valueString, 64)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = tFloat64

	case kindsettings.InternalTypeTime:

		tTime, err = time.Parse(kindsettings.TimeFormatA, valueString)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.Value = tTime

	default:
		err = fmt.Errorf(
			ErrFormatTypeUnsupported,
			this.ValueDbTypeFromCollection.InternalType,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}
