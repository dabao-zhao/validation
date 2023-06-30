package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	var v = 1
	var v2 *int
	tests := []struct {
		tag    string
		values []interface{}
		value  interface{}
		err    string
	}{
		{"t0", []interface{}{1, 2}, 0, ""},
		{"t1", []interface{}{1, 2}, 1, ""},
		{"t2", []interface{}{1, 2}, 2, ""},
		{"t3", []interface{}{1, 2}, 3, "the fotest field must exist in [1 2]."},
		{"t4", []interface{}{}, 3, ""},
		{"t5", []interface{}{1, 2}, "1", "the fotest field must exist in [1 2]."},
		{"t6", []interface{}{1, 2}, &v, ""},
		{"t7", []interface{}{1, 2}, v2, ""},
		{"t8", []interface{}{[]byte{1}, 1, 2}, []byte{1}, ""},
	}

	for _, test := range tests {
		r := In(test.values...)
		err := r.Validate("fotest", test.value)
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}
