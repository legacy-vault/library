package urlmaker

import (
	"testing"
)

func Test_URLArgument_IsValid(t *testing.T) {

	var argument URLArgument

	// Test #1. Negative Test.
	argument = URLArgument{Name: "", Value: 123}
	if argument.IsValid() != false {
		t.FailNow()
	}

	// Test #2. Positive Test.
	argument = URLArgument{Name: "x", Value: 123}
	if argument.IsValid() != true {
		t.FailNow()
	}
}
