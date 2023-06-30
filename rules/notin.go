package rules

import (
	"reflect"

	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrNotIn = validation.NewError("the {{.field}} field must not exist in {{.elements}}.")
)

type NotInRule struct {
	elements []interface{}
	err      validation.Error
}

// NotIn 验证字段必须未包含在给定的值列表中
func NotIn(values ...interface{}) NotInRule {
	return NotInRule{
		elements: values,
		err:      ErrNotIn,
	}
}

func (r NotInRule) Validate(key, value interface{}) error {
	value, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(value) {
		return nil
	}
	if len(r.elements) == 0 {
		return nil
	}
	for _, e := range r.elements {
		if reflect.DeepEqual(e, value) {
			return r.err.Parse(map[string]interface{}{"field": key, "elements": r.elements})
		}
	}
	return nil
}

func (r NotInRule) Error(message string) NotInRule {
	r.err = r.err.SetMessage(message)
	return r
}
