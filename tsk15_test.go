package main

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrades(t *testing.T) {
	testTable := []struct {
		tableGrades []int
		expected    string
	}{
		{
			tableGrades: []int{1, 2, 3, 4, 5},
			expected:    "Very Bad, Bad, Average, Good, Excellent",
		},
		{
			tableGrades: []int{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			expected:    "Unknown grade, Unknown grade, Unknown grade, Unknown grade, Average, Good, Good, Good, Unknown grade",
		},
		{
			tableGrades: []int{},
			expected:    "",
		},
	}

	for _, tc := range testTable {
		result := Grades(tc.tableGrades)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result))
	}
}
