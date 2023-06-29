package rules

import (
	"reflect"

	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
)

var (
	ErrNumericGreaterEqualThanRequired = validation.NewError("the {{.field}} field must not be greater than or equal to {{.threshold}}.")
	ErrNumericLessEqualThanRequired    = validation.NewError("the {{.field}} field must not be less than or equal to {{.threshold}}.")

	ErrStringGreaterEqualThanRequired = validation.NewError("the {{.field}} field must not be greater than or equal to {{.threshold}} characters.")
	ErrStringLessEqualThanRequired    = validation.NewError("the {{.field}} field must not be less than or equal to {{.threshold}} characters.")

	ErrItemsGreaterEqualThanRequired = validation.NewError("the {{.field}} field must not be greater than or equal to {{.threshold}} items.")
	ErrItemsLessEqualThanRequired    = validation.NewError("the {{.field}} field must not be less than or equal to {{.threshold}} items.")
)

const (
	greaterEqualThan = iota
	lessEqualThan
)

type ThresholdRule struct {
	threshold interface{}
	operator  int
	err       validation.Error
}

func Min(min interface{}) ThresholdRule {
	return ThresholdRule{
		threshold: min,
		operator:  greaterEqualThan,
		err:       ErrNumericLessEqualThanRequired,
	}
}

func Max(max interface{}) ThresholdRule {
	return ThresholdRule{
		threshold: max,
		operator:  lessEqualThan,
		err:       ErrNumericGreaterEqualThanRequired,
	}
}

func (r ThresholdRule) Validate(key, value interface{}) error {
	value, isNil := util.Indirect(value)
	if isNil || util.IsEmpty(value) {
		return nil
	}
	v := reflect.ValueOf(value)
	rv := reflect.ValueOf(r.threshold)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := util.ToInt(value)
		if err != nil {
			return err
		}
		if r.compareInt(rv.Int(), v) {
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := util.ToUint(value)
		if err != nil {
			return err
		}
		if r.compareUint(rv.Uint(), v) {
			return nil
		}
	case reflect.Float32, reflect.Float64:
		v, err := util.ToFloat(value)
		if err != nil {
			return err
		}
		if r.compareFloat(rv.Float(), v) {
			return nil
		}
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		if r.operator == greaterEqualThan {
			if v.Kind() == reflect.String {
				r.err = ErrStringGreaterEqualThanRequired
			} else {
				r.err = ErrItemsGreaterEqualThanRequired
			}
		} else {
			if v.Kind() == reflect.String {
				r.err = ErrStringLessEqualThanRequired
			} else {
				r.err = ErrItemsLessEqualThanRequired
			}
		}
		if r.compareStringSliceMapArray(rv.Int(), int64(v.Len())) {
			return nil
		}
	}
	return r.err.Parse(map[string]interface{}{"field": key, "threshold": r.threshold})
}

func (r ThresholdRule) compareInt(threshold, value int64) bool {
	switch r.operator {
	case greaterEqualThan:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareUint(threshold, value uint64) bool {
	switch r.operator {
	case greaterEqualThan:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareFloat(threshold, value float64) bool {
	switch r.operator {
	case greaterEqualThan:
		return value >= threshold
	default:
		return value <= threshold
	}
}

func (r ThresholdRule) compareStringSliceMapArray(threshold, l int64) bool {
	switch r.operator {
	case greaterEqualThan:
		return l >= threshold
	default:
		return l <= threshold
	}
}
