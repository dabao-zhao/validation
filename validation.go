package validation

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

type (
	// Validatable is the interface indicating the type implementing it supports data validation.
	Validatable interface {
		// Validate validates the data and returns an error if validation fails.
		Validate() error
	}

	// ValidatableWithContext is the interface indicating the type implementing it supports context-aware data validation.
	ValidatableWithContext interface {
		// ValidateWithContext validates the data with the given context and returns an error if validation fails.
		ValidateWithContext(ctx context.Context) error
	}

	// Rule represents a validation rule.
	Rule interface {
		// Validate validates a value and returns a value if validation fails.
		Validate(key, value interface{}) error
	}

	// RuleWithContext represents a context-aware validation rule.
	RuleWithContext interface {
		// ValidateWithContext validates a value and returns a value if validation fails.
		ValidateWithContext(ctx context.Context, value interface{}) error
	}

	// RuleFunc represents a validator function.
	// You may wrap it as a Rule by calling By().
	RuleFunc func(key, value interface{}) error

	// RuleWithContextFunc represents a validator function that is context-aware.
	// You may wrap it as a Rule by calling WithContext().
	RuleWithContextFunc func(ctx context.Context, value interface{}) error

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
