package main

import (
	"fmt"
)

func strings() {
	names := []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"}
	var (
		maxstring       string = names[0]
		maxcountsymbols int64  = 0
		currentlens     int64  = 0
	)
	for _, v := range names {
		for range v {
			currentlens++

		}
		if currentlens > maxcountsymbols {
			maxcountsymbols = currentlens
			maxstring = v
		}
		currentlens = 0
	}
	if maxcountsymbols%2 == 0 {
		fmt.Println("Максимальная четное имя", maxstring)
	} else {
		fmt.Println("Максимальное нечетное имя", maxstring)
	}
}
