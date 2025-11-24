package main

func bigone(num []int64) int64 {
	var (
		maxvalue int64 = num[0]
	)

	for _, v := range num {
		if v > maxvalue {
			maxvalue = v
		}
	}
	return maxvalue
}
