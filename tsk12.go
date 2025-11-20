package main

import "fmt"

func delldbl() {
	deldbl := []int64{-98, -7, -7, 1000, 3, 4, 4, 4, 48}
	var (
		vs int = 1
	)
	for i := 1; i < len(deldbl); i++ {
		if deldbl[i] != deldbl[vs-1] {
			deldbl[vs] = deldbl[i]
			vs++
		}
	}
	deldbl = deldbl[:vs]
	fmt.Println(deldbl)
}
