package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	testTable := []struct {
		nums     []int64
		idx      int
		expected []int64
	}{
		{
			nums:     []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			idx:      0,
			expected: []int64{-7, -4, 1000, 3, 4, 8, 34, 48},
		},
		{
			nums:     []int64{},
			idx:      0,
			expected: []int64{},
		},
		{
			nums:     []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			idx:      4,
			expected: []int64{-98, -7, -4, 1000, 4, 8, 34, 48},
		},
	}

	for _, tc := range testTable {
		result := RemoveElement(tc.nums, tc.idx)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
