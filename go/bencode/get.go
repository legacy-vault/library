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

// get.go.

// 'Bencode' Encoding Reader-Functions.

// Last Update Time: 2018-10-30.

package bencode

import (
	"bufio"
	"bytes"
	"fmt"
)

// Gets a Byte String from the Stream (Reader).
func getByteString(
	reader *bufio.Reader,
	byteStringMaxLen int,
) ([]byte, error) {

	var b byte
	var ba []byte
	var bytesAccumulator *bytes.Buffer
	var byteStringLen uint64
	var err error
	var errorArea []byte
	var i uint64
	var sizeHeader []byte

	// Prepare Data.
	ba = []byte{}
	sizeHeader = []byte{}

	// Read the first Byte.
	b, err = reader.ReadByte()
	if err != nil {
		return ba, err
	}

	for b != HeaderStringSizeValueDelimiter {

		// Syntax Check.
		if !byteIsNonNegativeASCIINumeric(b) {
			errorArea = append(sizeHeader, []byte{b}...)
			err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
			return ba, err
		}

		// Save Byte to Size Header.
		if len(sizeHeader) < byteStringMaxLen {
			sizeHeader = append(sizeHeader, b)
		} else {
			// The Length Header is too big!
			return ba, ErrHeaderLength
		}

		// Read next Byte.
		b, err = reader.ReadByte()
		if err != nil {
			return ba, err
		}
	}

	// We have read the Size Header.
	// Check that it is not empty.
	if len(sizeHeader) == 0 {
		errorArea = append(sizeHeader, []byte{b}...)
		err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
		return ba, err
	}

	// Convert Size Header into normal integer Size Value.
	byteStringLen, err = byteStringToNonNegativeInteger(sizeHeader)
	if err != nil {
		return ba, err
	}

	// Now we should read the Byte String.
	i = 0
	bytesAccumulator = bytes.NewBuffer([]byte{})
	for i < byteStringLen {

		// Read next Symbol.
		b, err = reader.ReadByte()
		if err != nil {
			return ba, err
		}

		// Save the Byte to the Accumulator.
		err = bytesAccumulator.WriteByte(b)
		if err != nil {
			return ba, err
		}

		i++
	}

	return bytesAccumulator.Bytes(), nil
}

// Gets a Dictionary.
// We suppose that the Header of Dictionary ('d')
// has already been read from the Stream.
func getDictionary(reader *bufio.Reader) (interface{}, error) {

	var b byte
	var dict []DictionaryItem
	var dictKey []byte
	var dictValue interface{}
	var err error

	// Prepare Data.
	dict = make([]DictionaryItem, 0)

	// Probe the Next Byte to check the End of Dictionary.
	b, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}
	for b != FooterCommon {

		// That single Byte (we probed) was not an End!
		// We must get back, rewind that Byte.
		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Get the Key.
		dictKey, err = getDictionaryKey(reader)
		if err != nil {
			return nil, err
		}

		// Get the Value.
		dictValue, err = getDictionaryValue(reader)
		if err != nil {
			return nil, err
		}

		// Save Item into Dictionary.
		dict = append(
			dict,
			DictionaryItem{
				// System Fields.
				Key:   dictKey,
				Value: dictValue,

				// Additional Fields for special Purposes.
				KeyStr:   string(dictKey),
				ValueStr: interfaceToString(dictValue),
			},
		)

		// Probe the Next Byte to check the End of Dictionary.
		b, err = reader.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}

// Gets Dictionary's Key.
func getDictionaryKey(reader *bufio.Reader) ([]byte, error) {

	return getByteString(reader, ByteStringMaxLength)
}

// Gets Dictionary's Value.
func getDictionaryValue(reader *bufio.Reader) (interface{}, error) {

	return getBencodedValue(reader)
}

// Gets an Integer from the Stream (Reader).
// We suppose that the Header of Integer ('i')
// has already been read from the Stream.
func getInteger(
	reader *bufio.Reader,
	integerMaxLen int,
) (int64, error) {

	var b byte
	var err error
	var errorArea []byte
	var value int64
	var valueBA []byte

	// Prepare Data.
	valueBA = []byte{}

	// Read the first Byte.
	b, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	for b != FooterCommon {

		// Syntax Check.
		if !byteIsASCIINumeric(b) {
			errorArea = append(valueBA, []byte{b}...)
			err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
			return 0, err
		}

		// Save Byte to Value Byte Array.
		if len(valueBA) < integerMaxLen {
			valueBA = append(valueBA, b)
		} else {
			// The Integer is too big!
			return 0, ErrIntegerLength
		}

		// Read next Byte.
		b, err = reader.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// We have read the Value.
	// Check that it is not empty.
	if len(valueBA) == 0 {
		errorArea = append(valueBA, []byte{b}...)
		err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
		return 0, err
	}

	// Convert Value into normal integer Value.
	value, err = byteStringToInteger(valueBA)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// Gets a List from the Stream (Reader).
// We suppose that the Header of List ('l')
// has already been read from the Stream.
func getList(reader *bufio.Reader) ([]interface{}, error) {

	var b byte
	var err error
	var list []interface{}
	var listItem interface{}

	// Prepare Data.
	list = make([]interface{}, 0)

	// Probe the Next Byte to check the End of List.
	b, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}
	for b != FooterCommon {

		// That single Byte (we probed) was not an End!
		// We must get back, rewind that Byte.
		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Get the Item.
		listItem, err = getBencodedValue(reader)
		if err != nil {
			return nil, err
		}

		// Save Item into Dictionary.
		list = append(list, listItem)

		// Probe the Next Byte to check the End of List.
		b, err = reader.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

// Gets a raw "bencoded" Value, including its Sub-values.
func getBencodedValue(reader *bufio.Reader) (interface{}, error) {

	var b byte
	var err error
	var errorArea []byte
	var result interface{}
	var resultByteString []byte
	var resultInteger int64

	// Get the first Byte from Stream to know its Type.
	b, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}

	// Analyze the Type.
	if b == HeaderDictionary {

		// Dictionary.
		result, err = getDictionary(reader)
		if err != nil {
			return nil, err
		}

		return result, nil

	} else if b == HeaderList {

		// List.
		result, err = getList(reader)
		if err != nil {
			return nil, err
		}

		return result, nil

	} else if b == HeaderInteger {

		// Integer.
		resultInteger, err = getInteger(reader, IntegerMaxLength)
		if err != nil {
			return nil, err
		}

		return resultInteger, nil

	} else if byteIsNonNegativeASCIINumeric(b) {

		// It must be an ASCII Number indicating a Byte String.
		// => Byte String.

		// Rewind the Cursor back, as it does not have a Type-Prefix!
		// The 'bencode' Encoding is ugly...
		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Read the Byte String.
		resultByteString, err = getByteString(reader, ByteStringMaxLength)
		if err != nil {
			return nil, err
		}

		return resultByteString, nil

	} else {

		// Otherwise, it is a Syntax Error.
		errorArea = []byte{b}
		err = fmt.Errorf(ErrFmtSyntaxErrorAt, errorArea)
		return nil, err
	}

	// This Code is unreachable.
	// And, yes, I do believe in Anomalies! ;-)
	return nil, ErrAnomaly
}
