package db

import (
	"errors"
	"fmt"
)

var (
	keyHeaderValueExistsError = "row with key header '%s' and value '%s' already exists"
)

/*
Test for:
  - The rows are returned correctly
*/
func (r *Rows) GetRows() []RowI {
	return r.Items
}

/*
Test for:
  - The row already exists
  - The row is added correctly
*/
func (r *Rows) AddRow(row RowI) error {
	h, v := row.GetKeyHeaderAndValue()
	for _, ro := range r.Items {
		if ro.KeyHeaderValueEqual(v.GetValue()) {
			return errors.New(fmt.Sprintf(keyHeaderValueExistsError, h.GetName(), v.GetValue()))
		}
	}

	r.Items = append(r.Items, row)

	return nil
}

/*
Test for:
  - The row is deleted correctly
*/
func (r *Rows) DeleteRow(row RowI) {
	_, v := row.GetKeyHeaderAndValue()
	for i, ro := range r.Items {
		_, val := ro.GetKeyHeaderAndValue()
		if v.GetValue() == val.GetValue() {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
		}
	}
}

/*
Test for:
  - The header is deleted correctly
  - The header is deleted in all rows
*/
func (r *Rows) RemoveHeader(header string) {
	for _, row := range r.Items {
		row.RemoveHeader(header)
	}
}

/*
Tesst for:
  - The correct KeyHeader is chosen
  - The value is added to the correct header
  - If the key is incorrect, nothing happens
*/
func (r *Rows) AddValueToRowWithKeyHeader(value string, header string, key string) {
	for _, row := range r.Items {
		_, v := row.GetKeyHeaderAndValue()
		if v.GetValue() == key {
			row.UpdateHeaderValue(header, value)
			break
		}
	}
}

/*
Test for:
  - The correct row is returned
  - If no value exists for the KeyHeader, return nothing
*/
func (r *Rows) GetRowFromKeyHeader(keyHeaderValue string) RowI {
	for _, row := range r.Items {
		_, v := row.GetKeyHeaderAndValue()
		if v.GetValue() == keyHeaderValue {
			return row
		}
	}

	return nil
}

/*
Test for:
  - Returns the correct rows given the value and header
*/
func (r *Rows) GetRowsFromHeaderAndValue(header string, value string) []RowI {
	rows := make([]RowI, 0)
	for _, row := range r.Items {
		if row.GetValueFromHeader(header).GetValue() == value {
			rows = append(rows, row)
		}
	}

	return rows
}
