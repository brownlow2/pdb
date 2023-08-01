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

func TestGetHeaders(t *testing.T) {
	hExists := func(h string, headers []HeaderI) bool {
		for _, header := range headers {
			if header.GetName() == h {
				return true
			}
		}
		return false
	}

	db, _ := newDBWithValues()
	headers := db.GetHeaders()
	assert.NotNil(t, headers)
	assert.Equal(t, 2, len(headers))
	assert.True(t, hExists("Title", headers))
	assert.True(t, hExists("Value", headers))
}

func TestGetHeadersString(t *testing.T) {
	hExists := func(h string, headers []string) bool {
		for _, header := range headers {
			if header == h {
				return true
			}
		}
		return false
	}

	db, _ := newDBWithValues()
	headers := db.GetHeadersString()
	assert.NotNil(t, headers)
	assert.Equal(t, 2, len(headers))
	assert.True(t, hExists("Title (K)", headers))
	assert.True(t, hExists("Value", headers))
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

func TestDBAddRow(t *testing.T) {
	db := &DBImpl{
		Name:      "Test",
		KeyHeader: "Key",
		Headers: map[HeaderI]struct{}{
			&Header{"Key", true, VALUE_STRING}:     struct{}{},
			&Header{"NotKey", false, VALUE_NUMBER}: struct{}{},
		},
		Rows: &Rows{},
	}

	row := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", false, VALUE_STRING}: &Value{"key value"},
			// This header's KeyHeader value is intentionally set to true
			&Header{"NotKey", true, VALUE_NUMBER}: &Value{"not key value"},
		},
	}
	err := db.AddRow(row)
	assert.Error(t, err)

	row = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{""},
			&Header{"NotKey", false, VALUE_NUMBER}: &Value{"not key value"},
		},
	}
	err = db.AddRow(row)
	assert.Error(t, err)

	row = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{""},
			&Header{"NotKey", false, VALUE_NUMBER}: &Value{"not key value"},
			&Header{"Extra", false, VALUE_STRING}:  &Value{""},
		},
	}
	err = db.AddRow(row)
	assert.Error(t, err)

	row = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{"new key value"},
			&Header{"NotKey", false, VALUE_NUMBER}: &Value{"not key value"},
		},
	}
	err = db.AddRow(row)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(db.GetRows()))
}

func TestVerifyHeaders(t *testing.T) {
	db := &DBImpl{
		Name:      "Test",
		KeyHeader: "Key",
		Headers: map[HeaderI]struct{}{
			&Header{"Key", true, VALUE_STRING}:     struct{}{},
			&Header{"NotKey", false, VALUE_NUMBER}: struct{}{},
		},
		Rows: &Rows{},
	}

	row := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:     &Value{"new key value"},
			&Header{"NotKey", false, VALUE_NUMBER}: &Value{"not key value"},
		},
	}
	err := db.verifyHeaders(row)
	assert.Nil(t, err)

	row = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}:         &Value{"new key value"},
			&Header{"NotKey", false, VALUE_NUMBER}:     &Value{"not key value"},
			&Header{"NotPresent", false, VALUE_STRING}: &Value{"not exist"},
		},
	}
	err = db.verifyHeaders(row)
	assert.Error(t, err)

	row = &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Key", true, VALUE_STRING}: &Value{"new key value"},
		},
	}
	err = db.verifyHeaders(row)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(row.GetRowMap()))
}

func TestDBGetRows(t *testing.T) {
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
	assert.Nil(t, err)

	row := db.GetRowFromKeyHeader("test")
	assert.Equal(t, 2, len(row.GetRowMap()))
	v, err := row.GetValueFromHeader("Value")
	assert.Nil(t, err)
	assert.Equal(t, "after", v.GetValue())
}

func TestDBHeaderExists(t *testing.T) {
	db, _ := newDBWithValues()
	exists := db.headerExists("Title")
	assert.True(t, exists)

	exists = db.headerExists("Value")
	assert.True(t, exists)

	exists = db.headerExists("Not exists")
	assert.False(t, exists)
}

func TestDBGetRowFromKeyHeader(t *testing.T) {
	db, _ := newDBWithValues()
	row := db.GetRowFromKeyHeader("test")
	v, err := row.GetValueFromHeader("Title")
	assert.Nil(t, err)
	assert.Equal(t, "test", v.GetValue())
	v, err = row.GetValueFromHeader("Value")
	assert.Nil(t, err)
	assert.Equal(t, "test2", v.GetValue())

	row = db.GetRowFromKeyHeader("fail")
	assert.Nil(t, row)
}

func TestDBGetRowsFromHeaderAndValue(t *testing.T) {
	db, _ := newDBWithValues()
	r := &Row{
		RowMap: map[HeaderI]ValueI{
			&Header{"Title", true, VALUE_STRING}:  &Value{"next"},
			&Header{"Value", false, VALUE_STRING}: &Value{"test2"},
		},
	}
	err := db.AddRow(r)
	assert.Nil(t, err)

	rows, err := db.GetRowsFromHeaderAndValue("Value", "test2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows))

	rows, err = db.GetRowsFromHeaderAndValue("Value", "not exist")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(rows))

	rows, err = db.GetRowsFromHeaderAndValue("Not Exist", "")
	assert.Error(t, err)
	assert.Nil(t, rows)
}

func TestRemoveRow(t *testing.T) {
	db := &DBImpl{
		Name:      "Test",
		KeyHeader: "Key",
		Headers: map[HeaderI]struct{}{
			&Header{"Key", true, VALUE_STRING}:     struct{}{},
			&Header{"NotKey", false, VALUE_NUMBER}: struct{}{},
		},
		Rows: &Rows{
			Items: []RowI{
				&Row{
					RowMap: map[HeaderI]ValueI{
						&Header{"Key", true, VALUE_STRING}:     &Value{"key value"},
						&Header{"NotKey", false, VALUE_NUMBER}: &Value{"not key value"},
					},
				},
			},
		},
	}

	err := db.RemoveRow("")
	assert.Error(t, err)
	assert.Equal(t, 1, len(db.GetRows()))

	err = db.RemoveRow("key value")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(db.GetRows()))
}

func TestGetRowsFromHeaderAndValueNumberOperation(t *testing.T) {
	db := &DBImpl{
		Name:      "Test",
		KeyHeader: "Key",
		Headers: map[HeaderI]struct{}{
			&Header{"Key", true, VALUE_STRING}:     struct{}{},
			&Header{"NotKey", false, VALUE_NUMBER}: struct{}{},
		},
		Rows: &Rows{
			Items: []RowI{
				&Row{
					RowMap: map[HeaderI]ValueI{
						&Header{"Key", true, VALUE_STRING}:     &Value{"key value"},
						&Header{"NotKey", false, VALUE_NUMBER}: &Value{"3.4"},
					},
				},
				&Row{
					RowMap: map[HeaderI]ValueI{
						&Header{"Key", true, VALUE_STRING}:     &Value{"diff key value"},
						&Header{"NotKey", false, VALUE_NUMBER}: &Value{"5.0"},
					},
				},
				&Row{
					RowMap: map[HeaderI]ValueI{
						&Header{"Key", true, VALUE_STRING}:     &Value{"also key value"},
						&Header{"NotKey", false, VALUE_NUMBER}: &Value{"2"},
					},
				},
			},
		},
	}

	rows, err := db.GetRowsFromHeaderAndValueNumberOperation("NotKey", "4.0", "<")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows))

	rows, err = db.GetRowsFromHeaderAndValueNumberOperation("Key", "not number", "<")
	assert.Error(t, err)
	assert.Nil(t, rows)

	rows, err = db.GetRowsFromHeaderAndValueNumberOperation("NotKey", "4", "<")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows))

	rows, err = db.GetRowsFromHeaderAndValueNumberOperation("NotAKey", "not number", "<")
	assert.Error(t, err)
	assert.Nil(t, rows)

	rows, err = db.GetRowsFromHeaderAndValueNumberOperation("NotKey", "2.1", ">")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(rows))

	rows, err = db.GetRowsFromHeaderAndValueNumberOperation("NotKey", "1.9", ">")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(rows))
}
