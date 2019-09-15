package email

import (
	"errors"
	"fmt"
	"strings"

	"github.com/legacy-vault/library/go/lv-libraries/unicode"
)

// Errors
const (
	ErrEmailShort    = "E-Mail address is too short"
	ErrEmailAtSymbol = "'@' symbol is either absent or occurs multiple times"

	ErrfEmailStart  = "Forbidden symbol '%v' at the start position"
	ErrfEmailEnd    = "Forbidden symbol '%v' at the end position"
	ErrfEmailSymbol = "Forbidden symbol '%v' at position %v"
)

// E-Mail Constants.
const (
	EmailUsernameHostnameDelimiter = "@"
)

// EmailAddressIsValid Function checks whether the specified E-mail Address is
// valid or not.
func EmailAddressIsValid(emailAddress string) (bool, error) {

	var err error
	var ok bool
	var symbols []rune

	// Get Username and Hostname.
	emailParts := strings.Split(emailAddress, EmailUsernameHostnameDelimiter)
	if len(emailParts) != 2 {
		return false, errors.New(ErrEmailAtSymbol)
	}

	// Check the Username and Hostname Parts.
	for i := 0; i <= 1; i++ {
		emailPartLen := len(emailParts[i])
		if emailPartLen < 1 {
			return false, errors.New(ErrEmailShort)
		}
		symbols = []rune(emailParts[i])
		for pos, symbol := range symbols {
			ok, err = emailAddressPartSymbolIsAllowed(symbol, pos, emailPartLen)
			if !ok {
				return ok, err
			}
		}
	}

	return true, nil
}

// emailAddressPartSymbolIsAllowed Function checks whether the Symbol is allowed
// in a E-Mail Address Part. This Checker is compliant with the original
// Specification of Hostnames in the RFC 952, which mandated that Labels could
// not start with a Digit or with a Hyphen, and must not end with a Hyphen.
//
// Rules:
//
// 1. Allowed Symbols at the Start Position of the Username Part and Hostname Part
// of the Address: Letters only.
//
// 2. Allowed Symbols at the End Position of the Username Part and Hostname Part
// of the Address: Letters and Numbers only.
func emailAddressPartSymbolIsAllowed(
	symbol rune,
	position int, // Zero-based Index from 0 to "Length minus 1".
	stringLength int,
) (bool, error) {

	// Latin Letters may be at any Place.
	if unicode.SymbolIsLatLetter(symbol) {
		return true, nil
	}

	// Numbers may be at any Place except the Start Position.
	if unicode.SymbolIsNumber(symbol) {
		if position == 0 {
			return false, fmt.Errorf(ErrfEmailEnd, string(symbol))
		}
		return true, nil
	}

	// Process a non-letter and non-number Symbol.
	// It must not be at the Corner Position.
	if position == 0 {
		return false, fmt.Errorf(ErrfEmailStart, string(symbol))
	}
	if position == stringLength-1 {
		return false, fmt.Errorf(ErrfEmailEnd, string(symbol))
	}
	if (symbol == '.') || (symbol == '-') || (symbol == '_') {
		return true, nil
	}
	return false, fmt.Errorf(ErrfEmailSymbol, string(symbol), position)
}
