package rules

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
	"reflect"
)

var (
	ErrUrl = validation.NewError("the {{.field}} field must be a valid URL.")
)

type UrlRule struct {
	err validation.Error
}

// URL 验证字段必须为有效的 URL
var URL = UrlRule{
	err: ErrUrl,
}

func (r UrlRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the URL rule requires a type of string")
	}
	if !govalidator.IsURL(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r UrlRule) Error(message string) UrlRule {
	if r.err == nil {
		r.err = ErrUrl
	}
	r.err = r.err.SetMessage(message)
	return r
}
