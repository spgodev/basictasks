package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDubl(t *testing.T) {
	testTable := []struct {
		rpt1     []int64
		expected []int64
	}{
		// все упало (
		{
			rpt1:     nil,
			expected: nil,
		},
		{
			rpt1:     []int64{},
			expected: []int64{},
		},
		{
			rpt1:     []int64{1, 2, 10, 15},
			expected: []int64{1, 1, 2, 2, 10, 10, 15, 15},
		},
		{
			rpt1:     []int64{},
			expected: []int64{},
		},
		{
			rpt1:     []int64{0, 1},
			expected: []int64{0, 0, 1, 1},
		},
	}

	for _, tc := range testTable {
		result := dubl(tc.rpt1)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
