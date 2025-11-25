package main

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrades(t *testing.T) {
	testTable := []struct {
		tablegrades []int
		expected    string
	}{
		{
			tablegrades: []int{1, 2, 3, 4, 5},
			expected:    "Very Bad, Bad, Average, Good, Excellent",
		},
		{
			tablegrades: []int{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			expected:    "Unknown grade, Unknown grade, Unknown grade, Unknown grade, Average, Good, Good, Good, Unknown grade",
		},
		{
			tablegrades: []int{},
			expected:    "",
		},
	}

	for _, tc := range testTable {
		result := grades(tc.tablegrades)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result))
	}
}
