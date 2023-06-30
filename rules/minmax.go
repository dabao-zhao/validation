package rules

import (
	"reflect"

	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrNumericMax = validation.NewError("the {{.field}} field must be less than or equal to {{.threshold}}.")
	ErrNumericMin = validation.NewError("the {{.field}} field must be greeter than or equal to {{.threshold}}.")

	ErrStringMax = validation.NewError("the {{.field}} field must be less than or equal to {{.threshold}} characters.")
	ErrStringMin = validation.NewError("the {{.field}} field must be greeter than or equal to {{.threshold}} characters.")

	ErrItemsMax = validation.NewError("the {{.field}} field must be less than or equal to {{.threshold}} items.")
	ErrItemsMin = validation.NewError("the {{.field}} field must be greeter than or equal to {{.threshold}} items.")
)

const (
	maxOp = iota
	minOp
)

type ThresholdRule struct {
	threshold interface{}
	operator  int
	err       validation.Error
}

func Min(min interface{}) ThresholdRule {
	return ThresholdRule{
		threshold: min,
		operator:  minOp,
		err:       ErrNumericMin,
	}
}

func Max(max interface{}) ThresholdRule {
	return ThresholdRule{
		threshold: max,
		operator:  maxOp,
		err:       ErrNumericMax,
	}
}

func (r ThresholdRule) Validate(key, value interface{}) error {
	value, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(value) {
		return nil
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := util.ToInt(value)
		if err != nil {
			return err
		}
		threshold, err := util.ToInt(r.threshold)
		if err != nil {
			return err
		}
		if r.compareInt(threshold, v) {
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := util.ToUint(value)
		if err != nil {
			return err
		}
		threshold, err := util.ToUint(r.threshold)
		if err != nil {
			return err
		}
		if r.compareUint(threshold, v) {
			return nil
		}
	case reflect.Float32, reflect.Float64:
		v, err := util.ToFloat(value)
		if err != nil {
			return err
		}
		threshold, err := util.ToFloat(r.threshold)
		if err != nil {
			return err
		}
		if r.compareFloat(threshold, v) {
			return nil
		}
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		if r.operator == maxOp {
			if v.Kind() == reflect.String {
				r.err = ErrStringMax
			} else {
				r.err = ErrItemsMax
			}
		} else {
			if v.Kind() == reflect.String {
				r.err = ErrStringMin
			} else {
				r.err = ErrItemsMin
			}
		}
		threshold, err := util.ToInt(r.threshold)
		if err != nil {
			return err
		}
		if r.compareStringSliceMapArray(threshold, int64(v.Len())) {
			return nil
		}
	}
	return r.err.Parse(map[string]interface{}{"field": key, "threshold": r.threshold})
}

func (r ThresholdRule) compareInt(threshold, value int64) bool {
	switch r.operator {
	case minOp:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareUint(threshold, value uint64) bool {
	switch r.operator {
	case minOp:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareFloat(threshold, value float64) bool {
	switch r.operator {
	case minOp:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareStringSliceMapArray(threshold, l int64) bool {
	switch r.operator {
	case minOp:
		return l >= threshold
	default:
		return l <= threshold
	}
}

func (r ThresholdRule) Error(message string) ThresholdRule {
	r.err = r.err.SetMessage(message)
	return r
}
