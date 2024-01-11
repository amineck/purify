package purify

type (
	// Sanitizable is the interface indicating that the object can be sanitized.
	Sanitizable interface {
		Sanitize() error
	}

	// Rule is the interface indicating that the object can be sanitized.
	Rule interface {
		Apply(value any) (any, error)
	}
)

// Sanitize sanitizes the given value using the specified rules.
func Sanitize(value any, rules ...Rule) (any, error) {
	for _, rule := range rules {
		res, err := rule.Apply(value)
		if err != nil {
			return value, err
		}
		if res == nil {
			return value, nil
		}
		value = res
	}

	return value, nil
}
