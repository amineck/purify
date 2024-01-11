package purify

import (
	"errors"
	"reflect"
)

type (
	// KeyRules represents a rule set associated with a map key.
	KeyRules struct {
		key      any
		optional bool
		rules    []Rule
	}
)

// Key specifies a map key and the corresponding validation rules.
func Key(key any, rules ...Rule) *KeyRules {
	return &KeyRules{
		key:   key,
		rules: rules,
	}
}

// SanitizeMap sanitizes the given map by checking the specified map keys
// against the corresponding sanitization rules.
func SanitizeMap(value any, fields ...*KeyRules) error {
	rv := reflect.ValueOf(value)
	if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && rv.IsNil() {
		return nil
	}
	if rv.Kind() != reflect.Map {
		// must be a map
		return errors.New("only a map can be sanitized with SanitizeMap")
	}

	for _, fr := range fields {
		if fr.key == nil {
			continue
		}
		// TODO: handle nested keys
		key := reflect.ValueOf(fr.key)
		if kValue := rv.MapIndex(key); kValue.IsValid() {
			sanitized, err := Sanitize(kValue.Interface(), fr.rules...)
			if err != nil {
				return err
			}
			rv.SetMapIndex(key, reflect.ValueOf(sanitized))
		}
	}
	return nil
}
