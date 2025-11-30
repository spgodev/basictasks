package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// а если путсой slice передам?
func TestZapros(t *testing.T) {
	testTable := []struct {
		names    []string
		imput    string
		expected bool
	}{
		{
			names:    []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"},
			imput:    "Sergey",
			expected: true,
		},
		{
			names:    []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"},
			imput:    "Anton",
			expected: false,
		},
		{
			names:    []string{"Vova"},
			imput:    "Sergey",
			expected: false,
		},
	}

	for _, tc := range testTable {
		result := zapros(tc.names, tc.imput)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Input: %v, Search: %s, Expected: %v, Got: %v", tc.names, tc.imput, tc.expected, result),
		)
	}
}
