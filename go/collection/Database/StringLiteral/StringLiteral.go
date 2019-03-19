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

// StringLiteral.go.

package stringliteral

// String Literal.

const SymbolSingleQuote = '\''
const SymbolDoubleQuote = '"'
const SymbolBackSlash = '\\'
const SymbolPercent = '%'

type StringLiteral struct {
	Value string
}

// Returns a Value with special Symbols escaped.
func (this StringLiteral) Escaped() string {

	return escape(this.Value)
}

// Returns a Value with special Symbols escaped & put into single Quotes.
func (this StringLiteral) EscapedAndSingleQuoted() string {

	var result string

	result = string(SymbolSingleQuote) +
		escape(this.Value) +
		string(SymbolSingleQuote)

	return result
}

// Escapes special Symbols.
func escape(
	str string,
) string {

	var result string

	symbols := []rune(str)
	for _, symbol := range symbols {

		switch symbol {

		case SymbolDoubleQuote:
			result = result +
				string(SymbolBackSlash) + string(SymbolDoubleQuote)

		case SymbolSingleQuote:
			result = result +
				string(SymbolBackSlash) + string(SymbolSingleQuote)

		case SymbolBackSlash:
			result = result +
				string(SymbolBackSlash) + string(SymbolBackSlash)

		case SymbolPercent:
			result = result +
				string(SymbolBackSlash) + string(SymbolPercent)

		default:
			result = result + string(symbol)

		}
	}

	return result
}
