package main

import (
	"fmt"
)

func dellel() {
	delel := []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48}
	var (
		idx int = 2
	)
	if idx >= len(delel) {
		fmt.Println("idx out of range")
	} else {
		for i := idx; i < len(delel)-1; i++ {
			delel[i] = delel[i+1]
		}
	}
	delel = delel[:len(delel)-1]
	fmt.Println(delel)
}
