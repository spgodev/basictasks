package main

func delldbl(deldbl []int64) []int64 {
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
	return deldbl
}
