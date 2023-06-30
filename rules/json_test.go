package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	r := JSON

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, "the JSON rule requires a type of string"},
		{"t2", "", ""},
		{"t3", nil, ""},
		{"t4", "[1,2]", ""},
		{"t5", "[1,2,]", "the fortest field must be a valid json."},
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
