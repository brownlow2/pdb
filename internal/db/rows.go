package db

import (
	"errors"
	"fmt"
)

func (r *Rows) GetRows() []RowI {
	return r.Items
}

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

func (r *Rows) DeleteRow(row RowI) {
	_, v := row.GetKeyHeaderAndValue()
	for i, ro := range r.Items {
		_, val := ro.GetKeyHeaderAndValue()
		if v.GetValue() == val.GetValue() {
			r.Items = append(r.Items[:i], r.Items[i+1:]...)
		}
	}
}

func (r *Rows) RemoveHeader(header string) {
	for _, row := range r.Items {
		row.RemoveHeader(header)
	}
}

func (r *Rows) AddValueToRowWithKeyHeader(value string, header string, key string) {
	for _, row := range r.Items {
		_, v := row.GetKeyHeaderAndValue()
		if v.GetValue() == key {
			row.UpdateHeaderValue(header, value)
			break
		}
	}
}

func (r *Rows) GetRowFromKeyHeader(keyHeaderValue string) RowI {
	for _, row := range r.Items {
		_, v := row.GetKeyHeaderAndValue()
		if v.GetValue() == keyHeaderValue {
			return row
		}
	}

	return nil
}

func (r *Rows) GetRowsFromHeaderAndValue(header string, value string) ([]RowI, error) {
	rows := make([]RowI, 0)
	for _, row := range r.Items {
		v, err := row.GetValueFromHeader(header)
		if err != nil {
			return nil, err
		}

		if v.GetValue() == value {
			rows = append(rows, row)
		}
	}

	return rows, nil
}
