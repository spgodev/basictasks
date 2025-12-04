package main

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func TestCalculatedSum(t *testing.T) {

	testTable := []struct {
		nums     []float64
		expected float64
	}{
		{
			nums:     []float64{-1.2, -0.7, 1, 2.3, 3.4},
			expected: 4.8,
		},
		{
			nums:     []float64{1, 2.3, 3.4},
			expected: 6.7,
		},
		{
			nums:     []float64{-1.3, 0, 1.3},
			expected: 0,
		},
	}
	for _, testCase := range testTable {
		result := CalculatedSum(testCase.nums)
		// Лучше пользоваться стандартными функциями. Можно использовать require.InDelta
		// а почему вообще рассхождение происходит?
		if !almostEqual(result, testCase.expected, 1e-9) {
			t.Errorf("Incorrect summa. Expected: %f, got: %f", testCase.expected, result)
		}
	}
}
