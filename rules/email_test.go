package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	r := Email

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, "the Email rule requires a type of string"},
		{"t2", "", ""},
		{"t3", nil, ""},
		{"t4", "dabao@github.com", ""},
		{"t5", "githubcom", "the fortest field must be a valid email."},
	}

	for _, test := range tests {
		err := r.Validate("fortest", test.value)
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}
