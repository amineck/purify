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
			expectedErr: errors.New("purify: TrimSpace: expected string, got int"),
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
			name:     "string",
			input:    "john doe123",
			expected: "john doe",
		},
		{
			name:        "non string",
			input:       123,
			expected:    nil,
			expectedErr: errors.New(""),
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
