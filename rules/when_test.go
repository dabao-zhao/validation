package rules

import (
	"testing"

	"github.com/dabao-zhao/validation"
	"github.com/stretchr/testify/assert"
)

func TestWhen(t *testing.T) {

	r1 := Required
	r2 := Length(5, 10)

	tests := []struct {
		tag       string
		condition bool
		value     interface{}
		rule      validation.Rule
		elseRule  validation.Rule
		err       string
	}{
		// True condition
		{"t1", true, "1", r1, r2, ""},
		{"t2", false, "1", r1, r2, "the fortest field length must be between 5 and 10."},
		{"t3", true, "", r1, r2, "the fortest field is required."},
		{"t4", false, "", r1, r2, "the fortest field length must be between 5 and 10."},
		{"t1", false, "1", r1, nil, ""},
	}

	for _, test := range tests {
		err := When(test.condition, test.rule).Else(test.elseRule).Validate("fortest", test.value)
		if test.err == "" {
			assert.NoError(t, err, test.tag)
		} else {
			assert.EqualError(t, err, test.err, test.tag)
		}
	}
}
