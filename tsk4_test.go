package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тесты не проходят.
func TestArithm(t *testing.T) {
	testTable := []struct {
		num      []int64
		expected string
	}{
		{
			num:      []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48},
			expected: "Big one!",
		},
		{
			num:      []int64{0, 10, 20, 30},
			expected: "Среднее арифмитическое = 15",
		},
		{
			num:      []int64{1, 1},
			expected: "Среднее арифмитическое = 2",
		},
	}

	for _, tc := range testTable {
		result := arithm(tc.num)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %q, got %q", tc.expected, result),
		)
	}
}
