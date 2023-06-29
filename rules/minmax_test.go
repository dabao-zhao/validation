package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	r := Max(10)

	assert.EqualError(t, r.Validate("fortest", 12), "the fortest field must not be greater than or equal to 10.")
	assert.Nil(t, r.Validate("fortest", 9))
	assert.Nil(t, r.Validate("fortest", []int{1, 2, 3, 4, 5, 6}))
}

func TestMin(t *testing.T) {
	r := Min(10)

	assert.EqualError(t, r.Validate("fortest", 7), "the fortest field must not be less than or equal to 10.")
	assert.Nil(t, r.Validate("fortest", 12))
	assert.Nil(t, r.Validate("fortest", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
