package main

import (
	"fmt"
)

func arithm() {
	num := []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48}
	var arithm int64 = 0
	var sum int64 = 0
	var count int64 = 0

	for _, v := range num {
		sum += v
		count++
	}
	arithm = sum / count
	if arithm > 100 {
		fmt.Println("Big one!")
	} else {
		fmt.Println("Среднее арифмитическое=", arithm)
	}
}
