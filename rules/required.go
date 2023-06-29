package rules

import (
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrRequired = validation.NewError("the {{.field}} field is required.")
)

type RequiredRule struct {
	err validation.Error
}

// Required 验证非空.
var Required = RequiredRule{
	err: ErrRequired,
}

func (r RequiredRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}
	return nil
}

func (r RequiredRule) Error(message string) RequiredRule {
	if r.err == nil {
		r.err = ErrRequired
	}
	r.err = r.err.SetMessage(message)
	return r
}
