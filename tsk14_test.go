package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestFoundindex(t *testing.T) {
	testTable := []struct {
		indexf   []int64
		zapros   int64
		expected int
	}{
		{
			indexf:   []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			zapros:   -98,
			expected: 0,
		},
		{
			indexf:   []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48},
			zapros:   34,
			expected: -1,
		},
		{
			indexf:   []int64{},
			zapros:   1,
			expected: -1,
		},
	}

	for _, tc := range testTable {
		result := foundindex(tc.indexf, tc.zapros)

		assert.Equal(
			t,
			tc.expected,
			result,
			fmt.Sprintf("Expected %v, got %v", tc.expected, result))
	}
}
