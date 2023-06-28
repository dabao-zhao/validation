package validation

import (
	"fmt"
	"reflect"
	"strings"
)

func (v *Validation) ValidateStruct() error {
	value := reflect.ValueOf(v.val).Elem()
	errs := Errors{}
	for i, fr := range v.fieldRules {
		fv := reflect.ValueOf(fr.field)
		if fv.Kind() != reflect.Ptr {
			return fmt.Errorf("field #%v must be specified as a pointer", i)
		}
		ft := findStructField(value, fv)
		if ft == nil {
			return fmt.Errorf("field #%v cannot be found in the struct", i)
		}

		for _, rule := range fr.rules {
			err := rule.Validate(getFullFieldName(value, fv, v.fieldTag), fv.Elem().Interface())
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

func findStructField(structValue reflect.Value, fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()
	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if sf.Type == fieldValue.Elem().Type() {
				return &sf
			}
		}
		fi := structValue.Field(i)
		if sf.Type.Kind() == reflect.Ptr {
			fi = fi.Elem()
		}
		if fi.Kind() == reflect.Struct {
			if f := findStructField(fi, fieldValue); f != nil {
				return f
			}
		}
	}
	return nil
}

func getErrorFieldName(f *reflect.StructField, tagName string) string {
	if tag := f.Tag.Get(tagName); tag != "" && tag != "-" {
		if cps := strings.SplitN(tag, ",", 2); cps[0] != "" {
			return cps[0]
		}
	}
	return f.Name
}

func getFullFieldName(structValue reflect.Value, fieldValue reflect.Value, tagName string) string {
	ptr := fieldValue.Pointer()
	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if sf.Type == fieldValue.Elem().Type() {
				return getErrorFieldName(&sf, tagName)
			}
		}
		fi := structValue.Field(i)
		if sf.Type.Kind() == reflect.Ptr {
			fi = fi.Elem()
		}
		tag := getErrorFieldName(&sf, tagName)
		if fi.Kind() == reflect.Struct {
			if f := findStructField(fi, fieldValue); f != nil {
				return tag + "." + getErrorFieldName(f, tagName)
			}
		}
	}
	return ""
}
