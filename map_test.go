package purify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		rules    []*KeyRules
		expected any
	}{
		{
			name: "map[string]string",
			input: map[string]string{
				"name":    " john doe123 ",
				"email":   " John@EXAMPLE.com ",
				"country": "US",
			},
			rules: []*KeyRules{
				{key: "name", rules: []Rule{TrimSpace, ToAlpha, ToTitleCase}},
				{key: "email", rules: []Rule{TrimSpace, ToEmail}},
			},
			expected: map[string]string{
				"name":    "John Doe",
				"email":   "John@example.com",
				"country": "US",
			},
		},
		{
			name: "map[string]interface{}",
			input: map[string]interface{}{
				"name":  " john doe123 ",
				"email": " John@EXAMPLE.com ",
			},
			rules: []*KeyRules{
				{key: "name", rules: []Rule{TrimSpace, ToAlpha, ToTitleCase}},
				{key: "email", rules: []Rule{TrimSpace, ToEmail, ToSHA256}},
			},
			expected: map[string]interface{}{
				"name":  "John Doe",
				"email": "20740450eae791b7928c1869d3a0c964e8685544eb0d70c386d6ba825270b12e",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SanitizeMap(tc.input, tc.rules...)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
