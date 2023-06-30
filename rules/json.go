package rules

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrJson = validation.NewError("the {{.field}} field must be a valid json.")
)

type JsonRule struct {
	err validation.Error
}

// JSON 验证字段必须为有效的 JSON
var JSON = JsonRule{
	err: ErrJson,
}

func (r JsonRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the JSON rule requires a type of string")
	}
	if !govalidator.IsJSON(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r JsonRule) Error(message string) JsonRule {
	if r.err == nil {
		r.err = ErrJson
	}
	r.err = r.err.SetMessage(message)
	return r
}
