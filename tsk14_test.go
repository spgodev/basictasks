package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestFoundIndex(t *testing.T) {
	testTable := []struct {
		nums     []int64
		input    int64
		expected int
	}{
		{
			nums:     []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			input:    -98,
			expected: 0,
		},
		{
			nums:     []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			input:    34,
			expected: -1,
		},
		{
			nums:     []int64{},
			input:    1,
			expected: -1,
		},
	}

	for _, tc := range testTable {
		result := FoundIndex(tc.nums, tc.input)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result))
	}
}
