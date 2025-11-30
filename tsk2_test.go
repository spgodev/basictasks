package main

import (
	"testing"
)

func TestCounteven(t *testing.T) {

	testTable := []struct {
		even     []int64
		expected int64
	}{
		{
			even:     []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: 10,
		},
		{
			even:     []int64{0, 30, 30, 4, 8, 1},
			expected: 5,
		},
		{
			even:     []int64{1, 1, 1, 1, 1},
			expected: 0,
		},
	}
	for _, testCase := range testTable {
		result := counteven(testCase.even)
		// можно использовать require.Equal
		if result != testCase.expected {
			t.Errorf("Incorrect count. Expected %d, got %d", testCase.expected, result)
		}
	}
}
