package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNew(t *testing.T, name string, headers []string) {
	db, err := New(name, headers, "Title")
	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.Equal(t, name, db.Name)
	assert.Equal(t, len(headers), len(db.Headers))

	h := make(map[string]struct{}, 0)
	for _, header := range headers {
		h[header] = struct{}{}
	}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	_, err = New(name, headers, "")
	assert.Error(t, err)
}

func TestNew(t *testing.T) {
	testNew(t, "test db", []string{})
	testNew(t, "test db", []string{"header1"})
	testNew(t, "test db", []string{"header1", "header2"})
}

func TestGetName(t *testing.T) {
	db := &DBImpl{Name: "test db"}
	assert.Equal(t, "test db", db.GetName())
}

func TestAddHeader(t *testing.T) {
	db := &DBImpl{"test", "Title", map[string]struct{}{}, Rows{}}
	db.AddHeader("header1")
	h := map[string]struct{}{"header1": struct{}{}}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	db.AddHeader("header2")
	h["header2"] = struct{}{}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	db.AddHeader("header1")
	assert.True(t, reflect.DeepEqual(h, db.Headers))
}

func TestRemoveHeader(t *testing.T) {
	db := &DBImpl{
		Name: "test",
		Headers: map[string]struct{}{
			"header1": struct{}{},
			"header2": struct{}{},
		},
	}

	db.RemoveHeader("header1")
	h := map[string]struct{}{"header2": struct{}{}}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	db.RemoveHeader("header2")
	h = map[string]struct{}{}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	db.RemoveHeader("header3")
	assert.True(t, reflect.DeepEqual(h, db.Headers))
}

func TestAddRow(t *testing.T) {
	db := &DBImpl{"Test", "Key", map[string]struct{}{}, Rows{}}
	row := map[string]string{"test": "test"}
	err := db.AddRow(row)
	assert.Error(t, err)

	headers := map[string]struct{}{"test": struct{}{}}
	db.Headers = headers
	err = db.AddRow(row)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(row, db.Rows.Items[0]))
}

func newDBWithValues() (*DBImpl, []map[string]string) {
	rows := Rows{
		Items: []map[string]string{
			{
				"Title": "test",
				"Value": "test",
			},
		},
	}
	db := &DBImpl{
		Name:      "test",
		KeyHeader: "Title",
		Headers: map[string]struct{}{
			"Title": struct{}{},
			"Value": struct{}{},
		},
		Rows: rows,
	}
	return db, rows.Items
}

func TestGetRows(t *testing.T) {
	db, rows := newDBWithValues()

	assert.True(t, reflect.DeepEqual(rows, db.GetRows()))
}

func TestAddValueToHeader(t *testing.T) {
	db, _ := newDBWithValues()
	err := db.AddValueToHeader("after", "Value", "test")
	assert.Nil(t, err)

	err = db.AddValueToHeader("after", "not exist", "test")
	assert.Error(t, err)

	err = db.AddValueToHeader("after", "Value", "not exists")
	assert.Error(t, err)
}

func TestHeaderExists(t *testing.T) {
	db, _ := newDBWithValues()
	exists := db.headerExists("Title")
	assert.True(t, exists)

	exists = db.headerExists("Value")
	assert.True(t, exists)

	exists = db.headerExists("Not exists")
	assert.False(t, exists)
}

func TestGetRowFromKeyHeader(t *testing.T) {
	db, _ := newDBWithValues()
	row := db.GetRowFromKeyHeader("test")
	expectedRow := map[string]string{"Title": "test", "Value": "test"}
	assert.True(t, reflect.DeepEqual(row, expectedRow))

	row = db.GetRowFromKeyHeader("fail")
	expectedRow = map[string]string{}
	assert.True(t, reflect.DeepEqual(row, expectedRow))
}

func TestGetRowsFromHeaderAndValue(t *testing.T) {
	db, _ := newDBWithValues()
	expectedRows := []map[string]string{
		{
			"Title": "test",
			"Value": "test",
		},
		{
			"Title": "test2",
			"Value": "test",
		},
	}
	db.Rows.Items = expectedRows
	rows := db.GetRowsFromHeaderAndValue("Value", "test")
	assert.True(t, reflect.DeepEqual(rows, expectedRows))

	rows = db.GetRowsFromHeaderAndValue("Values", "test")
	expectedRows = make([]map[string]string, 0)
	assert.True(t, reflect.DeepEqual(rows, expectedRows))

	rows = db.GetRowsFromHeaderAndValue("Value", "fail")
	assert.True(t, reflect.DeepEqual(rows, expectedRows))
}
