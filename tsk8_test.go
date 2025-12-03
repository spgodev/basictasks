package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOdd(t *testing.T) {
	testTable := []struct {
		nums     []int64
		expected []int64
	}{
		{
			nums:     []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			expected: []int64{-7, 3},
		},
		{
			nums:     []int64{-98, -4, 1000, 4, 8, 34, 48},
			expected: []int64{},
		},
		{
			nums:     []int64{0, 1, 2, 3},
			expected: []int64{1, 3},
		},
	}

	for _, tc := range testTable {
		result := RemoveOdd(tc.nums)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
