package main

import "fmt"

func alt() {
	delodd := []int64{-98, -7, -4, 1000, 3, 4, 8, 34, 48}
	var (
		podsnos int64 = 0
	)
	for range delodd {
		podsnos++
	}
	for i := 0; i < int(podsnos); i++ {
		if delodd[i]%2 == 0 {
			delodd[i] = delodd[int(podsnos)-1]
			podsnos--
		}
		delodd = delodd[:int(podsnos)]
	}
	fmt.Println(delodd)
}
