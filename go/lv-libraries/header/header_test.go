//+build test

package header

import (
	"testing"

	"github.com/legacy-vault/library/go/lv-libraries/tester"
)

func Test_MakeListOfHeaders(t *testing.T) {

	var aTest *tester.Test
	var headers = []string{"aa", "bb", "cc"}
	var result string

	aTest = tester.NewTest(t)
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "aa, bb, cc")
}
