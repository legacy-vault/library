// +build test

package urlmaker

import (
	"fmt"
	"testing"
)

func Test_NewURL(t *testing.T) {

	var arguments []URLArgument
	var err error
	var url *URL

	// Test #1. Negative Test: 'prepareArguments' return an Error.
	arguments = []URLArgument{
		{Name: "x", Value: 123},
		{Name: "x", Value: "John"},
	}
	url, err = NewURL("protocol", "path", arguments)
	if (err == nil) || (url != nil) {
		t.FailNow()
	}

	// Test #2. Positive Test: 'prepareArguments' return no Error.
	arguments = []URLArgument{
		{Name: "x", Value: 123},
		{Name: "y", Value: "John"},
	}
	url, err = NewURL("protocol", "path", arguments)
	if (err != nil) ||
		(url.Protocol != "protocol") || (url.Path != "path") ||
		(len(url.arguments) != 2) ||
		(url.arguments[0].Name != "x") || (url.arguments[0].Value != 123) ||
		(url.arguments[1].Name != "y") || (url.arguments[1].Value != "John") ||
		(len(url.argumentsCache) != 2) ||
		(url.argumentsCache["x"] != 123) ||
		(url.argumentsCache["y"] != "John") {
		t.FailNow()
	}
}

func Test_prepareArguments(t *testing.T) {

	var argsMap map[string]interface{}
	var arguments []URLArgument
	var err error

	// Test #1. Negative Test: Argument is not valid.
	arguments = []URLArgument{
		{Name: "x", Value: 123},
		{Name: "", Value: "John"},
	}
	argsMap, err = prepareArguments(arguments)
	fmt.Println(err) // Debug.
	if (err == nil) || (argsMap != nil) {
		t.FailNow()
	}

	// Test #2. Negative Test: Duplicate Arguments.
	arguments = []URLArgument{
		{Name: "x", Value: 123},
		{Name: "x", Value: "John"},
	}
	argsMap, err = prepareArguments(arguments)
	fmt.Println(err) // Debug.
	if (err == nil) || (argsMap != nil) {
		t.FailNow()
	}

	// Test #3. Positive Test.
	arguments = []URLArgument{
		{Name: "x", Value: 123},
		{Name: "y", Value: "John"},
	}
	argsMap, err = prepareArguments(arguments)
	if (err != nil) || (len(argsMap) != 2) ||
		(argsMap["x"] != 123) || (argsMap["y"] != "John") {
		t.FailNow()
	}
}
