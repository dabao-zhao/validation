package validation

import (
	"bytes"
	"encoding/json"
	"errors"
	"text/template"
)

type (
	Errors []error

	Error interface {
		Error() string
		SetMessage(string) Error
		Parse(interface{}) error
	}

	InternalError struct {
		message string
	}
)

func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}
	b, _ := json.Marshal(es)
	return string(b)
}

func (es Errors) MarshalJSON() ([]byte, error) {
	errs := make([]interface{}, len(es))
	for i, err := range es {
		errs[i] = err.Error()
	}
	return json.Marshal(errs)
}

func NewError(message string) Error {
	return InternalError{
		message: message,
	}
}

func (e InternalError) SetMessage(message string) Error {
	e.message = message
	return e
}

func (e InternalError) Error() string {
	return e.message
}

func (e InternalError) Parse(field interface{}) error {
	res := bytes.Buffer{}
	_ = template.Must(template.New("err").Parse(e.message)).Execute(&res, field)

	return errors.New(res.String())
}
