package rules

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrIp = validation.NewError("the {{.field}} field must be a valid ip.")
)

type IpRule struct {
	err validation.Error
}

// IP 验证字段必须为有效的 IP
var IP = IpRule{
	err: ErrIp,
}

func (r IpRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the IP rule requires a type of string")
	}
	if !govalidator.IsIP(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r IpRule) Error(message string) IpRule {
	if r.err == nil {
		r.err = ErrIp
	}
	r.err = r.err.SetMessage(message)
	return r
}
