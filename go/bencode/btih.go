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

// btih.go.

// BitTorrent Info Hash (BTIH) Check Sum Calculator.

package bencode

import (
	"crypto/sha1"
	"encoding/hex"
)

// Calculates the BitTorrent Info Hash (BTIH) Check Sum.
func (do *DecodedObject) CalculateBTIH() error {

	var err error
	var infoSection interface{}
	var infoSectionBA []byte

	// Get the 'info' Section from the decoded Object.
	infoSection, err = do.GetInfoSection()
	if err != nil {
		return err
	}

	// Encode the 'info' Section.
	infoSectionBA, err = Encode(infoSection)
	if err != nil {
		return err
	}

	// Calculate BTIH.
	do.BTIH.Bytes, do.BTIH.Text = CalculateSHA1(infoSectionBA)

	return nil
}

// Calculates SHA-1 Check Sum and
// returns it as a Hexadecimal Text and Byte Array.
func CalculateSHA1(ba []byte) ([20]byte, string) {

	var sha1sum [20]byte
	var sha1sumStrHex string

	sha1sum = sha1.Sum(ba)
	sha1sumStrHex = hex.EncodeToString(sha1sum[:])

	return sha1sum, sha1sumStrHex
}

// Gets an 'info' Section from the Object.
func (do DecodedObject) GetInfoSection() (interface{}, error) {

	var dictItem DictionaryItem
	var dictionary []DictionaryItem
	var ok bool

	// Get Dictionary.
	dictionary, ok = do.DecodedObject.([]DictionaryItem)
	if !ok {
		return nil, ErrTypeAssertion
	}

	// Get the 'info' Section from the decoded Object.
	for _, dictItem = range dictionary {

		if string(dictItem.Key) == FileSectionInfo {
			return dictItem.Value, nil
		}
	}

	return nil, ErrSectionDoesNotExist
}
