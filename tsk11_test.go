package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDellel(t *testing.T) {
	testTable := []struct {
		delel    []int64
		idx      int
		expected []int64
	}{
		{
			delel:    []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			idx:      0,
			expected: []int64{-7, -4, 1000, 3, 4, 8, 34, 48},
		},
		{
			delel:    []int64{},
			idx:      0,
			expected: []int64{},
		},
		{
			delel:    []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			idx:      4,
			expected: []int64{-98, -7, -4, 1000, 4, 8, 34, 48},
		},
	}

	for _, tc := range testTable {
		result := dellel(tc.delel, tc.idx)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
