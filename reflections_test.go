// Copyright (c) 2013 Théo Crevon
//
// See the file LICENSE for copying permission.

package reflections

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	unexported uint64
	Dummy      string `test:"dummytag"`
	Yummy      int    `test:"yummytag"`
}

func TestGetField_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	value, err := GetField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, value, "test")
}

func TestGetField_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
	}

	value, err := GetField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, value, "test")
}

func TestGetField_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetField(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	_, err := GetField(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetField_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	assert.Panics(t, func() {
		GetField(dummyStruct, "unexported") //nolint:errcheck,gosec
	})
}

func TestGetFieldKind_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := GetFieldKind(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.String)

	kind, err = GetFieldKind(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.Int)
}

func TestGetFieldKind_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := GetFieldKind(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.String)

	kind, err = GetFieldKind(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.Int)
}

func TestGetFieldKind_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldKind(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetFieldKind_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	_, err := GetFieldKind(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetFieldType_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	typeString, err := GetFieldType(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, typeString, "string")

	typeString, err = GetFieldType(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, typeString, "int")
}

func TestGetFieldType_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	typeString, err := GetFieldType(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, typeString, "string")

	typeString, err = GetFieldType(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, typeString, "int")
}

func TestGetFieldType_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldType(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetFieldType_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	_, err := GetFieldType(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetFieldTag_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{}

	tag, err := GetFieldTag(dummyStruct, "Dummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "dummytag")

	tag, err = GetFieldTag(dummyStruct, "Yummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "yummytag")
}

func TestGetFieldTag_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{}

	tag, err := GetFieldTag(dummyStruct, "Dummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "dummytag")

	tag, err = GetFieldTag(dummyStruct, "Yummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "yummytag")
}

func TestGetFieldTag_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := GetFieldTag(dummy, "Dummy", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{}

	_, err := GetFieldTag(dummyStruct, "obladioblada", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	_, err := GetFieldTag(dummyStruct, "unexported", "test")
	assert.Error(t, err)
}

func TestSetField_on_struct_with_valid_value_type(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "Dummy", "abc")
	assert.NoError(t, err)
	assert.Equal(t, dummyStruct.Dummy, "abc")
}

func TestSetField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "obladioblada", "life goes on")
	assert.Error(t, err)
}

func TestSetField_invalid_value_type(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := SetField(&dummyStruct, "Yummy", "123")
	assert.Error(t, err)
}

func TestSetField_non_exported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
	}

	assert.Error(t, SetField(&dummyStruct, "unexported", "fail, bitch"))
}

func TestFields_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := Fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestFields_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := Fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestFields_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Fields(dummy)
	assert.Error(t, err)
}

func TestFields_with_non_exported_fields(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 6789,
		Dummy:      "test",
		Yummy:      123,
	}

	fields, err := Fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestHasField_on_struct_with_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.True(t, has)
}

func TestHasField_on_struct_pointer_with_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.True(t, has)
}

func TestHasField_non_existing_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := HasField(dummyStruct, "Test")
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestHasField_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := HasField(dummy, "Test")
	assert.Error(t, err)
}

func TestHasField_unexported_field(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		unexported: 7890,
		Dummy:      "test",
		Yummy:      123,
	}

	has, err := HasField(dummyStruct, "unexported")
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestTags_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Tags(dummyStruct, "test")
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func TestTags_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Tags(dummyStruct, "test")
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func TestTags_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Tags(dummy, "test")
	assert.Error(t, err)
}

func TestItems_on_struct(t *testing.T) {
	t.Parallel()

	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Items(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func TestItems_on_struct_pointer(t *testing.T) {
	t.Parallel()

	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := Items(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func TestItems_on_non_struct(t *testing.T) {
	t.Parallel()

	dummy := "abc 123"

	_, err := Items(dummy)
	assert.Error(t, err)
}

//nolint:unused
func TestItems_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "Decumanus maximus"
	p.Number = 17

	items, err := Items(p)
	assert.NoError(t, err)
	itemsDeep, err := ItemsDeep(p)
	assert.NoError(t, err)

	assert.Equal(t, len(items), 2)
	assert.Equal(t, len(itemsDeep), 3)
	assert.Equal(t, itemsDeep["Name"], "John")
	assert.Equal(t, itemsDeep["Street"], "Decumanus maximus")
	assert.Equal(t, itemsDeep["Number"], 17)
}

//nolint:unused
func TestTags_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "Decumanus maximus"
	p.Number = 17

	tags, err := Tags(p, "tag")
	assert.NoError(t, err)
	tagsDeep, err := TagsDeep(p, "tag")
	assert.NoError(t, err)

	assert.Equal(t, len(tags), 2)
	assert.Equal(t, len(tagsDeep), 3)
	assert.Equal(t, tagsDeep["Name"], "bu")
	assert.Equal(t, tagsDeep["Street"], "be")
	assert.Equal(t, tagsDeep["Number"], "bi")
}

//nolint:unused
func TestFields_deep(t *testing.T) {
	t.Parallel()

	type Address struct {
		Street string `tag:"be"`
		Number int    `tag:"bi"`
	}

	type unexportedStruct struct{}

	type Person struct {
		Name string `tag:"bu"`
		Address
		unexportedStruct
	}

	p := Person{}
	p.Name = "John"
	p.Street = "street?"
	p.Number = 17

	fields, err := Fields(p)
	assert.NoError(t, err)
	fieldsDeep, err := FieldsDeep(p)
	assert.NoError(t, err)

	assert.Equal(t, len(fields), 2)
	assert.Equal(t, len(fieldsDeep), 3)
	assert.Equal(t, fieldsDeep[0], "Name")
	assert.Equal(t, fieldsDeep[1], "Street")
	assert.Equal(t, fieldsDeep[2], "Number")
}
