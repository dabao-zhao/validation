package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowercaseRule(t *testing.T) {
	r := Lowercase

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"success", "abc", ""},
		{"error", "ABC", "the fortest field must be lowercase."},
		{"notString1", []int{1, 2, 3, 4}, "the Lowercase rule requires a type of string"},
		{"noteString2", 1, "the Lowercase rule requires a type of string"},
		{"numeric", "123", ""},
	}
	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			if tt.err == "" {
				assert.Nil(t, r.Validate("fortest", tt.value))
			} else {
				assert.EqualError(t, r.Validate("fortest", tt.value), tt.err)
			}
		})
	}
}
