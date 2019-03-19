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

// Utf8_test.go.

package utf8

// Functions for Analysis of Text encoded in 'UTF-8' Format.

import "testing"

// Checks whether a Rune is ASCII Letter or Number.
func Test_All(t *testing.T) {

	var result bool

	result = SymbolIsASCIILetterOrNumber('!')
	if result != false {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('0')
	if result != true {
		t.FailNow()
	}
	result = SymbolIsASCIILetterOrNumber('1')
	if result != true {
		t.FailNow()
	}
	result = SymbolIsASCIILetterOrNumber('5')
	if result != true {
		t.FailNow()
	}
	result = SymbolIsASCIILetterOrNumber('9')
	if result != true {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber(':')
	if result != false {
		t.FailNow()
	}
	result = SymbolIsASCIILetterOrNumber('=')
	if result != false {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('X')
	if result != true {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('_')
	if result != false {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('x')
	if result != true {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('|')
	if result != false {
		t.FailNow()
	}
	result = SymbolIsASCIILetterOrNumber('Ы')
	if result != false {
		t.FailNow()
	}

	result = SymbolIsASCIILetterOrNumber('_')
	if result != false {
		t.FailNow()
	}

	if StringHasOnlyASCIILetterOrNumber("") != false {
		t.FailNow()
	}
	if StringHasOnlyASCIILetterOrNumber("abcxyzABCXYZ0123456789") != true {
		t.FailNow()
	}
	if StringHasOnlyASCIILetterOrNumber("abcxyzABCXYZ0123456789!+ЖЩЫ") != false {
		t.FailNow()
	}

	result = SymbolIsASCIILetterNumberOrUnderline('_')
	if result != true {
		t.FailNow()
	}
	result = SymbolIsASCIILetterNumberOrUnderline('*')
	if result != false {
		t.FailNow()
	}

	if StringHasOnlyASCIILetterNumberOrUnderline("abcxyzABCXYZ0123456789_") != true {
		t.FailNow()
	}
	if StringHasOnlyASCIILetterNumberOrUnderline("abcxyzABCXYZ0123456789_*") != false {
		t.FailNow()
	}
}
