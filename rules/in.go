package rules

import (
	"reflect"

	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrIn = validation.NewError("the {{.field}} field must exist in {{.elements}}.")
)

type InRule struct {
	elements []interface{}
	err      validation.Error
}

// In 验证字段必须包含在给定的值列表中
func In(values ...interface{}) InRule {
	return InRule{
		elements: values,
		err:      ErrIn,
	}
}

func (r InRule) Validate(key, value interface{}) error {
	value, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(value) {
		return nil
	}
	if len(r.elements) == 0 {
		return nil
	}
	for _, e := range r.elements {
		if reflect.DeepEqual(e, value) {
			return nil
		}
	}
	return r.err.Parse(map[string]interface{}{"field": key, "elements": r.elements})
}

func (r InRule) Error(message string) InRule {
	if r.err == nil {
		r.err = ErrIn
	}
	r.err = r.err.SetMessage(message)
	return r
}
