package purify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type struct1 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestSanitizeStruct(t *testing.T) {
	var1 := &struct1{
		Name:  " john doe123 ",
		Email: " John@EXAMPLE.com ",
	}

	testCases := []struct {
		name     string
		input    any
		rules    []*FieldRules
		expected any
	}{
		{
			name:  "struct1 with data",
			input: var1,
			rules: []*FieldRules{
				{fieldPtr: &var1.Name, rules: []Rule{TrimSpace, ToAlpha, ToTitleCase}},
				{fieldPtr: &var1.Email, rules: []Rule{TrimSpace, ToEmail}},
			},
			expected: &struct1{
				Name:  "John Doe",
				Email: "John@example.com",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SanitizeStruct(tc.input, tc.rules...)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
