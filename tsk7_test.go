package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// что будет если 1 элемент в  slice?
func TestSwap(t *testing.T) {
	testTable := []struct {
		rot      []int64
		expected []int64
	}{
		{
			rot:      []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			expected: []int64{48, 34, 8, 4, 3, 1000, -4, -7, -98},
		},
		{
			rot:      []int64{1, 2},
			expected: []int64{2, 1},
		},
		{
			rot:      []int64{1, 2, 3},
			expected: []int64{3, 2, 1},
		},
	}

	for _, tc := range testTable {
		result := swap(tc.rot)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
