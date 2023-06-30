package rules

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
	"reflect"
)

var (
	ErrUppercase = validation.NewError("the {{.field}} field must be uppercase.")
)

type UppercaseRule struct {
	err validation.Error
}

// Uppercase 验证的字段必须是大写
var Uppercase = UppercaseRule{
	err: ErrUppercase,
}

func (r UppercaseRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the Uppercase rule requires a type of string")
	}
	if !govalidator.IsUpperCase(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r UppercaseRule) Error(message string) UppercaseRule {
	if r.err == nil {
		r.err = ErrUppercase
	}
	r.err = r.err.SetMessage(message)
	return r
}
