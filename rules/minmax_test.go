package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	r := Max(3)

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"numericSuccess", 1, ""},
		{"numericError", 4, "the fortest field must be less than or equal to 3."},
		{"sliceSuccess", []int{1, 2, 3}, ""},
		{"sliceError", []int{1, 2, 3, 4, 5, 6}, "the fortest field must be less than or equal to 3 items."},
		{"stringSuccess", "1", ""},
		{"stringError", "1234", "the fortest field must be less than or equal to 3 characters."},
		{"mapSuccess", map[string]string{"a": "a"}, ""},
		{"mapError", map[string]string{
			"a": "a",
			"b": "b",
			"c": "c",
			"d": "d",
		}, "the fortest field must be less than or equal to 3 items."},
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

func TestMin(t *testing.T) {
	r := Min(3)

	tests := []struct {
		tag   string
		value interface{}
		err   string
	}{
		{"numericSuccess", 4, ""},
		{"numericError", 2, "the fortest field must be greeter than or equal to 3."},
		{"sliceSuccess", []int{1, 2, 3, 4}, ""},
		{"sliceError", []int{1, 2}, "the fortest field must be greeter than or equal to 3 items."},
		{"stringSuccess", "123", ""},
		{"stringError", "12", "the fortest field must be greeter than or equal to 3 characters."},
		{"mapSuccess", map[string]string{"a": "a", "b": "b", "c": "c", "d": "d"}, ""},
		{"mapError", map[string]string{
			"a": "a",
			"b": "b",
		}, "the fortest field must be greeter than or equal to 3 items."},
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

func TestError(t *testing.T) {
	r := Min(1.223)
	assert.EqualError(t, r.Validate("f", 1), "cannot convert float64 to int64")
}
