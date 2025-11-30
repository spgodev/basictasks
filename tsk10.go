package main

func dubl(rpt1 []int64) []int64 {
	rpt2 := make([]int64, len(rpt1)*2)
	var (
		count int
	)
	// реши используя стандартный индекс из for i,value := range rpt1
	for _, value := range rpt1 {
		rpt2[count] = value
		count++
		rpt2[count] = value
		count++
	}
	return rpt2
}
