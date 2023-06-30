package rules

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrIpv6 = validation.NewError("the {{.field}} field must be a valid ipv6.")
)

type Ipv6Rule struct {
	err validation.Error
}

// IPv6 验证字段必须为有效的 IPv6
var IPv6 = Ipv6Rule{
	err: ErrIpv6,
}

func (r Ipv6Rule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(v) {
		return nil
	}
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("the IPv6 rule requires a type of string")
	}
	if !govalidator.IsIPv6(v.(string)) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}

	return nil
}

func (r Ipv6Rule) Error(message string) Ipv6Rule {
	if r.err == nil {
		r.err = ErrIpv6
	}
	r.err = r.err.SetMessage(message)
	return r
}
