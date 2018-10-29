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

// encode.go.

// 'Bencode' Encoding Functions.

// Last Update Time: 2018-10-29.

package bencode

import (
	"reflect"
	"strconv"
)

// Cached Prefixes.
var prefixInteger = []byte{HeaderInteger}

// Encodes an Interface into an Array of Bytes.
func Encode(ifc interface{}) ([]byte, error) {

	return encodeInterface(ifc)
}

// Encodes an Interface.
func encodeInterface(ifc interface{}) ([]byte, error) {

	var ba []byte
	var dictionary []DictionaryItem
	var dictItem DictionaryItem
	var err error
	var ifcType reflect.Kind
	var ifcElementType reflect.Kind
	var intVar int
	var int8var int8
	var int16var int16
	var int32var int32
	var int64var int64
	var keyBA []byte
	var list []interface{}
	var listItem interface{}
	var ok bool
	var stringVar string
	var uint8var uint8
	var uint16var uint16
	var uint32var uint32
	var uint64var uint64
	var uintVar uint
	var valueBA []byte

	// Check an Interface's Type.
	ifcType = reflect.TypeOf(ifc).Kind()

	// Array?
	if ifcType == reflect.Slice {

		// Get Type of Sub-Elements.
		ifcElementType = reflect.TypeOf(ifc).Elem().Kind()
		if ifcElementType == reflect.Uint8 {

			// Array of Bytes.
			// => 'bencode' Byte String.

			// Convert the Type.
			ba, ok = ifc.([]byte)
			if !ok {
				return nil, ErrTypeAssertion
			}

			// Add Prefixes and Postfixes to the Byte String.
			ba = append(createSizePrefix(uint64(len(ba))), ba...)

			return ba, nil
		}

		// Try to change Type to Dictionary.
		dictionary, ok = ifc.([]DictionaryItem)
		if ok {

			// => 'bencode' Dirctionary.

			// Dictionary Prefix.
			ba = []byte{HeaderDictionary}

			// Add Keys and Values.
			for _, dictItem = range dictionary {

				// Add Key.
				keyBA, err = encodeInterface(dictItem.Key)
				if err != nil {
					return nil, err
				}
				ba = append(ba, keyBA...)

				// Add Value.
				valueBA, err = encodeInterface(dictItem.Value)
				if err != nil {
					return nil, err
				}
				ba = append(ba, valueBA...)
			}

			// Dictionary Postfix.
			ba = append(ba, FooterCommon)

			return ba, nil
		}

		// Try to change Type to List.
		list, ok = ifc.([]interface{})
		if ok {

			// => 'bencode' List.

			// Dictionary Prefix.
			ba = []byte{HeaderList}

			// Add Values.
			for _, listItem = range list {

				// Add Value.
				valueBA, err = encodeInterface(listItem)
				if err != nil {
					return nil, err
				}
				ba = append(ba, valueBA...)
			}

			// Dictionary Postfix.
			ba = append(ba, FooterCommon)

			return ba, nil
		}

		// Unknown Type.
		return nil, ErrDataType

	} else if ifcType == reflect.Uint64 {

		// => 'bencode' Integer.

		// Convert the Type.
		uint64var, ok = ifc.(uint64)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createUIntegerText(uint64var)

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Int64 {

		// => 'bencode' Integer.

		// Convert the Type.
		int64var, ok = ifc.(int64)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(int64var)

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Int {

		// => 'bencode' Integer.

		// Convert the Type.
		intVar, ok = ifc.(int)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(int64(intVar))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Uint {

		// => 'bencode' Integer.

		// Convert the Type.
		uintVar, ok = ifc.(uint)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createUIntegerText(uint64(uintVar))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.String {

		// String.
		// => 'bencode' Byte String.

		// Convert the Type.
		stringVar, ok = ifc.(string)
		if !ok {
			return nil, ErrTypeAssertion
		}

		// String → Byte Array.
		ba = []byte(stringVar)

		// Add Prefixes and Postfixes to the Byte String.
		ba = append(createSizePrefix(uint64(len(ba))), ba...)

		return ba, nil

	} else if ifcType == reflect.Uint32 {

		// => 'bencode' Integer.

		// Convert the Type.
		uint32var, ok = ifc.(uint32)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createUIntegerText(uint64(uint32var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Int32 {

		// => 'bencode' Integer.

		// Convert the Type.
		int32var, ok = ifc.(int32)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(int64(int32var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Uint16 {

		// => 'bencode' Integer.

		// Convert the Type.
		uint16var, ok = ifc.(uint16)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createUIntegerText(uint64(uint16var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Int16 {

		// => 'bencode' Integer.

		// Convert the Type.
		int16var, ok = ifc.(int16)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(int64(int16var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Uint8 {

		// => 'bencode' Integer.

		// Convert the Type.
		uint8var, ok = ifc.(uint8)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createUIntegerText(uint64(uint8var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	} else if ifcType == reflect.Int8 {

		// => 'bencode' Integer.

		// Convert the Type.
		int8var, ok = ifc.(int8)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(int64(int8var))

		// Add Prefixes and Postfixes to the Integer.
		ba = append(prefixInteger, ba...)
		ba = append(ba, FooterCommon)

		return ba, nil

	}

	// Unknown Type.
	return nil, ErrDataType
}

// Creates a Size Prefix with a Delimiter.
func createSizePrefix(size uint64) []byte {

	return append(
		[]byte(strconv.FormatUint(size, 10)),
		HeaderStringSizeValueDelimiter,
	)
}

// Creates an ASCII Text (Byte Array) of an unsigned Integer.
func createUIntegerText(value uint64) []byte {

	return []byte(strconv.FormatUint(value, 10))
}

// Creates an ASCII Text (Byte Array) of a signed Integer.
func createIntegerText(value int64) []byte {

	return []byte(strconv.FormatInt(value, 10))
}
