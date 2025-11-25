package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDelldbl(t *testing.T) {
	testTable := []struct {
		deldbl   []int64
		expected []int64
	}{
		{
			deldbl:   []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			expected: []int64{-98, -7, 1000, 3, 4, 48},
		},
		{
			deldbl:   []int64{4, 4, 4},
			expected: []int64{4},
		},
		{
			deldbl:   []int64{0, 0, 0},
			expected: []int64{0},
		},
	}

	for _, tc := range testTable {
		result := delldbl(tc.deldbl)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result),
		)
	}
}
