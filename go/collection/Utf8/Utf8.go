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

// Utf8.go.

package utf8

// Functions for Analysis of Text encoded in 'UTF-8' Format.

// Checks whether a String consists only of ASCII Letters and Numbers.
func StringHasOnlyASCIILetterNumberOrUnderline(
	s string,
) bool {

	var letters []rune

	letters = []rune(s)
	if len(letters) == 0 {
		return false
	}

	for _, letter := range letters {
		if SymbolIsASCIILetterNumberOrUnderline(letter) == false {
			return false
		}
	}

	return true
}

// Checks whether a String consists only of ASCII Letters and Numbers.
func StringHasOnlyASCIILetterOrNumber(
	s string,
) bool {

	var letters []rune

	letters = []rune(s)
	if len(letters) == 0 {
		return false
	}

	for _, letter := range letters {
		if SymbolIsASCIILetterOrNumber(letter) == false {
			return false
		}
	}

	return true
}

// Checks whether a Rune is an ASCII Letter, Number or Underline.
func SymbolIsASCIILetterNumberOrUnderline(
	r rune,
) bool {

	if ('A' <= r) && (r <= 'Z') {
		return true
	}

	if r == '_' {
		return true
	}

	if ('a' <= r) && (r <= 'z') {
		return true
	}

	if ('0' <= r) && (r <= '9') {
		return true
	}

	return false
}

// Checks whether a Rune is an ASCII Letter or Number.
func SymbolIsASCIILetterOrNumber(
	r rune,
) bool {

	if ('A' <= r) && (r <= 'Z') {
		return true
	}

	if ('a' <= r) && (r <= 'z') {
		return true
	}

	if ('0' <= r) && (r <= '9') {
		return true
	}

	return false
}
