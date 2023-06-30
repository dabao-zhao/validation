package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPv6(t *testing.T) {
	r := IPv6

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"t1", 123, "the IPv6 rule requires a type of string"},
		{"t2", "", ""},
		{"t3", nil, ""},
		{"t4", "ABCD:EF01:2345:6789:ABCD:EF01:2345:6789", ""},
		{"t5", "githubcom", "the fortest field must be a valid ipv6."},
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
