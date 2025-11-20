package main

import (
	"fmt"
)

func foundindex() {
	indexf := []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48}
	var (
		zapros int64 = 1000
		found  bool  = false
	)
	for i, v := range indexf {
		if v == zapros {
			fmt.Println(i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println(-1)
	}
}
