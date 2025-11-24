package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestSringsfortest(t *testing.T) {
	testTable := []struct {
		names    []string
		expected string
	}{
		{
			names:    []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"},
			expected: "Alexander",
		},
		{
			names:    []string{},
			expected: "",
		},
		{
			names:    []string{"Privet", "Poka"},
			expected: "Privet",
		},
	}

	for _, tc := range testTable {
		result := stringsfortest(tc.names)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
