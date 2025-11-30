package main

import "fmt"

// Название функции должно быть глаголом, тк она ДЕЛАЕТ какую-то работу
// Например calculateArithmeticMean
func arithm(num []int64) string {
	var sum int64
	for _, v := range num {
		sum += v
	}

	if len(num) == 0 {
		return "empty slice"
	}

	avg := sum / int64(len(num))

	if avg > 100 {
		return "Big one!"
	}

	return fmt.Sprintf("Среднее арифмитическое = %d", avg)
}
