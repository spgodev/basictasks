package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// а если путсой slice передам?
func TestContaines(t *testing.T) {
	testTable := []struct {
		names    []string
		input    string
		expected bool
	}{
		{
			names:    []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"},
			input:    "Sergey",
			expected: true,
		},
		{
			names:    []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"},
			input:    "Anton",
			expected: false,
		},
		{
			names:    []string{"Vova"},
			input:    "Sergey",
			expected: false,
		},
	}

	for _, tc := range testTable {
		result := Containes(tc.names, tc.input)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Input: %v, Search: %s, Expected: %v, Got: %v", tc.names, tc.input, tc.expected, result),
		)
	}
}
