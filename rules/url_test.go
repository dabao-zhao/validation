package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {

	r := URL

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, "the URL rule requires a type of string"},
		{"t2", "", ""},
		{"t3", nil, ""},
		{"t4", "http://github.com", ""},
		{"t5", "githubcom", "the fortest field must be a valid URL."},
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
