package main

func calculateSum(nums []float64) float64 {

	var (
		sum float64
	)

	for _, v := range nums {
		sum += v
	}
	return sum
}
