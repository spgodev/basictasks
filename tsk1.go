package main

func calculateSum(nums []float64) float64 {

	sum := 0.0

	for _, v := range nums {
		sum += v
	}
	return sum
}
