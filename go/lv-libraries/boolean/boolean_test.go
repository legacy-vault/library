// +build test

package boolean

import (
	"testing"

	"github.com/legacy-vault/library/go/lv-libraries/tester"
)

func Test_StringToFloat(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result bool

	aTest = tester.NewTest(t)

	// Test #1.
	result, err = StringToFloat("TRuE")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, true)

	// Test #2.
	result, err = StringToFloat("FalsE")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, false)

	// Test #3.
	result, err = StringToFloat("x")
	aTest.MustBeAnError(err)

	// Test #4.
	result, err = StringToFloat("123456789")
	aTest.MustBeAnError(err)
}
