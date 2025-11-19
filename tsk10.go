package main

import (
	"fmt"
)

func dubl() {
	rpt1 := []int64{1, 2, 10, 15}
	rpt2 := make([]int64, len(rpt1)*2)
	var (
		count int
	)
	for _, value := range rpt1 {
		rpt2[count] = value
		count++
		rpt2[count] = value
		count++
	}
	fmt.Println(rpt2)
}
