package db

import (
	"errors"
	"fmt"
)

var (
	keyHeaderValueExistsError = "row with key header '%s' and value '%s' already exists"
)

func (r *Rows) AddRow(row RowI) error {
	h, v := row.GetKeyHeaderAndValue()
	for _, ro := range r.Items {
		if ro.KeyHeaderValueEqual(v.GetValue()) {
			err := fmt.Sprintf(keyHeaderValueExistsError, h.GetName(), v.GetValue())
			return errors.New(err)
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
