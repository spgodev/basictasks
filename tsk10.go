package main

func Duplicate(slice []int64) []int64 {
	if slice == nil {
		return nil
	}
	result := make([]int64, len(slice)*2)
	// реши используя стандартный индекс из for i,value := range rpt1
	for i, v := range slice {
		result[i*2] = v
		result[i*2+1] = v
	}

	return result
}
