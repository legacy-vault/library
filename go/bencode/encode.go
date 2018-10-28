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
	var integer uint64
	var keyBA []byte
	var list []interface{}
	var listItem interface{}
	var ok bool
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
		integer, ok = ifc.(uint64)
		if !ok {
			return nil, ErrTypeAssertion
		}
		ba = createIntegerText(integer)

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

// Creates an ASCII Text (Byte Array) of an Integer.
func createIntegerText(value uint64) []byte {

	return []byte(strconv.FormatUint(value, 10))
}
