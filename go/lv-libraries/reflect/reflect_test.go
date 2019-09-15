// +build test

package reflect

import (
	"testing"

	"github.com/fatih/structtag"

	"github.com/legacy-vault/library/go/lv-libraries/tester"
)

func Test_List1stLevelFields(t *testing.T) {

	type SampleClassA struct {
		Age  int
		Name string
	}
	type SampleClassB struct {
		Age  int    `json:"x,y"`
		Name string `json:"a,b" sql:"c,d,e"`
	}

	var aTest *tester.Test
	var err error
	var fields []StructField
	var tag *structtag.Tag

	aTest = tester.NewTest(t)

	// Test #1. Not a Struct.
	fields, err = List1stLevelFields(123)
	aTest.MustBeAnError(err)
	if fields != nil {
		t.FailNow()
	}

	// Test #2. A Struct without Tags.
	fields, err = List1stLevelFields(
		SampleClassA{
			Age:  10,
			Name: "John",
		},
	)
	aTest.MustBeNoError(err)
	field1Tags := (fields[0].Tags)
	field2Tags := (fields[1].Tags)
	if (len(fields) != 2) ||
		(fields[0].Name != "Age") || (fields[0].TagRaw != "") || (field1Tags.Len() != 0) ||
		(fields[1].Name != "Name") || (fields[1].TagRaw != "") || (field2Tags.Len() != 0) {
		t.FailNow()
	}

	// Test #3. A Struct with Tags.
	fields, err = List1stLevelFields(
		SampleClassB{
			Age:  10,
			Name: "John",
		},
	)
	aTest.MustBeNoError(err)
	field1Tags = (fields[0].Tags)
	field2Tags = (fields[1].Tags)
	if len(fields) != 2 {
		t.FailNow()
	}

	// Check the first Field ('Age').
	if (fields[0].Name != "Age") || (fields[0].TagRaw == "") || (field1Tags.Len() != 1) {
		t.FailNow()
	}
	tag, err = field1Tags.Get("json")
	aTest.MustBeNoError(err)
	if (tag.Name != "x") || (len(tag.Options) != 1) || (tag.Options[0] != "y") {
		t.FailNow()
	}

	// Check the second Field ('Name').
	if (fields[1].Name != "Name") || (fields[1].TagRaw == "") || (field2Tags.Len() != 2) {
		t.FailNow()
	}
	tag, err = field2Tags.Get("json")
	aTest.MustBeNoError(err)
	if (tag.Name != "a") || (len(tag.Options) != 1) || (tag.Options[0] != "b") {
		t.FailNow()
	}
	tag, err = field2Tags.Get("sql")
	aTest.MustBeNoError(err)
	if (tag.Name != "c") || (len(tag.Options) != 2) || (tag.Options[0] != "d") || (tag.Options[1] != "e") {
		t.FailNow()
	}
}

func Test_List1stLevelFieldsWithTagType(t *testing.T) {

	type SampleClassC struct {
		ID         int
		Age        int    `json:"x,y"`
		Name       string `json:"a,b" sql:"c,d,e"`
		Size       int    `sql:"f,g"`
		HiddenName string `json:"-"`
	}

	var aTest *tester.Test
	var err error
	var fieldNames []string

	aTest = tester.NewTest(t)

	// Test #1.
	fieldNames, err = List1stLevelFieldsWithTagType(SampleClassC{}, "json")
	aTest.MustBeNoError(err)
	if (len(fieldNames) != 2) ||
		(fieldNames[0] != "Age") || (fieldNames[1] != "Name") {
		t.FailNow()
	}
}

func Test_List1stLevelFieldTagsWithTagType(t *testing.T) {

	type SampleClassC struct {
		ID   int
		Age  int    `json:"x,y" zz:"Age,101"`
		Name string `qq:""`
		Size int    `zz:"Size,102" sql:"f,g"`
	}

	var aTest *tester.Test
	var err error
	var fieldTagNames []string

	aTest = tester.NewTest(t)

	// Test #1.
	fieldTagNames, err = List1stLevelFieldTagsWithTagType(SampleClassC{}, "zz")
	aTest.MustBeNoError(err)
	if (len(fieldTagNames) != 2) ||
		(fieldTagNames[0] != "Age") || (fieldTagNames[1] != "Size") {
		t.FailNow()
	}
}
