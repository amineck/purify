package purify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhen(t *testing.T) {
	testCases := []struct {
		name      string
		condition bool
		rules     []Rule
		elseRules []Rule
		expected  any
	}{
		{
			name:      "no rules",
			condition: true,
			rules:     []Rule{},
			elseRules: []Rule{},
			expected:  "john doe123!",
		},
		{
			name:      "true condition",
			condition: true,
			rules:     []Rule{ToAlphaNumeric},
			elseRules: []Rule{ToAlpha},
			expected:  "john doe123",
		},
		{
			name:      "false condition",
			condition: false,
			rules:     []Rule{ToAlphaNumeric},
			elseRules: []Rule{ToAlpha},
			expected:  "john doe",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Sanitize("john doe123!", When(tc.condition, tc.rules...).Else(tc.elseRules...))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}
