package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDuplicate(t *testing.T) {
	testTable := []struct {
		slice    []int64
		expected []int64
	}{
		// все упало (
		{
			slice:    nil,
			expected: nil,
		},
		{
			slice:    []int64{},
			expected: []int64{},
		},
		{
			slice:    []int64{1, 2, 10, 15},
			expected: []int64{1, 1, 2, 2, 10, 10, 15, 15},
		},
		{
			slice:    []int64{},
			expected: []int64{},
		},
		{
			slice:    []int64{0, 1},
			expected: []int64{0, 0, 1, 1},
		},
	}

	for _, tc := range testTable {
		result := Duplicate(tc.slice)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
