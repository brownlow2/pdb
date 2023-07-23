package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRow() (RowI, map[HeaderI]ValueI) {
	rowMap := map[HeaderI]ValueI{
		&Header{"Key", true, VALUE_STRING}:     &Value{"key value"},
		&Header{"NotKey", false, VALUE_NUMBER}: &Value{""},
	}

	return &Row{
		rowMap,
	}, rowMap
}

func TestGetKeyHeaderAndValue(t *testing.T) {
	row, _ := createRow()

	h, v := row.GetKeyHeaderAndValue()
	assert.NotNil(t, h)
	assert.NotNil(t, v)
	assert.Equal(t, "Key", h.GetName())
	assert.Equal(t, "key value", v.GetValue())
}

func TestGetValueFromHeader(t *testing.T) {
	row, _ := createRow()

	v := row.GetValueFromHeader("Key")
	assert.NotNil(t, v)
	assert.Equal(t, "key value", v.GetValue())

	v = row.GetValueFromHeader("NotKey")
	assert.NotNil(t, v)
	assert.Equal(t, "", v.GetValue())

	v = row.GetValueFromHeader("NotPresent")
	assert.Nil(t, v)
}

func TestGetRowMap(t *testing.T) {
	row, rowMap := createRow()

	dbRowMap := row.GetRowMap()
	assert.NotNil(t, dbRowMap)
	assert.True(t, reflect.DeepEqual(rowMap, dbRowMap))
}

func TestHeaderExists(t *testing.T) {
	row, _ := createRow()

	assert.True(t, row.HeaderExists("Key"))
	assert.True(t, row.HeaderExists("NotKey"))
	assert.False(t, row.HeaderExists("NotPresent"))
}

func TestKeyHeaderValueEqual(t *testing.T) {
	row, _ := createRow()

	assert.True(t, row.KeyHeaderValueEqual("key value"))
	assert.False(t, row.KeyHeaderValueEqual(""))
}

func TestAddHeaderWithValue(t *testing.T) {
	row, _ := createRow()

	err := row.AddHeaderWithValue("New", false, VALUE_STRING, "new value")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(row.GetRowMap()))
	v := row.GetValueFromHeader("New")
	assert.Equal(t, "new value", v.GetValue())

	err = row.AddHeaderWithValue("NewKey", true, VALUE_STRING, "new key value")
	assert.Error(t, err)
	assert.Equal(t, 3, len(row.GetRowMap()))
	v = row.GetValueFromHeader("NewKey")
	assert.Nil(t, v)
}

func TestUpdateHeaderValue(t *testing.T) {
	row, _ := createRow()

	row.UpdateHeaderValue("Key", "newer value")
	assert.Equal(t, "newer value", row.GetValueFromHeader("Key").GetValue())
	assert.Equal(t, "", row.GetValueFromHeader("NotKey").GetValue())

	row.UpdateHeaderValue("NotKey", "something")
	assert.Equal(t, "newer value", row.GetValueFromHeader("Key").GetValue())
	assert.Equal(t, "something", row.GetValueFromHeader("NotKey").GetValue())

	row.UpdateHeaderValue("NotPresent", "shouldn't exist")
	assert.Equal(t, "newer value", row.GetValueFromHeader("Key").GetValue())
	assert.Equal(t, "something", row.GetValueFromHeader("NotKey").GetValue())
	assert.Equal(t, 2, len(row.GetRowMap()))
}

func TestRowRemoveHeader(t *testing.T) {
	row, _ := createRow()

	err := row.RemoveHeader("Key")
	assert.Error(t, err)
	assert.Equal(t, 2, len(row.GetRowMap()))

	err = row.RemoveHeader("NotKey")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(row.GetRowMap()))

	err = row.RemoveHeader("NotPresent")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(row.GetRowMap()))
}
