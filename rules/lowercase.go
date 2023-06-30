package rules

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrLowercase = validation.NewError("the {{.field}} field must be lowercase.")
)

type LowercaseRule struct {
	err validation.Error
}

// Lowercase 验证的字段必须是小写
var Lowercase = LowercaseRule{
	err: ErrLowercase,
}

func (r LowercaseRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the Lowercase rule requires a type of string")
	}
	if !isLower(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r LowercaseRule) Error(message string) LowercaseRule {
	if r.err == nil {
		r.err = ErrLowercase
	}
	r.err = r.err.SetMessage(message)
	return r
}

func isLower(s string) bool {
	match, _ := regexp.MatchString("^[a-z0-9]+$", s)
	return match
}
