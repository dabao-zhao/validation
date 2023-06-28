package rules

import (
	"github.com/dabao-zhao/validation"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	s1 := "123"
	s2 := ""
	var time1 time.Time
	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, ""},
		{"t2", "", "the fortest field is required."},
		{"t3", &s1, ""},
		{"t4", &s2, "the fortest field is required."},
		{"t5", nil, "the fortest field is required."},
		{"t6", time1, "the fortest field is required."},
	}

	for _, test := range tests {
		r := Required
		err := r.Validate("fortest", test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestRequiredRule_Error(t *testing.T) {
	r := Required.Error("{{.field}} 不能为空")

	assert.Equal(t, validation.NewError("{{.field}} 不能为空"), r.err)

	err := r.Validate("fortest", "")
	assert.EqualError(t, err, "fortest 不能为空")
}

func assertError(t *testing.T, expected string, err error, tag string) {
	if expected == "" {
		assert.NoError(t, err, tag)
	} else {
		assert.EqualError(t, err, expected, tag)
	}
}
