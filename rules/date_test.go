package rules

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	r := Date

	assert.Nil(t, r.Validate("fortest", time.Now()))
	assert.Nil(t, r.Validate("fortest", "2023-06-28 17:44:04"))
	assert.Nil(t, r.SetLayout("2006-01-02").Validate("fortest", "2023-06-28"))

	assert.EqualError(t, r.Validate("fortest", "2023"), "the fortest field must be a valid date.")
}

func TestDateEqual(t *testing.T) {
	tim := time.Now()
	r := DateEqual(tim)

	assert.Nil(t, r.Validate("fortest", tim))
	assert.EqualError(t, r.Validate("fortest", tim.Add(time.Second)), "the fortest field must be a date equal to "+tim.Format(r.layout)+".")
}

func TestDateBefore(t *testing.T) {
	tim := time.Now()
	r := DateBefore(tim)

	assert.Nil(t, r.Validate("fortest", tim.AddDate(-1, 0, 0)))
	assert.EqualError(t, r.Validate("fortest", tim), "the fortest field must be a date before "+tim.Format(r.layout)+".")
}

func TestDateAfter(t *testing.T) {
	tim := time.Now()
	r := DateAfter(tim)

	assert.Nil(t, r.Validate("fortest", tim.AddDate(1, 0, 0)))
	assert.EqualError(t, r.Validate("fortest", tim), "the fortest field must be a date after "+tim.Format(r.layout)+".")
}
