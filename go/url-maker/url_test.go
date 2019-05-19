package urlmaker

import (
	"fmt"
	"testing"
)

func Test_URL_AddArgument(t *testing.T) {

	var argument URLArgument
	var err error
	var url *URL

	url = &URL{
		arguments: []URLArgument{
			{Name: "x", Value: 123},
		},
		argumentsCache: map[string]interface{}{
			"x": 123,
		},
	}

	// Test #1. Negative Test: Argument is not valid.
	argument = URLArgument{Name: "", Value: "John"}
	err = url.AddArgument(argument)
	if (err == nil) ||
		(len(url.arguments) != 1) ||
		(url.arguments[0].Name != "x") || (url.arguments[0].Value != 123) ||
		(len(url.argumentsCache) != 1) ||
		(url.argumentsCache["x"] != 123) {
		t.FailNow()
	}

	// Test #2. Negative Test: Argument is already set.
	argument = URLArgument{Name: "x", Value: 456}
	err = url.AddArgument(argument)
	if (err == nil) ||
		(len(url.arguments) != 1) ||
		(url.arguments[0].Name != "x") || (url.arguments[0].Value != 123) ||
		(len(url.argumentsCache) != 1) ||
		(url.argumentsCache["x"] != 123) {
		t.FailNow()
	}

	// Test #3. Positive Test.
	argument = URLArgument{Name: "y", Value: "John"}
	err = url.AddArgument(argument)
	if (err != nil) ||
		(len(url.arguments) != 2) ||
		(url.arguments[0].Name != "x") || (url.arguments[0].Value != 123) ||
		(url.arguments[1].Name != "y") || (url.arguments[1].Value != "John") ||
		(len(url.argumentsCache) != 2) ||
		(url.argumentsCache["x"] != 123) ||
		(url.argumentsCache["y"] != "John") {
		t.FailNow()
	}
}

func Test_URL_addArgument(t *testing.T) {

	var argument URLArgument
	var url *URL

	url = &URL{
		arguments: []URLArgument{
			{Name: "x", Value: 123},
		},
		argumentsCache: map[string]interface{}{
			"x": 123,
		},
	}
	argument = URLArgument{Name: "y", Value: "John"}

	// Test #1.
	url.addArgument(argument)
	if (len(url.arguments) != 2) ||
		(url.arguments[0].Name != "x") || (url.arguments[0].Value != 123) ||
		(url.arguments[1].Name != "y") || (url.arguments[1].Value != "John") ||
		(len(url.argumentsCache) != 2) ||
		(url.argumentsCache["x"] != 123) ||
		(url.argumentsCache["y"] != "John") {
		t.FailNow()
	}
}

func Test_URL_ArgumentIsSet(t *testing.T) {

	var argumentIsSet bool
	var err error
	var url *URL

	url = &URL{
		argumentsCache: map[string]interface{}{
			"x": 123,
		},
	}

	// Test #1. Negative Test: Argument is not valid.
	argumentIsSet, err = url.ArgumentIsSet("")
	if err == nil {
		t.FailNow()
	}

	// Test #2. Positive Test.
	argumentIsSet, err = url.ArgumentIsSet("x")
	if (err != nil) || (argumentIsSet != true) {
		t.FailNow()
	}
}

func Test_URL_argumentIsSet(t *testing.T) {

	var url *URL

	url = &URL{
		argumentsCache: map[string]interface{}{
			"x": 123,
		},
	}

	// Test #1. Positive Test.
	if url.argumentIsSet("x") != true {
		t.FailNow()
	}

	// Test #2. Negative Test.
	if url.argumentIsSet("y") != false {
		t.FailNow()
	}
}

func Test_URL_String(t *testing.T) {

	var err error
	var str string
	var url *URL

	// Test #1. No Arguments.
	url, err = NewURL(
		"http",
		"localhost",
		[]URLArgument{},
	)
	if err != nil {
		t.FailNow()
	}
	str = url.String()
	fmt.Println(str) // Debug.
	if str != "http://localhost" {
		t.FailNow()
	}

	// Test #2. Arguments are set.
	url, err = NewURL(
		"http",
		"localhost",
		[]URLArgument{
			{Name: "x", Value: 123},
			{Name: "y", Value: nil},
			{Name: "z", Value: "Q&Z АБВ"},
		},
	)
	if err != nil {
		t.FailNow()
	}
	str = url.String()
	fmt.Println(str) // Debug.
	if str != "http://localhost?x=123&y&z=Q%26Z+%D0%90%D0%91%D0%92" {
		t.FailNow()
	}
}
