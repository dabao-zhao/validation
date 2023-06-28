package rules

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLength(t *testing.T) {
	tests := []struct {
		tag      string
		min, max int
		value    interface{}
		err      string
	}{
		{"t1", 2, 4, "abc", ""},
		{"t2", 2, 4, "", ""},
		{"t3", 2, 4, "abcdf", "the fortest field length must be between 2 and 4"},
		{"t4", 0, 4, "ab", ""},
		{"t5", 0, 4, "abcde", "the fortest field length must be between 0 and 4"},
		{"t6", 2, 0, "ab", "the length max must be more than min"},
		{"t6", 0, 0, "ab", "the length max must be more than min"},
		{"t6", 2, 2, "ab", "the length max must be more than min"},
		{"t9", 2, 10, 123, "cannot get the length of int"},
		{"t10", 2, 4, sql.NullString{String: "abc", Valid: true}, ""},
		{"t11", 2, 4, sql.NullString{String: "", Valid: true}, ""},
		{"t12", 2, 4, &sql.NullString{String: "abc", Valid: true}, ""},
	}

	for _, test := range tests {
		r := Length(test.min, test.max)
		err := r.Validate("fortest", test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestRuneLength(t *testing.T) {
	tests := []struct {
		tag      string
		min, max int
		value    interface{}
		err      string
	}{
		{"t1", 2, 4, "abc", ""},
		{"t2", 2, 4, "", ""},
		{"t3", 2, 4, "abcdf", "the fortest field length must be between 2 and 4"},
		{"t4", 0, 4, "ab", ""},
		{"t5", 0, 4, "abcde", "the fortest field length must be between 0 and 4"},
		{"t6", 2, 0, "ab", "the length max must be more than min"},
		{"t6", 0, 0, "ab", "the length max must be more than min"},
		{"t6", 2, 2, "ab", "the length max must be more than min"},
		{"t9", 2, 10, 123, "the RuneLength rule requires a type of string"},
		{"t10", 2, 4, sql.NullString{String: "abc", Valid: true}, ""},
		{"t11", 2, 4, sql.NullString{String: "", Valid: true}, ""},
		{"t12", 2, 4, &sql.NullString{String: "abc", Valid: true}, ""},
	}

	for _, test := range tests {
		r := RuneLength(test.min, test.max)
		err := r.Validate("fortest", test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestLengthRule_Error(t *testing.T) {
	r := Length(10, 100).Error("{{.field}} 长度应该介于 {{.min}} 和 {{.max}} 之间")
	err := r.Validate("fortest", "10")
	assert.EqualError(t, err, "fortest 长度应该介于 10 和 100 之间")
}
