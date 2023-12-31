package db

import (
	"errors"
	"fmt"
	"strconv"
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

func (db *DBImpl) GetHeader(header string) HeaderI {
	for h := range db.Headers {
		if h.GetName() == header {
			return h
		}
	}

	return &Header{}
}

func (db *DBImpl) GetHeaders() []HeaderI {
	headers := []HeaderI{}
	for h := range db.Headers {
		headers = append(headers, h)
	}

	return headers
}

func (db *DBImpl) GetHeadersString() []string {
	headersString := []string{}
	for h := range db.Headers {
		if h.IsKeyHeader() {
			headersString = append(headersString, fmt.Sprintf("%s (K)", h.GetName()))
		} else {
			headersString = append(headersString, h.GetName())
		}
	}

	return headersString
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

func (db *DBImpl) RemoveRow(keyValue string) error {
	if keyValue == "" {
		return errors.New(keyValueEmptyError)
	}

	db.Rows.DeleteRowWithValue(keyValue)
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

func (db *DBImpl) GetRowsFromHeaderAndValueNumberOperation(header string, value string, op string) ([]RowI, error) {
	valueF, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(notANumberError, value))
	}

	if !db.headerExists(header) {
		return nil, errors.New(fmt.Sprintf(headerNotExistError, header))
	}

	h := db.GetHeader(header)
	rows := make([]RowI, 0)
	for _, row := range db.GetRows() {
		// Don't need to check error, header definitely exists
		v, _ := row.GetValueFromHeader(header)
		vF, err := h.Number(v)
		if err != nil {
			return nil, err
		}

		if op == "<" && vF < valueF {
			rows = append(rows, row)
		}

		if op == ">" && vF > valueF {
			rows = append(rows, row)
		}
	}

	return rows, nil
}
