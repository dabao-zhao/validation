package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotIn(t *testing.T) {
	v := 1
	var v2 *int
	var tests = []struct {
		tag    string
		values []interface{}
		value  interface{}
		err    string
	}{
		{"t0", []interface{}{1, 2}, 0, ""},
		{"t1", []interface{}{1, 2}, 1, "the fortest field must not exist in [1 2]."},
		{"t2", []interface{}{1, 2}, 2, "the fortest field must not exist in [1 2]."},
		{"t3", []interface{}{1, 2}, 3, ""},
		{"t4", []interface{}{}, 3, ""},
		{"t5", []interface{}{1, 2}, "1", ""},
		{"t6", []interface{}{1, 2}, &v, "the fortest field must not exist in [1 2]."},
		{"t7", []interface{}{1, 2}, v2, ""},
	}

	for _, test := range tests {
		r := NotIn(test.values...)
		err := r.Validate("fortest", test.value)
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}
