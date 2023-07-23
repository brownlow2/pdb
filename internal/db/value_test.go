package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	v := &Value{"test"}
	assert.Equal(t, "test", v.GetValue())
	assert.NotEqual(t, "incorrect", v.GetValue())
}

func TestSetValue(t *testing.T) {
	v := &Value{"test"}
	v.SetValue("changed")
	assert.Equal(t, "changed", v.GetValue())
	assert.NotEqual(t, "test", v.GetValue())
}
