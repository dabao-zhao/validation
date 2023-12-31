package rules

import (
	"testing"
	"time"

	"github.com/dabao-zhao/validation"
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
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}

func TestRequiredRule_Error(t *testing.T) {
	r := Required.Error("{{.field}} 不能为空")

	assert.Equal(t, validation.NewError("{{.field}} 不能为空"), r.err)

	err := r.Validate("fortest", "")
	assert.EqualError(t, err, "fortest 不能为空")
}
