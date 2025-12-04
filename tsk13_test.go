package main

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestSeparationParity(t *testing.T) {
	testTable := []struct {
		slice        []int64
		expectedEven []int64
		expectedOdd  []int64
	}{
		{
			slice:        []int64{1, 2, 3, 4, 6, 7, 8, 10},
			expectedEven: []int64{2, 4, 6, 8, 10},
			expectedOdd:  []int64{1, 3, 7},
		},
		{
			slice:        []int64{1, 3, 11},
			expectedEven: []int64{},
			expectedOdd:  []int64{1, 3, 11},
		},
		{
			slice:        []int64{},
			expectedEven: []int64{},
			expectedOdd:  []int64{},
		},
	}

	for _, tc := range testTable {
		even, odd := SeparationParity(tc.slice)

		assert.Equal(
			t,
			tc.expectedEven,
			even,
			fmt.Sprintf("Expected %v, got %v", tc.expectedEven, even))
		assert.Equal(
			t,
			tc.expectedOdd,
			odd,
			fmt.Sprintf("Expected %v, got %v", tc.expectedOdd, odd))
	}
}
