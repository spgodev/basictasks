package main

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func TestSumma(t *testing.T) {

	testTable := []struct {
		sums     []float64
		expected float64
	}{
		{
			sums:     []float64{-1.2, -0.7, 1, 2.3, 3.4},
			expected: 4.8,
		},
		{
			sums:     []float64{1, 2.3, 3.4},
			expected: 6.7,
		},
		{
			sums:     []float64{-1.3, 0, 1.3},
			expected: 0,
		},
	}
	for _, testCase := range testTable {
		result := summa(testCase.sums)
		if !almostEqual(result, testCase.expected, 1e-9) {
			t.Errorf("Incorrect summa. Expected: %f, got: %f", testCase.expected, result)
		}
	}
}
