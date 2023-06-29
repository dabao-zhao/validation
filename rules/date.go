package rules

import (
	"errors"
	"reflect"
	"time"

	"github.com/dabao-zhao/validation"
)

var (
	ErrDate       = validation.NewError("the {{.field}} field must be a valid date.")
	ErrDateBefore = validation.NewError("the {{.field}} field must be a date before {{.date}}.")
	ErrDateAfter  = validation.NewError("the {{.field}} field must be a date after {{.date}}.")
	ErrDateEqual  = validation.NewError("the {{.field}} field must be a date equal to {{.date}}.")
)

const (
	date = iota
	dateBefore
	dateAfter
	dateEqual

	layout = "2006-01-02 15:04:05"
)

type DateRule struct {
	time   time.Time
	err    validation.Error
	typ    int
	layout string
}

// Date 验证字段必须是 time.parse 可解析的有效日期
var Date = DateRule{
	err:    ErrDate,
	typ:    date,
	layout: layout,
}

// DateBefore 验证字段必须是给定日期之前的日期
// 给定日期支持 time 和 string
// 若为 string 会根据 layout 格式进行转换
func DateBefore(t time.Time) DateRule {
	return DateRule{
		time:   t,
		err:    ErrDateBefore,
		typ:    dateBefore,
		layout: layout,
	}
}

// DateAfter 验证字段必须是给定日期之后的日期
// 给定日期支持 time 和 string
// 若为 string 会根据 layout 格式进行转换
func DateAfter(t time.Time) DateRule {
	return DateRule{
		time:   t,
		err:    ErrDateAfter,
		typ:    dateAfter,
		layout: layout,
	}
}

// DateEqual 验证字段必须等于给定日期
// 给定日期支持 time 和 string
// 若为 string 会根据 layout 格式进行转换
func DateEqual(t time.Time) DateRule {
	return DateRule{
		time:   t,
		err:    ErrDateEqual,
		typ:    dateEqual,
		layout: layout,
	}
}

func (r DateRule) Validate(key, value interface{}) error {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}
	var (
		t   time.Time
		err error
		ok  bool
	)
	typeOf := reflect.TypeOf(value)
	if typeOf.Kind() != reflect.String && typeOf.String() != "time.Time" {
		return errors.New("only string,time.Time can be validated")
	} else if typeOf.Kind() == reflect.String {
		t, err = time.Parse(r.layout, value.(string))
	} else if typeOf.String() == "time.Time" {
		t, ok = value.(time.Time)
		if !ok {
			err = errors.New("not time.Time")
		}
	}
	switch r.typ {
	case date:
		if err != nil {
			return r.err.Parse(map[string]interface{}{
				"field": key,
			})
		}
	case dateBefore:
		if !t.Before(r.time) {
			return r.err.Parse(map[string]interface{}{
				"field": key,
				"date":  r.time.Format(r.layout),
			})
		}
	case dateAfter:
		if !t.After(r.time) {
			return r.err.Parse(map[string]interface{}{
				"field": key,
				"date":  r.time.Format(r.layout),
			})
		}
	case dateEqual:
		if !t.Equal(r.time) {
			return r.err.Parse(map[string]interface{}{
				"field": key,
				"date":  r.time.Format(r.layout),
			})
		}
	}
	return nil
}

func (r DateRule) SetLayout(layout string) DateRule {
	r.layout = layout
	return r
}
