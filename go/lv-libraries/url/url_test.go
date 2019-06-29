// +build test

package url

import (
	"net/http"
	"strings"
	"testing"

	hh "github.com/legacy-vault/library/go/lv-libraries/http-helper"
	"github.com/legacy-vault/library/go/lv-libraries/tester"
)

func Test_ParseUrlParameters(t *testing.T) {

	var aTest *tester.Test
	var err error
	var request *http.Request
	var result map[string]string

	aTest = tester.NewTest(t)

	// Test #1. Normal Data.
	request, err = http.NewRequest(
		"GET",
		"hello.com?x=1&y=2",
		strings.NewReader(""),
	)
	aTest.MustBeNoError(err)
	result, err = ParseURLParameters(request)
	aTest.MustBeNoError(err)
	if (result["x"] != "1") || (result["y"] != "2") {
		t.FailNow()
	}

	// Test #2. Bad Data: double Parameter 'x'.
	request, err = http.NewRequest(
		"GET",
		"hello.com?x=1&y=2&x=3",
		strings.NewReader(""),
	)
	aTest.MustBeNoError(err)
	result, err = ParseURLParameters(request)
	aTest.MustBeAnError(err)
}

func Test_ReadURLParameterOfUint64Type(t *testing.T) {

	var aSimpleHttpTest hh.SimpleTest
	var aTest *tester.Test
	var err error

	aTest = tester.NewTest(t)

	// Test #1. Positive.
	aSimpleHttpTest = hh.SimpleTest{
		Parameter: hh.SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org?x=123",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: hh.SimpleTestResult{
			ResponseBodyString: "",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {

		var herr error
		var rawParams map[string]string

		rawParams, herr = ParseURLParameters(r)
		aTest.MustBeNoError(herr)
		parameterExists, parameterValue, herr :=
			ReadURLParameterOfUint64Type(rawParams, "x")
		aTest.MustBeNoError(herr)
		aTest.MustBeEqual(parameterExists, true)
		aTest.MustBeEqual(parameterValue, uint64(123))

		w.WriteHeader(http.StatusAccepted)
	}
	err = hh.PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)
}

func Test_ReadURLParameterOfStringType(t *testing.T) {

	var aSimpleHttpTest hh.SimpleTest
	var aTest *tester.Test
	var err error

	aTest = tester.NewTest(t)

	// Test #1. Positive.
	aSimpleHttpTest = hh.SimpleTest{
		Parameter: hh.SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org?xyz=John",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: hh.SimpleTestResult{
			ResponseBodyString: "",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {

		var herr error
		var rawParams map[string]string

		rawParams, herr = ParseURLParameters(r)
		aTest.MustBeNoError(herr)
		parameterExists, parameterValue, herr :=
			ReadURLParameterOfStringType(rawParams, "xyz")
		aTest.MustBeNoError(herr)
		aTest.MustBeEqual(parameterExists, true)
		aTest.MustBeEqual(parameterValue, "John")

		w.WriteHeader(http.StatusAccepted)
	}
	err = hh.PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)
}

func Test_ReadHeaderParameterOfUint64Type(t *testing.T) {

	const TestedHeaderName = "X-Counter"
	const TestedHeaderValue = "12345"

	var anAverageHttpTest hh.AverageTest
	var aTest *tester.Test
	var err error

	aTest = tester.NewTest(t)
	headersCommon := http.Header{}
	headersCommon.Add(TestedHeaderName, TestedHeaderValue)

	// Test #1. Positive.
	anAverageHttpTest = hh.AverageTest{
		Parameter: hh.AverageTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestHeaders: headersCommon,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: hh.AverageTestResult{
			ResponseStatusCode: http.StatusAccepted,
			ResponseHeaders:    http.Header{},
			ResponseBody:       []byte{},
		},
	}
	anAverageHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {

		var herr error

		// Get the incoming Header.
		paramExists, paramValue, herr :=
			ReadHeaderParameterOfUint64Type(r, TestedHeaderName)
		aTest.MustBeNoError(herr)

		// Verify the incoming Header.
		aTest.MustBeEqual(paramExists, true)
		aTest.MustBeEqual(paramValue, uint64(12345))

		w.WriteHeader(http.StatusAccepted)
	}
	err = hh.PerformAverageHttpTest(&anAverageHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(anAverageHttpTest.ResultReceived, anAverageHttpTest.ResultExpected)
}

func Test_ReadHeaderParameterOfStringType(t *testing.T) {

	const TestedHeaderName = "X-Name"
	const TestedHeaderValue = "Alice"

	var anAverageHttpTest hh.AverageTest
	var aTest *tester.Test
	var err error

	aTest = tester.NewTest(t)
	headersCommon := http.Header{}
	headersCommon.Add(TestedHeaderName, TestedHeaderValue)

	// Test #1. Positive.
	anAverageHttpTest = hh.AverageTest{
		Parameter: hh.AverageTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestHeaders: headersCommon,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: hh.AverageTestResult{
			ResponseStatusCode: http.StatusAccepted,
			ResponseHeaders:    http.Header{},
			ResponseBody:       []byte{},
		},
	}
	anAverageHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {

		var herr error

		// Get the incoming Header.
		paramExists, paramValue, herr :=
			ReadHeaderParameterOfStringType(r, TestedHeaderName)
		aTest.MustBeNoError(herr)

		// Verify the incoming Header.
		aTest.MustBeEqual(paramExists, true)
		aTest.MustBeEqual(paramValue, TestedHeaderValue)

		w.WriteHeader(http.StatusAccepted)
	}
	err = hh.PerformAverageHttpTest(&anAverageHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(anAverageHttpTest.ResultReceived, anAverageHttpTest.ResultExpected)
}
