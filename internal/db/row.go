package db

import (
	"errors"
	"fmt"
)

var (
	keyHeaderAlreadyExistsError = "key header already exists"
	deleteKeyHeaderError        = "cannot delete key header '%s'"
)

func (r *Row) GetKeyHeaderAndValue() (HeaderI, ValueI) {
	for h, v := range r.RowMap {
		if h.IsKeyHeader() {
			return h, v
		}
	}

	return nil, nil
}

func (r *Row) GetValueFromHeader(header string) (ValueI, error) {
	for h, v := range r.RowMap {
		if h.GetName() == header {
			return v, nil
		}
	}

	return nil, errors.New(fmt.Sprintf(headerNotExistError, header))
}

func (r *Row) GetRowMap() map[HeaderI]ValueI {
	return r.RowMap
}

func (r *Row) HeaderExists(header string) bool {
	for h := range r.RowMap {
		if h.GetName() == header {
			return true
		}
	}

	return false
}

func (r *Row) KeyHeaderValueEqual(value string) bool {
	for h, v := range r.RowMap {
		if h.IsKeyHeader() {
			if v.GetValue() == value {
				return true
			}
		}
	}

	return false
}

func (r *Row) AddHeaderWithValue(header string, keyHeader bool, t Type, value string) error {
	key, _ := r.GetKeyHeaderAndValue()
	if key != nil && keyHeader {
		return errors.New(keyHeaderAlreadyExistsError)
	}

	h := &Header{header, keyHeader, t}
	v := &Value{value}
	r.RowMap[h] = v
	return nil
}

func (r *Row) RemoveHeader(header string) error {
	// Return error if trying to delete key header
	h, _ := r.GetKeyHeaderAndValue()
	if h.GetName() == header {
		return errors.New(fmt.Sprintf(deleteKeyHeaderError, h.GetName()))
	}

	newRowMap := map[HeaderI]ValueI{}
	for h, v := range r.RowMap {
		if h.GetName() != header {
			newRowMap[h] = v
		}
	}
	r.RowMap = newRowMap

	return nil
}

func (r *Row) UpdateHeaderValue(header string, value string) {
	for h := range r.RowMap {
		if h.GetName() == header {
			r.RowMap[h].SetValue(value)
			break
		}
	}
}
