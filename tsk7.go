package main

import (
	"fmt"
)

func swap() {
	rot := []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48}
	var (
		temp int64
		sum  int64
		half int64
	)
	for range rot {
		sum++
	}
	half = sum / 2
	for i := 0; i < int(half); i++ {
		temp = rot[i]
		rot[i] = rot[int(sum)-1-i]
		rot[int(sum)-1-i] = temp
	}
	fmt.Println(rot)
}
