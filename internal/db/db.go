package db

import (
	"errors"
	"fmt"
)

var (
	keyHeaderIncorrect  = "key header '%s' incorrect, expected '%s'"
	keyHeaderEmptyError = "key header '%s' must not be empty"
	headerNotExistError = "header '%s' does not exist"
)

func New(name string, headers []HeaderI, keyHeader string) (*DBImpl, error) {
	if keyHeader == "" {
		return &DBImpl{}, errors.New("key header must exist and not be empty")
	}

	db := &DBImpl{name, keyHeader, map[HeaderI]struct{}{}, &Rows{}}
	for _, header := range headers {
		db.AddHeader(header)
	}

	if len(headers) == 0 || !db.headerExists(keyHeader) {
		return &DBImpl{}, errors.New(fmt.Sprintf(keyHeaderEmptyError, keyHeader))
	}

	return db, nil
}

func (db *DBImpl) GetName() string {
	return db.Name
}

func (db *DBImpl) GetKeyHeader() string {
	return db.KeyHeader
}

func (db *DBImpl) AddHeader(header HeaderI) {
	// Don't need to add if it already exists
	if !db.headerExists(header.GetName()) {
		db.Headers[header] = struct{}{}

		// Add the header to each of the rows in the db
		for _, row := range db.Rows.GetRows() {
			row.AddHeaderWithValue(header.GetName(), header.IsKeyHeader(), header.GetType(), "")
		}
	}
}

func (db *DBImpl) RemoveHeader(header string) error {
	if header == db.KeyHeader {
		return errors.New(fmt.Sprintf(headerNotExistError, header))
	}

	newHeaders := make(map[HeaderI]struct{}, 0)
	for h := range db.Headers {
		if h.GetName() != header {
			newHeaders[h] = struct{}{}
		}
	}
	db.Headers = newHeaders
	db.Rows.RemoveHeader(header)

	return nil
}

func (db *DBImpl) AddRow(row RowI) error {
	// Make sure the row's key header is correct
	h, v := row.GetKeyHeaderAndValue()
	if h.GetName() != db.GetKeyHeader() {
		return errors.New(fmt.Sprintf(keyHeaderIncorrect, h.GetName(), db.GetKeyHeader()))
	}

	// Make sure the key header's value is not empty
	if v.GetValue() == "" {
		return errors.New(fmt.Sprintf(keyHeaderEmptyError, h.GetName()))
	}

	// Verify that there are no extra headers in the row
	err := db.verifyHeaders(row)
	if err != nil {
		return err
	}

	err = db.Rows.AddRow(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBImpl) verifyHeaders(row RowI) error {
	// Verify that no extra headers exist in row, and if any are missing add them
	// as new empty values
	for h := range row.GetRowMap() {
		if !db.headerExists(h.GetName()) {
			return errors.New(fmt.Sprintf(headerNotExistError, h.GetName()))
		}
	}

	// Add any headers to row that should exist
	for h := range db.Headers {
		if !row.HeaderExists(h.GetName()) {
			// KeyHeader will already be in row so can just use false here
			row.AddHeaderWithValue(h.GetName(), false, h.GetType(), "")
		}
	}

	return nil
}

func (db *DBImpl) GetRows() []RowI {
	return db.Rows.GetRows()
}

// Adds a value to a given header for a row with KeyHeader == key
func (db *DBImpl) AddValueToHeader(value string, header string, key string) error {
	if !db.headerExists(header) {
		return errors.New(fmt.Sprintf(headerNotExistError, header))
	}

	db.Rows.AddValueToRowWithKeyHeader(value, header, key)

	return nil
}

func (db *DBImpl) headerExists(header string) bool {
	for h := range db.Headers {
		if h.GetName() == header {
			return true
		}
	}
	return false
}

func (db *DBImpl) GetRowFromKeyHeader(value string) RowI {
	return db.Rows.GetRowFromKeyHeader(value)
}

func (db *DBImpl) GetRowsFromHeaderAndValue(header string, value string) ([]RowI, error) {
	if !db.headerExists(header) {
		return nil, errors.New(fmt.Sprintf(headerNotExistError, header))
	}

	rows, err := db.Rows.GetRowsFromHeaderAndValue(header, value)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// TODO: add GetRowsFromHeaderAndValueLessThan
