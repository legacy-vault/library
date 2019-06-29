package reflect

import (
	"errors"
	"reflect"

	"github.com/fatih/structtag"
)

// Errors.
const (
	ErrNotStruct = "Not a Struct"
)

// Struct Tag Types.
const (
	StructTagTypeJSON = "json"
)

// JSON Tag specific Values.
const StructTagTypeJSONValueHidden = "-"

type StructField struct {
	Name   string
	TagRaw reflect.StructTag
	Tags   *structtag.Tags
}

// List1stLevelFields Function returns a List of Filed Names and Field Tags of
// a Struct.
func List1stLevelFields(obj interface{}) ([]StructField, error) {

	var err error
	var fieldsCount int
	var field reflect.StructField
	var fields []StructField
	var i int
	var objValue reflect.Value
	var tags *structtag.Tags

	// Input Data Verification.
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return nil, errors.New(ErrNotStruct)
	}

	// Prepare Data.
	objValue = reflect.ValueOf(obj)
	fieldsCount = objValue.NumField()
	fields = make([]StructField, 0)

	// Inspect all Fields of an Object.
	for i = 0; i < fieldsCount; i++ {
		field = objValue.Type().Field(i)

		tags, err = structtag.Parse(string(field.Tag))
		if err != nil {
			return nil, err
		}

		fields = append(
			fields,
			StructField{
				Name:   field.Name,
				Tags:   tags,
				TagRaw: field.Tag,
			},
		)
	}

	return fields, nil
}

// List1stLevelFieldsWithJsonTag Function returns a List of Struct's Fields
// which have a Tag with JSON Data.
func List1stLevelFieldsWithJsonTag(obj interface{}) ([]string, error) {
	return List1stLevelFieldsWithTagType(obj, StructTagTypeJSON)
}

// List1stLevelFieldsWithTagType Function returns a List of Struct's Field
// Names which have a Tag with the specified Type.
func List1stLevelFieldsWithTagType(obj interface{}, tagType string) ([]string, error) {

	var err error
	var fields []StructField
	var fieldNames []string
	var tag *structtag.Tag

	fields, err = List1stLevelFields(obj)
	if err != nil {
		return nil, err
	}

	fieldNames = make([]string, 0)
	for _, field := range fields {
		if field.Tags == nil {
			continue
		}
		tag, err = field.Tags.Get(tagType)
		if err != nil {
			continue
		}
		if tag.Name == StructTagTypeJSONValueHidden {
			continue
		}
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames, nil
}

// List1stLevelFieldTagsWithTagType Function returns a List of Struct's Field
// Tag Values which have the specified Type.
func List1stLevelFieldTagsWithTagType(obj interface{}, tagType string) ([]string, error) {

	var err error
	var fields []StructField
	var fieldTagNames []string
	var tag *structtag.Tag

	fields, err = List1stLevelFields(obj)
	if err != nil {
		return nil, err
	}

	fieldTagNames = make([]string, 0)
	for _, field := range fields {
		if field.Tags == nil {
			continue
		}
		tag, err = field.Tags.Get(tagType)
		if err != nil {
			continue
		}
		fieldTagNames = append(fieldTagNames, tag.Name)
	}

	return fieldTagNames, nil
}
