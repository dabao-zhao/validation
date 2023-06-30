package rules

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrEmail = validation.NewError("the {{.field}} field must be a valid email.")
)

type EmailRule struct {
	err validation.Error
}

// Email 验证字段必须为有效的 Email
var Email = EmailRule{
	err: ErrEmail,
}

func (r EmailRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the Email rule requires a type of string")
	}
	if !govalidator.IsEmail(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r EmailRule) Error(message string) EmailRule {
	if r.err == nil {
		r.err = ErrEmail
	}
	r.err = r.err.SetMessage(message)
	return r
}
