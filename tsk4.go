package main

import "fmt"

// Название функции должно быть глаголом, тк она ДЕЛАЕТ какую-то работу
// Например calculateArithmeticMean
func CalculateArithmeticMean(num []int64) string {
	sum := int64(0)
	for _, v := range num {
		sum += v
	}

	avg := sum / int64(len(num))

	if avg > 100 {
		return "Big one!"
	}

	return fmt.Sprintf("Среднее арифмитическое = %d", avg)
}
