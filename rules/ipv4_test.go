package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPv4(t *testing.T) {
	r := IPv4

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, "the IPv4 rule requires a type of string"},
		{"t2", "", ""},
		{"t3", nil, ""},
		{"t4", "127.0.0.1", ""},
		{"t5", "githubcom", "the fortest field must be a valid ipv4."},
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
