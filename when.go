package purify

// When returns a validation rule that executes the given list of rules when the condition is true.
func When(condition bool, rules ...Rule) WhenRule {
	return WhenRule{
		condition: condition,
		rules:     rules,
		elseRules: []Rule{},
	}
}

// WhenRule is a validation rule that executes the given list of rules when the condition is true.
type WhenRule struct {
	condition bool
	rules     []Rule
	elseRules []Rule
}

// Apply checks if the condition is true and if so, it sanitizes the value using the specified rules.
func (r WhenRule) Apply(value interface{}) (any, error) {
	if r.condition {
		return Sanitize(value, r.rules...)
	}
	return Sanitize(value, r.elseRules...)
}

// Else returns a validation rule that executes the given list of rules when the condition is false.
func (r WhenRule) Else(rules ...Rule) WhenRule {
	r.elseRules = rules
	return r
}
