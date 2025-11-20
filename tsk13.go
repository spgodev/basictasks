package main

import (
	"fmt"
)

func sortirovka() {
	sort := []int64{1, 2, 3, 4, 6, 7, 8, 10}
	var (
		evenCount int
		oddCount  int
	)
	for _, v := range sort {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	even := make([]int64, evenCount)
	odd := make([]int64, oddCount)
	eIdx := 0
	oIdx := 0

	for _, v := range sort {
		if v%2 == 0 {
			even[eIdx] = v
			eIdx++
		} else {
			odd[oIdx] = v
			oIdx++
		}
	}

	fmt.Println(even, odd)
}
