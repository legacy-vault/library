// +build test

package time

import (
	"testing"
	"time"

	"github.com/legacy-vault/library/go/lv-libraries/tester"
)

func Test_NewTimeString(t *testing.T) {

	var aTest *tester.Test
	var result string
	var resultExpected string

	aTest = tester.NewTest(t)
	resultExpected = "2019-06-11T16:44:05Z"

	// Test #1. Short Year.
	result = NewTimeStringRFC3339(
		19,
		6,
		11,
		16,
		44,
		5,
	)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2. Long Year.
	result = NewTimeStringRFC3339(
		2019,
		6,
		11,
		16,
		44,
		5,
	)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Minimum(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	aTest = tester.NewTest(t)

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time1)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Maximum(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	aTest = tester.NewTest(t)

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time2)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_IsEmpty(t *testing.T) {

	var aTest *tester.Test
	var err error
	var x time.Time

	aTest = tester.NewTest(t)

	// Test #1.
	x, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), false)

	// Test #2.
	x = time.Time{}
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), true)
}
