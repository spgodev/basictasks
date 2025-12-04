package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// что будет если 1 элемент в  slice?
// переписал функцию, теперь все работает корректно
func TestReverse(t *testing.T) {
	testTable := []struct {
		nums     []int64
		expected []int64
	}{
		{
			nums:     []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			expected: []int64{48, 34, 8, 4, 3, 1000, -4, -7, -98},
		},
		{
			nums:     []int64{1},
			expected: []int64{1},
		},
		{
			nums:     []int64{1, 2, 3},
			expected: []int64{3, 2, 1},
		},
	}

	for _, tc := range testTable {
		result := Reverse(tc.nums)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
