package urlmaker

import (
	"fmt"
)

// Package 'URL Maker' provides a convenient Way to create simple URLs.
// The main Focus in this Package is done on the URL Arguments.
// URL Arguments can be easily manipulated. This can not be with the Golang's
// built-in 'url' Package.

// NewURL Function creates a new URL.
// Arguments must have unique Names.
// The Order of Arguments is important.
func NewURL(
	protocol string,
	path string,
	arguments []URLArgument,
) (*URL, error) {

	var argsMap map[string]interface{}
	var err error
	var url = new(URL)

	url.Protocol = protocol

	url.Path = path

	argsMap, err = prepareArguments(arguments)
	if err != nil {
		return nil, err
	}
	url.arguments = arguments
	url.argumentsCache = argsMap
	return url, nil
}

// prepareArguments Function performs a Check combined with the Arguments
// Preprocess Stage. Returns an Error when Names of Arguments are neither
// unique nor valid.
func prepareArguments(
	arguments []URLArgument,
) (map[string]interface{}, error) {

	var argumentsCache = make(map[string]interface{})
	var duplicateName bool

	for _, arg := range arguments {

		// Name Validation.
		if !arg.IsValid() {
			return nil, fmt.Errorf(ErrfArgumentNotValid, arg.Name)
		}

		// Duplicates Check.
		_, duplicateName = argumentsCache[arg.Name]
		if duplicateName {
			return nil, fmt.Errorf(ErrfDuplicateArgumentName, arg.Name)
		}
		argumentsCache[arg.Name] = arg.Value
	}

	return argumentsCache, nil
}
