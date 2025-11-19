package main

import (
	"fmt"
)

func dellodd() {
	delodd := []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48}
	var (
		podsnos = 0
	)

	for _, v := range delodd {
		if v%2 != 0 {
			delodd[podsnos] = v
			podsnos++
		}
	}
	delodd = delodd[:podsnos]

	fmt.Println("Результат:", delodd)
}
