package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigone(t *testing.T) {
	testTable := []struct {
		num      []int64
		expected int64
	}{
		{
			num:      []int64{-98, -7, -4, 47, 3, 4, 8, 34, 48},
			expected: 48,
		},
		{
			num:      []int64{0, 10, 20, 30},
			expected: 30,
		},
		{
			num:      []int64{1, 1},
			expected: 1,
		},
	}

	for _, tc := range testTable {
		result := bigone(tc.num)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
