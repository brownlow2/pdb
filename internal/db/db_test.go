package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNew(t *testing.T, name string, headers []HeaderI) {
	db, err := New(name, headers, "Title")
	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.Equal(t, name, db.Name)
	assert.Equal(t, len(headers), len(db.Headers))

	h := make(map[HeaderI]struct{}, 0)
	for _, header := range headers {
		h[header] = struct{}{}
	}
	assert.True(t, reflect.DeepEqual(h, db.Headers))

	_, err = New(name, headers, "")
	assert.Error(t, err)
}

func testNewFail(t *testing.T, name string, headers []HeaderI) {
	db, err := New(name, headers, "Title")
	assert.True(t, reflect.DeepEqual(db, &DBImpl{}))
	assert.Error(t, err)

	db, err = New(name, headers, "")
	assert.True(t, reflect.DeepEqual(db, &DBImpl{}))
	assert.Error(t, err)
}

func TestNew(t *testing.T) {
	headers := []HeaderI{}
	testNewFail(t, "test db", headers)

	headers = []HeaderI{
		&Header{
			Name:      "Error",
			KeyHeader: true,
			Type:      VALUE_STRING,
		},
	}
	testNewFail(t, "test db", headers)

	headers = []HeaderI{
		&Header{
			Name:      "Title",
			KeyHeader: true,
			Type:      VALUE_STRING,
		},
	}
	testNew(t, "test db", headers)

	headers = []HeaderI{
		&Header{
			Name:      "Title",
			KeyHeader: true,
			Type:      VALUE_STRING,
		},
		&Header{
			Name:      "2",
			KeyHeader: false,
			Type:      VALUE_NUMBER,
		},
	}
	testNew(t, "test db", headers)
}

func TestGetName(t *testing.T) {
	db := &DBImpl{Name: "test db"}
	assert.Equal(t, "test db", db.GetName())
}

func TestGetKeyHeader(t *testing.T) {
	db := &DBImpl{Name: "test db", KeyHeader: "Test"}
	assert.Equal(t, "Test", db.GetKeyHeader())
}

func TestAddHeader(t *testing.T) {
	rows := []RowI{&Row{map[HeaderI]ValueI{}}}
	db := &DBImpl{"test", "Test", map[HeaderI]struct{}{}, &Rows{rows}}
	h := &Header{"Test", true, VALUE_STRING}
	hMap := map[HeaderI]struct{}{h: struct{}{}}
	db.AddHeader(h)
	assert.True(t, reflect.DeepEqual(hMap, db.Headers))

	h2 := &Header{"Test2", false, VALUE_NUMBER}
	hMap[h2] = struct{}{}
	db.AddHeader(h2)
	assert.True(t, reflect.DeepEqual(hMap, db.Headers))

	db.AddHeader(h)
	assert.True(t, reflect.DeepEqual(hMap, db.Headers))

	dbRow := db.GetRowFromKeyHeader("")
	assert.Equal(t, 2, len(dbRow.GetRowMap()))
}

func TestRemoveHeader(t *testing.T) {
	hToV := map[HeaderI]ValueI{
		&Header{"Test", true, VALUE_STRING}:   &Value{""},
		&Header{"Test2", false, VALUE_NUMBER}: &Value{""},
	}
	rows := []RowI{&Row{hToV}}
	db := &DBImpl{
		Name:      "test",
		KeyHeader: "Test",
		Headers: map[HeaderI]struct{}{
			&Header{"Test", true, VALUE_STRING}:   struct{}{},
			&Header{"Test2", false, VALUE_NUMBER}: struct{}{},
		},
		Rows: &Rows{rows},
	}
	hMap := map[HeaderI]struct{}{
		&Header{"Test", true, VALUE_STRING}:   struct{}{},
		&Header{"Test2", false, VALUE_NUMBER}: struct{}{},
	}

	err := db.RemoveHeader("Test")
	assert.Error(t, err)
	assert.Equal(t, len(hMap), len(db.Headers))

	hMap = map[HeaderI]struct{}{
		&Header{"Test", true, VALUE_STRING}: struct{}{},
	}

	err = db.RemoveHeader("Test2")
	assert.Nil(t, err)
	assert.Equal(t, len(hMap), len(db.Headers))

	dbRow := db.GetRowFromKeyHeader("")
	assert.Equal(t, 1, len(dbRow.GetRowMap()))
}

/*
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
*/
