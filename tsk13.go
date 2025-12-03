package main

func SeparationParity(slice []int64) ([]int64, []int64) {
	var (
		evenCount int
		oddCount  int
	)
	for _, v := range slice {
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

	for _, v := range slice {
		if v%2 == 0 {
			even[eIdx] = v
			eIdx++
		} else {
			odd[oIdx] = v
			oIdx++
		}
	}

	return even, odd
}
