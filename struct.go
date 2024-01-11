package purify

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	sanitizableType = reflect.TypeOf((*Sanitizable)(nil)).Elem()
)

type (
	// FieldRules represents a rule set associated with a struct field.
	FieldRules struct {
		fieldPtr any
		rules    []Rule
	}
)

// Field specifies a struct field and the corresponding validation rules.
// The struct field must be specified as a pointer to it.
func Field(fieldPtr any, rules ...Rule) *FieldRules {
	return &FieldRules{
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

// SanitizeStruct sanitizes the given struct by checking the specified struct fields
// against the corresponding sanitization rules.
func SanitizeStruct(structPtr any, fields ...*FieldRules) error {
	value := reflect.ValueOf(structPtr)
	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		// must be a pointer to a struct
		return errors.New("only a pointer to a struct can be sanitized")
	}
	if value.IsNil() {
		// treat a nil struct pointer as valid
		return nil
	}
	value = value.Elem()

	for i, fr := range fields {
		fv := reflect.ValueOf(fr.fieldPtr)
		if fv.Kind() != reflect.Ptr {
			return fmt.Errorf("field %d must be a pointer", i)
		}

		fValue := fv.Elem()
		sanitized, err := Sanitize(fValue.Interface(), fr.rules...)
		if err != nil {
			return err
		}

		fValue.Set(reflect.ValueOf(sanitized))
	}

	if err := findAndSanitizeStructField(structPtr); err != nil {
		return err
	}
	return nil
}

func findAndSanitizeStructField(structPtr any) error {
	value := reflect.ValueOf(structPtr).Elem()

	for i := 0; i < value.NumField(); i++ {
		fv := value.Field(i)
		if fv.Kind() != reflect.Ptr || fv.IsNil() || fv.Elem().Kind() != reflect.Struct {
			continue
		}

		if fv.Type().Implements(sanitizableType) {
			if err := fv.Interface().(Sanitizable).Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
