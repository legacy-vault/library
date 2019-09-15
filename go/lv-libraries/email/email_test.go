// +build test

package email

import (
	"testing"
)

func Test_EmailAddressIsValid(t *testing.T) {

	var err error
	var ok bool
	var testDataPositive []string
	var testDataNegative []string

	// Part 1. Positive Tests.
	testDataPositive = make([]string, 0)
	testDataPositive = append(testDataPositive, "username@example.org")
	testDataPositive = append(testDataPositive, "username-a.x_y.m15@example.org")
	testDataPositive = append(testDataPositive, "username@domain33")
	testDataPositive = append(testDataPositive, "x22@host55")

	for _, td := range testDataPositive {
		ok, err = EmailAddressIsValid(td)
		if (ok != true) || (err != nil) {
			t.Error("Positive Test")
			t.FailNow()
		}
	}

	// Part 2. Negative Tests.
	testDataNegative = make([]string, 0)
	testDataNegative = append(testDataNegative, "")            // No '@'.
	testDataNegative = append(testDataNegative, "aaa@bbb@ccc") // Double '@'.
	testDataNegative = append(testDataNegative, "@aa")         // Empty Username.
	testDataNegative = append(testDataNegative, "aa@")         // Empty Hostname.
	testDataNegative = append(testDataNegative, ".@xxx")       // Function Entry Test.
	testDataNegative = append(testDataNegative, "xxx@.")

	for _, td := range testDataNegative {
		ok, err = EmailAddressIsValid(td)
		if (ok != false) || (err == nil) {
			t.Error("Negative Test")
			t.FailNow()
		}
	}
}

func Test_emailAddressPartSymbolIsAllowed(t *testing.T) {

	type Dataset struct {
		Symbol       rune
		Position     int
		StringLength int
	}

	var err error
	var ok bool
	var testDataPositive []Dataset
	var testDataNegative []Dataset

	// Part 1. Positive Tests.
	testDataPositive = make([]Dataset, 0)
	testDataPositive = append(testDataPositive, Dataset{Symbol: 'a', Position: 0, StringLength: 10})
	testDataPositive = append(testDataPositive, Dataset{Symbol: 'b', Position: 4, StringLength: 10})
	testDataPositive = append(testDataPositive, Dataset{Symbol: 'c', Position: 9, StringLength: 10})

	testDataPositive = append(testDataPositive, Dataset{Symbol: '8', Position: 4, StringLength: 10})
	testDataPositive = append(testDataPositive, Dataset{Symbol: '8', Position: 9, StringLength: 10})

	testDataPositive = append(testDataPositive, Dataset{Symbol: '.', Position: 4, StringLength: 10})
	testDataPositive = append(testDataPositive, Dataset{Symbol: '-', Position: 4, StringLength: 10})
	testDataPositive = append(testDataPositive, Dataset{Symbol: '_', Position: 4, StringLength: 10})

	for _, td := range testDataPositive {
		ok, err = emailAddressPartSymbolIsAllowed(td.Symbol, td.Position, td.StringLength)
		if (ok != true) || (err != nil) {
			t.Error("Positive Test")
			t.FailNow()
		}
	}

	// Part 2. Negative Tests.
	testDataNegative = make([]Dataset, 0)
	testDataNegative = append(testDataNegative, Dataset{Symbol: '8', Position: 0, StringLength: 10})

	testDataNegative = append(testDataNegative, Dataset{Symbol: '.', Position: 0, StringLength: 10})
	testDataNegative = append(testDataNegative, Dataset{Symbol: '.', Position: 9, StringLength: 10})
	testDataNegative = append(testDataNegative, Dataset{Symbol: '-', Position: 0, StringLength: 10})
	testDataNegative = append(testDataNegative, Dataset{Symbol: '-', Position: 9, StringLength: 10})
	testDataNegative = append(testDataNegative, Dataset{Symbol: '_', Position: 0, StringLength: 10})
	testDataNegative = append(testDataNegative, Dataset{Symbol: '_', Position: 9, StringLength: 10})

	for _, td := range testDataNegative {
		ok, err = emailAddressPartSymbolIsAllowed(td.Symbol, td.Position, td.StringLength)
		if (ok != false) || (err == nil) {
			t.Error("Negative Test")
			t.FailNow()
		}
	}
}
