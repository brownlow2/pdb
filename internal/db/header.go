package db

import (
	"errors"
	"fmt"
	"strconv"
)

func (h *Header) GetName() string {
	return h.Name
}

func (h *Header) GetType() Type {
	return h.Type
}

func (h *Header) IsString() bool {
	return h.Type == VALUE_STRING
}

func (h *Header) IsNumber() bool {
	return h.Type == VALUE_NUMBER
}

func (h *Header) IsKeyHeader() bool {
	return h.KeyHeader
}

func (h *Header) Number(value ValueI) (float64, error) {
	if h.IsString() {
		err := fmt.Sprintf(notANumberError, value.GetValue())
		return 0.0, errors.New(err)
	}

	return strconv.ParseFloat(value.GetValue(), 64)
}
