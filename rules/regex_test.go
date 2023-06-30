package rules

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	var v2 *string
	tests := []struct {
		tag   string
		re    string
		value interface{}
		err   string
	}{
		{"t1", "[a-z]+", "abc", ""},
		{"t2", "[a-z]+", "", ""},
		{"t3", "[a-z]+", v2, ""},
		{"t4", "[a-z]+", "123", "the fortest field format is invalid."},
		{"t5", "[a-z]+", []byte("abc"), ""},
		{"t6", "[a-z]+", []byte("123"), "the fortest field format is invalid."},
		{"t7", "[a-z]+", []byte(""), ""},
		{"t8", "[a-z]+", nil, ""},
		{"t9", "[a-z]+", map[string]string{"1": "1"}, "the fortest field format is invalid."},
	}

	for _, test := range tests {
		r := Regex(regexp.MustCompile(test.re))
		err := r.Validate("fortest", test.value)
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}
