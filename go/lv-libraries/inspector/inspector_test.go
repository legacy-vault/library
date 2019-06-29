// +build test

package inspector

import (
	//"fmt" // Debug.
	"reflect"
	"testing"
)

func Test_VerifyStruct(t *testing.T) {

	type ClassX struct {
		Age  int
		Name string
	}
	type ClassY struct {
		Age   int
		Color int
		Name  string
	}

	var err error
	var objX1 = ClassX{
		Age:  10,
		Name: "John",
	}
	var objX2 = ClassX{
		Age:  10,
		Name: "John",
	}
	var objY1 = ClassY{
		Age:  10,
		Name: "John",
		// Additional Field.
		Color: 5,
	}
	var result bool

	// Test #1.1. Not a Struct.
	result, err = VerifyStruct("oops", objY1)
	if (result != false) || (err == nil) {
		t.FailNow()
	}

	// Test #1.2. Not a Struct.
	result, err = VerifyStruct(objY1, "oops")
	if (result != false) || (err == nil) {
		t.FailNow()
	}

	// Test #2.1. Totally equal Structs.
	result, err = VerifyStruct(objX1, objX2)
	if (result != true) || (err != nil) {
		t.FailNow()
	}

	// Test #2.2. Short-equal Structs.
	result, err = VerifyStruct(objX1, objY1)
	if (result != true) || (err != nil) {
		t.FailNow()
	}

	// Test #3. Not equal Structs.
	result, err = VerifyStruct(objY1, objX1)
	//fmt.Println("err:", err) // Debug.
	if (result != false) || (err == nil) {
		t.FailNow()
	}
}

func Test_listStructFields(t *testing.T) {

	type ClassX struct {
		Age  int
		Name string
	}

	var objX1 = ClassX{
		Age:  10,
		Name: "John",
	}

	// Test #1.1. Not a Struct.
	fieldList := listStructFields(reflect.ValueOf(objX1))
	//fmt.Println("fieldList:", fieldList) // Debug.
	if (len(fieldList) != 2) ||
		(fieldList["Age"] != 10) ||
		(fieldList["Name"] != "John") {
		t.FailNow()
	}
}
