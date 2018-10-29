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

// bencode.go.

// Common 'Bencode' Encoding related Data.

// Last Update Time: 2018-10-29.

package bencode

import (
	"errors"
)

// Special Symbols of 'bencode' Encoding.

// Headers.
const HeaderDictionary byte = 'd'
const HeaderInteger byte = 'i'
const HeaderList byte = 'l'
const HeaderStringSizeValueDelimiter byte = ':'

// Footers.
const FooterCommon byte = 'e'

// Sections of a BitTorrent File.
const FileSectionAnnounce = "announce"
const FileSectionAnnounceList = "announce-list"
const FileSectionCreationDate = "creation date"
const FileSectionComment = "comment"
const FileSectionCreatedBy = "created by"
const FileSectionEncoding = "encoding"
const FileSectionInfo = "info"

// Error Messages.
const ErrAnomalyError = "The Anomaly"
const ErrByteStringToIntError = "Byte String to Integer Conversion Error"
const ErrDataTypeError = "Unsupported Type"
const ErrHeaderLengthError = "The Length Header is too big"
const ErrIntegerLengthError = "The Integeris too big"
const ErrSectionDoesNotExistError = "Section does not exist"
const ErrSelfCheckError = "Self-Check Error"
const ErrTypeAssertionError = "Type Assertion Error"

// Formats of Error Messages.
const ErrFmtSyntaxErrorAt = "Syntax Error At: '%v'."

// Limitations...

// 1. Integer Size (Number of ASCII Letters allowed).
// Maximum Value of UInt64 is '18446744073709551615'.
const IntegerMaxLength = 20

// 2. Byte String Size Header (Number of ASCII Letters allowed).
// We are not going to read Byte Strings which have Length more than that.
const ByteStringMaxLength = IntegerMaxLength

type DictionaryItem struct {
	// System Fields.
	Key   []byte
	Value interface{}

	// Additional Fields for special Purposes.
	KeyStr   string
	ValueStr string
}

type DecodedObject struct {
	FilePath        string
	DecodeTimestamp int64
	SourceData      []byte
	DecodedObject   interface{}
	SelfChecked     bool
	BTIH            BTIHData
}

type BTIHData struct {
	Bytes [20]byte
	Text  string
}

// Cached Errors.
var ErrAnomaly = errors.New(ErrAnomalyError)
var ErrByteStringToInt = errors.New(ErrByteStringToIntError)
var ErrDataType = errors.New(ErrDataTypeError)
var ErrHeaderLength = errors.New(ErrHeaderLengthError)
var ErrIntegerLength = errors.New(ErrIntegerLengthError)
var ErrSectionDoesNotExist = errors.New(ErrSectionDoesNotExistError)
var ErrSelfCheck = errors.New(ErrSelfCheckError)
var ErrTypeAssertion = errors.New(ErrTypeAssertionError)
