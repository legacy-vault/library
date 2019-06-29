package url

import (
	"fmt"
	"net/http"
	"strconv"
)

// Functions to work with URL Parameters.

// Errors.
const (
	ErrfQueryParametersCount = "Query parameters syntax error. Count of the parameter '%v' is %v."
)

// ParseURLParameters Function parses URL Query Parameters into the Map.
// Requires every Parameter to be listed no more than once,
// e.g. the Request "hello.com?x=1&y=2&x=3" will raise an Error.
func ParseURLParameters(request *http.Request) (map[string]string, error) {

	var err error
	var result map[string]string

	err = request.ParseForm()
	if err != nil {
		return nil, err
	}

	result = make(map[string]string)

	for key, values := range request.Form {
		if len(values) != 1 {
			return nil, fmt.Errorf(
				ErrfQueryParametersCount,
				key,
				len(values),
			)
		}
		result[key] = values[0]
	}

	return result, nil
}

// ReadURLParameterOfUint64Type Function tries to get the Request Parameter of the
// 'uint64' Type from its URL.
func ReadURLParameterOfUint64Type(
	queryParameters map[string]string,
	parameterName string,
) (parameterExists bool, parameterValue uint64, err error) {

	var tmp string

	tmp, parameterExists = queryParameters[parameterName]
	if !parameterExists {
		return
	}

	parameterValue, err = strconv.ParseUint(tmp, 10, 64)
	return
}

// ReadURLParameterOfStringType Function tries to get the Request Parameter of the
// 'string' Type from its URL.
func ReadURLParameterOfStringType(
	queryParameters map[string]string,
	parameterName string,
) (parameterExists bool, parameterValue string, err error) {

	parameterValue, parameterExists = queryParameters[parameterName]
	return
}

// ReadHeaderParameterOfUint64Type Function tries to get the Request Parameter of the
// 'uint64' Type from its Header. The first returned Parameter is 'true' when
//// the Parameter is set in the Header.
func ReadHeaderParameterOfUint64Type(
	r *http.Request,
	parameterName string,
) (bool, uint64, error) {

	var err error
	var requestHeaderRawValue string
	var value uint64

	requestHeaderRawValue = r.Header.Get(parameterName)
	if len(requestHeaderRawValue) == 0 {
		return false, value, nil
	}

	value, err = strconv.ParseUint(requestHeaderRawValue, 10, 64)
	if err != nil {
		return true, value, err
	}

	return true, value, nil
}

// ReadHeaderParameterOfStringType Function tries to get the Request Parameter of the
// 'string' Type from its Header. The first returned Parameter is 'true' when
//// the Parameter is set in the Header.
func ReadHeaderParameterOfStringType(
	r *http.Request,
	parameterName string,
) (bool, string, error) {

	var requestHeaderRawValue string

	requestHeaderRawValue = r.Header.Get(parameterName)
	if len(requestHeaderRawValue) == 0 {
		return false, requestHeaderRawValue, nil
	}

	return true, requestHeaderRawValue, nil
}
