package boolean

import (
	"fmt"
	"strings"
)

const (
	TrueStringLc  = "true"
	FalseStringLc = "false"
)

const ErrfBadBoolean = "Bad boolean value: '%v'."

func StringToFloat(str string) (bool, error) {

	var err error

	if len(str) < 4 {
		err = fmt.Errorf(ErrfBadBoolean, str)
		return false, err
	}

	str = strings.ToLower(str)

	switch str {

	case TrueStringLc:
		return true, nil

	case FalseStringLc:
		return false, nil

	default:
		err = fmt.Errorf(ErrfBadBoolean, str)
		return false, err
	}
}
