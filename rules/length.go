package rules

import (
	"errors"
	"github.com/dabao-zhao/validation"
	"unicode/utf8"
)

var (
	ErrLength = validation.NewError("the {{.field}} field length must be between {{.min}} and {{.max}}.")
)

type LengthRule struct {
	min, max int
	rune     bool

	err validation.Error
}

// Length 长度验证
// map、slice、array、string 对应为 len
func Length(min, max int) LengthRule {
	return LengthRule{
		min: min,
		max: max,
		err: ErrLength,
	}
}

// RuneLength 长度验证
// 只验证 string 对应为 utf8.RuneCountInString
func RuneLength(min, max int) LengthRule {
	r := Length(min, max)
	r.rune = true

	return r
}

func (r LengthRule) Validate(key, value interface{}) error {
	if r.max <= r.min {
		return errors.New("the length max must be more than min")
	}

	v, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(v) {
		return nil
	}

	var (
		l   int
		err error
	)

	if r.rune {
		s, ok := v.(string)
		if !ok {
			return errors.New("the RuneLength rule requires a type of string")
		}
		l = utf8.RuneCountInString(s)
	} else {
		if l, err = validation.LengthOfValue(v); err != nil {
			return err
		}
	}

	if r.min > 0 && l < r.min || r.max > 0 && l > r.max || r.min == 0 && r.max == 0 && l > 0 {
		return r.err.Parse(map[string]interface{}{"field": key, "max": r.max, "min": r.min})
	}

	return nil
}

func (r LengthRule) Error(message string) LengthRule {
	if r.err == nil {
		r.err = ErrLength
	}
	r.err = r.err.SetMessage(message)
	return r
}
