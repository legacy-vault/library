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

// Identifier.go.

package identifier

// Database Identifier.

import (
	"unicode"

	"xxx/Utf8"
)

const AllowedSymbolUnderline = '_'
const SystemPrefix = AllowedSymbolUnderline
const Separator = AllowedSymbolUnderline

type Identifier struct {
	Name string
}

// Checks whether a Database Identifier has a good Name.
// The 'systemId' Parameter is used to check System IDs.
func (this Identifier) IsGood(
	systemId bool,
) bool {

	var firstLetter rune
	var lastLetter rune
	var letters []rune

	// Check Length.
	letters = []rune(this.Name)
	if len(letters) == 0 {
		return false
	}

	// Check first Symbol.
	firstLetter = letters[0]
	// Underline first Symbol is allowed for System IDs.
	if !systemId {
		if firstLetter == SystemPrefix {
			return false
		}
	}
	if unicode.IsNumber(firstLetter) {
		return false
	}

	// Check last Symbol.
	lastLetter = letters[len(letters)-1]
	if lastLetter == Separator {
		return false
	}

	// Check all Characters.
	if utf8.StringHasOnlyASCIILetterNumberOrUnderline(this.Name) {
		return true
	}

	return false
}
