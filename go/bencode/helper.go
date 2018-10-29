//============================================================================//
//
// Copyright © 2018 by McArcher.
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
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-29.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// helper.go.

// Helping Functions.

// Last Update Time: 2018-10-29.

package bencode

import (
	"math"
	"reflect"
	"strconv"
)

const ArchitectureIs64Bit bool = (strconv.IntSize == 64)
const ArchitectureIs32Bit bool = (strconv.IntSize == 32)

// Checks whether the Byte is ASCII numeric Symbol.
func byteIsASCIINumeric(b byte) bool {

	if (b == '0') ||
		(b == '1') ||
		(b == '2') ||
		(b == '3') ||
		(b == '4') ||
		(b == '5') ||
		(b == '6') ||
		(b == '7') ||
		(b == '8') ||
		(b == '9') {
		return true
	}

	return false
}

// Converts a Byte String into an unsigned 64-Bit Integer.
func byteStringToInteger(ba []byte) (uint64, error) {

	var err error
	var result uint64
	var str string

	if len(ba) > ByteStringMaxLength {
		return 0, ErrByteStringToInt
	}

	str = string(ba)
	result, err = strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Checks whether the unsigned Integer is able to be converted into 'int' Type.
func uint64IsConvertibleToInt(x uint64) bool {

	// 64-bit CPU Architecture.
	if ArchitectureIs64Bit {

		if x <= math.MaxInt64 {
			return true
		}
		return false
	}

	// 32-bit CPU Architecture.
	if ArchitectureIs32Bit {

		if x <= math.MaxInt32 {
			return true
		}
		return false
	}

	// UnKnown CPU Architecture.
	return false
}

// Tries to get a textual Data from an Interface.
func interfaceToString(i interface{}) string {

	var ba []byte
	var iType reflect.Kind
	var iElementType reflect.Kind
	var ok bool

	// Slice?
	iType = reflect.TypeOf(i).Kind()
	if iType == reflect.Slice {

		// Array Item's Type is Byte?
		iElementType = reflect.TypeOf(i).Elem().Kind()
		if iElementType == reflect.Uint8 {
			ba, ok = i.([]byte)
			if !ok {
				return ""
			}
			return string(ba)
		}
	}

	return ""
}
