package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRows() (RowsI, RowI, map[HeaderI]ValueI) {
	row, rowMap := createRow()
	rows := &Rows{
		Items: []RowI{row},
	}

	return rows, row, rowMap
}

func TestGetRows(t *testing.T) {
	rows, row, _ := createRows()

	assert.True(t, reflect.DeepEqual([]RowI{row}, rows.GetRows()))
	assert.Equal(t, 1, len(rows.GetRows()))
}

func TestAddRow(t *testing.T) {
	rows, row, _ := createRows()
	r := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{"key value"},
			&Header{"NotKey", false, VALUE_STRING}: &Value{"not key value"},
		},
	}

	err := rows.AddRow(r)
	assert.Error(t, err)
	assert.True(t, reflect.DeepEqual([]RowI{row}, rows.GetRows()))
	assert.Equal(t, 1, len(rows.GetRows()))

	r = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{"diff key value"},
			&Header{"NotKey", false, VALUE_STRING}: &Value{"diff not key value"},
		},
	}

	err = rows.AddRow(r)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows.GetRows()))
}

func TestDeleteRow(t *testing.T) {
	rows, _, _ := createRows()
	r := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:    &Value{"diff key value"},
			&Header{"NotKey", true, VALUE_STRING}: &Value{"diff not key value"},
		},
	}
	err := rows.AddRow(r)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows.GetRows()))

	rows.DeleteRow(r)
	assert.Equal(t, 1, len(rows.GetRows()))
}

func TestRowsRemoveHeader(t *testing.T) {
	rows, _, _ := createRows()
	rows.RemoveHeader("NotKey")
	r := rows.GetRows()
	for _, row := range r {
		assert.Equal(t, 1, len(row.GetRowMap()))
	}
}

func TestAddValueToRowWithKeyHeader(t *testing.T) {
	rows, _, _ := createRows()
	r := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:    &Value{"diff key value"},
			&Header{"NotKey", true, VALUE_STRING}: &Value{"diff not key value"},
		},
	}
	err := rows.AddRow(r)
	assert.Nil(t, err)
	rows.AddValueToRowWithKeyHeader("new not key value", "NotKey", "key value")
	row := rows.GetRowFromKeyHeader("key value")
	v, err := row.GetValueFromHeader("NotKey")
	assert.Nil(t, err)
	assert.Equal(t, "new not key value", v.GetValue())

	row2 := rows.GetRowFromKeyHeader("diff key value")
	v, err = row2.GetValueFromHeader("NotKey")
	assert.Nil(t, err)
	assert.Equal(t, "diff not key value", v.GetValue())
}

func TestGetRowFromKeyHeader(t *testing.T) {
	rows, row, _ := createRows()
	assert.True(t, reflect.DeepEqual(row, rows.GetRowFromKeyHeader("key value")))
	assert.Nil(t, rows.GetRowFromKeyHeader("not key value"))
}

func TestGetRowsFromHeaderAndValue(t *testing.T) {
	rows, row, _ := createRows()
	r, err := rows.GetRowsFromHeaderAndValue("NotExist", "")
	assert.Nil(t, r)
	assert.Error(t, err)

	r, err = rows.GetRowsFromHeaderAndValue("Key", "not exist")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(r))

	r, err = rows.GetRowsFromHeaderAndValue("NotKey", "")
	assert.Nil(t, err)
	assert.Equal(t, []RowI{row}, r)

	newRow := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:    &Value{"diff key value"},
			&Header{"NotKey", true, VALUE_STRING}: &Value{""},
		},
	}
	err = rows.AddRow(newRow)
	assert.Nil(t, err)

	r, err = rows.GetRowsFromHeaderAndValue("NotKey", "")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(r))
}
