package main

import (
	"fmt"
)

func countodd() {
	odd := []int64{-7, -4, 1, 3, 4, 8}
	var count int64 = 0
	{
		for _, v := range odd {
			if v%2 == 1 {
				count++
			}
		}
	}
	fmt.Println("Количество нечетных элементов:", count)
}
