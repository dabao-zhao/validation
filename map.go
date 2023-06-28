package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (v *Validation) ValidateMap() error {
	value := reflect.ValueOf(v.val)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.IsNil() {
		return nil
	}

	errs := Errors{}

	for _, fr := range v.fieldRules {
		field, ok := fr.field.(string)
		if !ok {
			return errors.New("the field must be string")
		}

		vv, err := getMapValue(value, field)
		if err != nil {
			return err
		}
		for _, rule := range fr.rules {
			err := rule.Validate(field, vv.Interface())
			if err == nil {
				continue
			}
			if v.stopOnFirstFailure {
				return err
			}
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func getMapValue(m reflect.Value, field string) (reflect.Value, error) {
	sp := strings.Split(field, ".")
	kt := m.Type().Key()
	if len(sp) > 1 {
		for i, k := range sp {
			kv := reflect.ValueOf(k)
			if !kt.AssignableTo(kv.Type()) {
				return reflect.Value{}, errors.New("key not the correct type")
			}
			vv := m.MapIndex(kv)

			if !vv.IsValid() {
				return reflect.Value{}, errors.New("required key is missing")
			}
			if i == (len(sp) - 1) {
				return vv, nil
			}

			if reflect.TypeOf(vv.Interface()).Kind() != reflect.Map {
				return reflect.Value{}, fmt.Errorf("the field %s not a map", k)
			}

			m = reflect.ValueOf(vv.Interface())
			if m.Kind() == reflect.Ptr {
				m = m.Elem()
			}
			if m.IsNil() {
				return reflect.Value{}, nil
			}
		}
	} else {
		kv := reflect.ValueOf(field)
		if !kt.AssignableTo(kv.Type()) {
			return reflect.Value{}, errors.New("key not the correct type")
		}
		vv := m.MapIndex(kv)
		if !vv.IsValid() {
			return reflect.Value{}, errors.New("required key is missing")
		}
		return vv, nil
	}
	return reflect.Value{}, nil
}
