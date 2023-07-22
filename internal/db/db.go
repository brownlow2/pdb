package db

import (
	"errors"
	"fmt"
)

func New(name string, headers []string, keyHeader string) (*DBImpl, error) {
	if keyHeader == "" {
		return &DBImpl{}, errors.New("Key Header must not be empty")
	}

	db := &DBImpl{name, keyHeader, map[string]struct{}{}, Rows{}}
	for _, header := range headers {
		db.AddHeader(header)
	}
	return db, nil
}

func (db *DBImpl) GetName() string {
	return db.Name
}

func (db *DBImpl) AddHeader(header string, t Type) {
	// Don't need to add if it already exists
	if _, exists := db.Headers[header]; !exists {
		db.Headers[header] = struct{}{}

		// Add the header to each of the rows in the db
		for _, row := range db.Rows.Items {
			row.AddHeaderWithValue(header, false, t, "")
		}
	}
}

func (db *DBImpl) RemoveHeader(header string) error {
	if header == db.KeyHeader {
		return errors.New("cannot delete key header '", header, "'")
	}

	newHeaders := make(map[string]struct{}, 0)
	for h := range db.Headers {
		if h != header {
			newHeaders[h] = struct{}{}
		}
	}
	db.Headers = newHeaders
	db.Rows.RemoveHeader(header)

	return nil
}

func (db *DBImpl) AddRow(row map[string]string) error {
	if len(row) > len(db.Headers) {
		return errors.New("Too many headers")
	}

	db.Rows.AddRow(row)

	return nil
}

func (db *DBImpl) GetRows() []map[string]string {
	return db.Items
}

// Adds a value to a given header for a row with Title == name
func (db *DBImpl) AddValueToHeader(value string, header string, key string) error {
	if !db.headerExists(header) {
		headerNotExists := fmt.Sprintf("header %s does not exist", header)
		return errors.New(headerNotExists)
	}

	for i, row := range db.Rows.Items {
		for h, v := range row {
			if h == db.KeyHeader && v == key {
				db.Rows.Items[i][header] = value
				return nil
			}
		}
	}

	keyNotExists := fmt.Sprintf("key %s does not exist", key)
	return errors.New(keyNotExists)
}

func (db *DBImpl) headerExists(header string) bool {
	for h := range db.Headers {
		if h == header {
			return true
		}
	}

	return false
}

func (db *DBImpl) GetRowFromKeyHeader(value string) map[string]string {
	for _, r := range db.Rows.Items {
		if value == r[db.KeyHeader] {
			return r
		}
	}

	return map[string]string{}
}

func (db *DBImpl) GetRowsFromHeaderAndValue(header string, value string) []map[string]string {
	rows := make([]map[string]string, 0)

	if !db.headerExists(header) {
		return rows
	}

	for _, r := range db.Rows.Items {
		if value == r[header] {
			rows = append(rows, r)
		}
	}

	return rows
}

func (db *DBImpl) GetRowsFromHeaderAndValueLessThan(header string, value string) []map[string]string {
	rows := make([]map[string]string, 0)
	// First check if header exists
	if !db.headerExists(header) {
		return rows
	}

	// Check if the header is a number
	if 
}
