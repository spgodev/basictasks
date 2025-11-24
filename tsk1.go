package main

func summa(sums []float64) float64 {

	var (
		plus float64
	)

	for _, v := range sums {
		plus += v
	}
	return plus
}
