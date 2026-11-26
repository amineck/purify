package purify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimSpace(t *testing.T) {
	testCases := []struct {
		name        string
		input       any
		expected    any
		expectedErr error
	}{
		{
			name:     "string with spaces",
			input:    "  john doe123 ",
			expected: "john doe123",
		},
		{
			name:     "string with tabs",
			input:    "\tjohn doe123\t",
			expected: "john doe123",
		},
		{
			name:     "string with newlines",
			input:    "\njohn doe123\n",
			expected: "john doe123",
		},
		{
			name:     "string with carriage returns",
			input:    "\rjohn doe123\r",
			expected: "john doe123",
		},
		{
			name:     "string with mixed whitespace",
			input:    " \t\n\rjohn doe123 \t\n\r",
			expected: "john doe123",
		},
		{
			name:        "non string",
			input:       123,
			expected:    nil,
			expectedErr: errors.New("expected string, got int"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := TrimSpace.Apply(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToAlpha(t *testing.T) {
	testCases := []struct {
		name        string
		input       any
		expected    any
		expectedErr error
	}{
		{
			name:     "with numbers",
			input:    "john doe123",
			expected: "john doe",
		},
		{
			name:     "with punctuation",
			input:    "john doe%*(?!+=",
			expected: "john doe",
		},
		{
			name:        "non string",
			input:       123,
			expected:    nil,
			expectedErr: errors.New("must be either a string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToAlpha.Apply(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToAlphaNumeric(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		expected  any
		expectErr bool
	}{
		{
			name:     "keeps letters and digits",
			input:    "johndoe123",
			expected: "johndoe123",
		},
		{
			name:     "removes spaces",
			input:    "john doe 123",
			expected: "johndoe123",
		},
		{
			name:     "removes leading and trailing spaces",
			input:    "  john doe123  ",
			expected: "johndoe123",
		},
		{
			name:     "removes tabs, newlines and carriage returns",
			input:    "john\tdoe\n123\r",
			expected: "johndoe123",
		},
		{
			name:     "removes punctuation",
			input:    "john%*(?!+=doe123",
			expected: "johndoe123",
		},
		{
			name:     "removes spaces and punctuation together",
			input:    " john doe%*(?!+= 123 ",
			expected: "johndoe123",
		},
		{
			name:     "keeps unicode letters",
			input:    "jöhn dóe 123",
			expected: "jöhndóe123",
		},
		{
			name:     "only spaces and punctuation becomes empty",
			input:    "  !? ",
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "byte slice",
			input:    []byte("john doe 123!"),
			expected: "johndoe123",
		},
		{
			name:      "non string",
			input:     123,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := ToAlphaNumeric.Apply(tt.input)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestToNumeric(t *testing.T) {
	testCases := []struct {
		name        string
		input       any
		expected    any
		expectedErr error
	}{
		{
			name:     "with alpha",
			input:    "john123",
			expected: "123",
		},
		{
			name:     "with ponctuation",
			input:    "john123 %*(?!+=",
			expected: "123",
		},
		{
			name:        "non string",
			input:       123,
			expected:    nil,
			expectedErr: errors.New("must be either a string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToNumeric.Apply(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}
