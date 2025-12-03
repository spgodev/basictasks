package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	testTable := []struct {
		input    []int64
		expected []int64
	}{
		// упал
		{
			input:    []int64{},
			expected: []int64{},
		},
		{
			input:    []int64{-98, -7, -7, 3, 4, 4, 4, 48},
			expected: []int64{-98, -7, 3, 4, 48},
		},
		{
			input:    []int64{4, 4, 4},
			expected: []int64{4},
		},
		{
			input:    []int64{0, 0, 0},
			expected: []int64{0},
		},
	}

	for _, tc := range testTable {
		result := RemoveDuplicate(tc.input)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
