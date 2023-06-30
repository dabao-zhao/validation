package rules

import (
	"github.com/dabao-zhao/validation"
	"github.com/dabao-zhao/validation/util"
	"regexp"
)

var (
	ErrRegex = validation.NewError("the {{.field}} field format is invalid.")
)

type RegexRule struct {
	re  *regexp.Regexp
	err validation.Error
}

// Regex 验证的字段必须匹配给定的正则表达式
func Regex(re *regexp.Regexp) RegexRule {
	return RegexRule{
		re:  re,
		err: ErrRegex,
	}
}

func (r RegexRule) Validate(key, value interface{}) error {
	v, isNil := util.Indirect(value)
	if isNil {
		return nil
	}

	isString, str, isBytes, bs := util.StringOrBytes(v)
	if isString && (str == "" || r.re.MatchString(str)) {
		return nil
	} else if isBytes && (len(bs) == 0 || r.re.Match(bs)) {
		return nil
	}

	return r.err.Parse(map[string]interface{}{"field": key})
}

func (r RegexRule) Error(message string) RegexRule {
	if r.err == nil {
		r.err = ErrRegex
	}
	r.err = r.err.SetMessage(message)
	return r
}
