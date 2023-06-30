package rules

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrIpv4 = validation.NewError("the {{.field}} field must be a valid ipv4.")
)

type Ipv4Rule struct {
	err validation.Error
}

// IPv4 验证字段必须为有效的 IPv4
var IPv4 = Ipv4Rule{
	err: ErrIpv4,
}

func (r Ipv4Rule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the IPv4 rule requires a type of string")
	}
	if !govalidator.IsIPv4(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r Ipv4Rule) Error(message string) Ipv4Rule {
	if r.err == nil {
		r.err = ErrIpv4
	}
	r.err = r.err.SetMessage(message)
	return r
}
