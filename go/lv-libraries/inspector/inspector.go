package inspector

import (
	"errors"
	"fmt"
	"reflect"
)

// VerifyStruct Function checks that all Fields of a short-Model Object exist
// in the long-Model Object and their Values match each other.
func VerifyStruct(
	objectShort interface{},
	objectLong interface{},
) (bool, error) {

	var err error
	var objectLongValue reflect.Value
	var objectLongFields map[string]interface{}
	var objectShortValue reflect.Value
	var objectShortFields map[string]interface{}

	objectShortValue = reflect.Indirect(reflect.ValueOf(objectShort))
	objectLongValue = reflect.Indirect(reflect.ValueOf(objectLong))

	if objectShortValue.Type().Kind() != reflect.Struct {
		return false, errors.New("Short Object is not a Struct")
	}
	if objectLongValue.Type().Kind() != reflect.Struct {
		return false, errors.New("Long Object is not a Struct")
	}

	objectShortFields = listStructFields(objectShortValue)
	objectLongFields = listStructFields(objectLongValue)

	//fmt.Println("objectShortFields:", objectShortFields) // Debug.
	//fmt.Println("objectLongFields:", objectLongFields)   // Debug.

	// Compare the Fields.
	for shortObjectFieldName, shortObjectFieldValue := range objectShortFields {
		longObjectFieldValue, longObjectFieldExists := objectLongFields[shortObjectFieldName]
		if !longObjectFieldExists {
			err = fmt.Errorf(
				"Field '%v' does not exist in long-Model Object",
				shortObjectFieldName,
			)
			return false, err
		}

		if longObjectFieldValue != shortObjectFieldValue {
			err = fmt.Errorf(
				"Field '%v' is not equal",
				shortObjectFieldName,
			)
			return false, err
		}
	}

	return true, nil
}

// listStructFields Function lists the Struct's Fields as a Map.
// Map's Key is a Field Name, Map's Value is the Field Value.
func listStructFields(structValue reflect.Value) map[string]interface{} {

	var structFields map[string]interface{}

	structFields = make(map[string]interface{})
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structValue.Type().Field(i).Name
		structFields[fieldName] = structValue.Field(i).Interface()
	}

	return structFields
}
