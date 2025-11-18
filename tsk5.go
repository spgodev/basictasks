package main

import (
	"fmt"
)

func bigone() {
	num := []int64{-98, -7, -4, 47, 3, 4, 8, 34, 48}
	var maxvalue int64 = num[0]

	for _, v := range num {
		if v > maxvalue {
			maxvalue = v
		}
	}
	fmt.Println("Максимальное значение", maxvalue)
}
