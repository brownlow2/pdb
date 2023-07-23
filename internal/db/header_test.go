package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderGetName(t *testing.T) {
	h := &Header{Name: "Test"}
	assert.Equal(t, "Test", h.GetName())
}

func TestGetType(t *testing.T) {
	h := &Header{Name: "Test", Type: VALUE_STRING}
	assert.Equal(t, VALUE_STRING, h.GetType())

	h = &Header{Name: "Test", Type: VALUE_NUMBER}
	assert.Equal(t, VALUE_NUMBER, h.GetType())
}

func TestIsString(t *testing.T) {
	h := &Header{Name: "Test", Type: VALUE_STRING}
	assert.True(t, h.IsString())

	h = &Header{Name: "Test", Type: VALUE_NUMBER}
	assert.False(t, h.IsString())
}

func TestIsNumber(t *testing.T) {
	h := &Header{Name: "Test", Type: VALUE_NUMBER}
	assert.True(t, h.IsNumber())

	h = &Header{Name: "Test", Type: VALUE_STRING}
	assert.False(t, h.IsNumber())
}

func TestIsKeyHeader(t *testing.T) {
	h := &Header{KeyHeader: true}
	assert.True(t, h.IsKeyHeader())

	h = &Header{KeyHeader: false}
	assert.False(t, h.IsKeyHeader())
}
