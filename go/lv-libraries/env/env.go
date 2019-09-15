package env

import (
	"fmt"
	"os"
)

const ErrfEnvEmpty = "Environment Variable '%v' is empty."

func GetEnv(variableName string) (string, error) {

	var envValue string
	var err error

	envValue = os.Getenv(variableName)
	if len(envValue) == 0 {
		err = fmt.Errorf(
			ErrfEnvEmpty,
			variableName,
		)
		return "", err
	}

	return envValue, nil
}
