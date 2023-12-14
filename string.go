package purify

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	bytesType  = reflect.TypeOf([]byte(nil))
	valuerType = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
)

var (
	reAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9]`)
	reNumeric      = regexp.MustCompile(`[^0-9]`)
	reAlpha        = regexp.MustCompile(`[^\p{L}\s]`)
)

type stringSanitizer func(value string) string

// StringRule handles sanitizing strings using a sanitizer function.
type StringRule struct {
	sanitizer stringSanitizer
}

// Apply applies the string sanitizer to the value.
func (r *StringRule) Apply(ptr any) (any, error) {
	value, isNil := validation.Indirect(ptr)
	if isNil {
		return nil, nil
	}

	str, err := EnsureString(value)
	if err != nil {
		return nil, err
	}

	return r.sanitizer(str), nil
}

// NewStringRule creates a new string rule.
func NewStringRule(sanitizer stringSanitizer) *StringRule {
	return &StringRule{
		sanitizer: sanitizer,
	}
}

// EnsureString ensures the given value is a string.
// If the value is a byte slice, it will be typecast into a string.
// An error is returned otherwise.
func EnsureString(value interface{}) (string, error) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.String {
		return v.String(), nil
	}
	if v.Type() == bytesType {
		return string(v.Interface().([]byte)), nil
	}
	return "", errors.New("must be either a string or byte slice")
}

func trimLeftSpace(value string) string {
	return strings.TrimLeft(value, " ")
}

func trimRightSpace(value string) string {
	return strings.TrimRight(value, " ")
}

func trim(value string) func(string) string {
	return func(s string) string {
		return strings.Trim(s, value)
	}
}

func toEmail(value string) string {
	splits := strings.Split(value, "@")
	if len(splits) != 2 {
		return value
	}
	return fmt.Sprintf("%s@%s", splits[0], strings.ToLower(splits[1]))
}

func toAlphaNumeric(value string) string {
	return reAlphaNumeric.ReplaceAllLiteralString(value, "")
}

func toAlpha(value string) string {
	return reAlpha.ReplaceAllLiteralString(value, "")
}

func toNumeric(value string) string {
	return reNumeric.ReplaceAllLiteralString(value, "")
}
