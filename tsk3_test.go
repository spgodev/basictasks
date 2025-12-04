package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOdd(t *testing.T) {
	testTable := []struct {
		nums     []int64
		expected int
	}{
		{
			nums:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: 10,
		},
		{
			nums:     []int64{0, 30, 30, 4, 8, 1},
			expected: 1,
		},
		{
			nums:     []int64{2, 2, 2, 2, 2},
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := CountOdd(testCase.nums)

		assert.Equal(
			t,
			testCase.expected,
			result,
			fmt.Sprintf("Incorrect result. Expected %d, got %d", testCase.expected, result),
		)
	}
}
