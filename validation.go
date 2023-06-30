package validation

import (
	"errors"
	"reflect"
	"sync"
)

type (
	Validatable interface {
		Validate() error
	}

	Rule interface {
		Validate(key, value interface{}) error
	}

	Validation struct {
		fieldTag           string
		stopOnFirstFailure bool
		val                interface{}
		fieldRules         []*FieldRules
		mu                 sync.Mutex
	}
)

func Make(val interface{}, fieldRules ...*FieldRules) *Validation {
	return &Validation{
		fieldTag:           "json",
		stopOnFirstFailure: false,
		val:                val,
		fieldRules:         fieldRules,
		mu:                 sync.Mutex{},
	}
}

func (v *Validation) SetStopOnFirstFailure(stopOnFirstFailure bool) *Validation {
	defer v.mu.Unlock()
	v.mu.Lock()
	v.stopOnFirstFailure = stopOnFirstFailure
	return v
}

func (v *Validation) SetFieldTag(fieldTag string) *Validation {
	defer v.mu.Unlock()
	v.mu.Lock()
	v.fieldTag = fieldTag
	return v
}

func (v *Validation) Validate() error {
	value := reflect.ValueOf(v.val)
	if value.Kind() != reflect.Ptr {
		return errors.New("validated val must be a pointer")
	}
	if value.IsNil() {
		return nil
	}
	switch value.Elem().Kind() {
	case reflect.Struct:
		return v.ValidateStruct()
	case reflect.Map:
		return v.ValidateMap()
	}
	return errors.New("only map,struct can be validated")
}

func (v *Validation) ValidateSlice() error {
	return nil
}
