package urlmaker

import (
	"fmt"
	"net/url"
)

type URL struct {
	Protocol string
	Path     string

	// Arguments Storage.
	// Stores Arguments Names, Values and their Order.
	arguments []URLArgument

	// Arguments Cache.
	// Is used to make 'Add' & 'Delete' Actions faster.
	argumentsCache map[string]interface{}
}

// AddArgument Method tries to add a new Argument to the End of the URL
// Argument's List. Also validates the Argument Name.
func (u *URL) AddArgument(
	arg URLArgument,
) error {
	if !arg.IsValid() {
		return fmt.Errorf(ErrfArgumentNotValid, arg.Name)
	}
	if u.argumentIsSet(arg.Name) {
		return fmt.Errorf(ErrfDuplicateArgumentName, arg.Name)
	}
	u.addArgument(arg)
	return nil
}

func (u *URL) addArgument(
	arg URLArgument,
) {
	u.arguments = append(u.arguments, arg)
	u.argumentsCache[arg.Name] = arg.Value
}

// ArgumentIsSet Method checks whether the specified Argument already exists in
// the URL. Also validates the Argument Name.
func (u *URL) ArgumentIsSet(
	argName string,
) (bool, error) {

	tmpArg := URLArgument{Name: argName, Value: nil}
	if !tmpArg.IsValid() {
		return false, fmt.Errorf(ErrfArgumentNotValid, tmpArg.Name)
	}

	return u.argumentIsSet(argName), nil
}

// argumentIsSet Method checks whether the specified Argument already exists in
// the URL.
func (u *URL) argumentIsSet(
	argName string,
) bool {
	_, argNameIsSet := u.argumentsCache[argName]
	return argNameIsSet
}

// String Method returns a full URL as a String.
func (u *URL) String() string {

	var argumentValueStr string
	var result string

	result = u.Protocol + ProtocolPostfix + u.Path
	if len(u.arguments) > 0 {
		result = result + ArgumentsPrefix
	}

	iLast := len(u.arguments) - 1
	for i, argument := range u.arguments {
		result = result + argument.Name
		if argument.Value != nil {
			argumentValueStr = fmt.Sprintf("%v", argument.Value)
			argumentValueStr = url.QueryEscape(argumentValueStr)
			result = result + ArgumentKeyValueSeparator + argumentValueStr
		}
		if i != iLast {
			result = result + ArgumentSeparator
		}
	}

	return result
}
