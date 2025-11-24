package main

func dellodd(delodd []int64) []int64 {
	var (
		podsnos = 0
	)

	for _, v := range delodd {
		if v%2 != 0 {
			delodd[podsnos] = v
			podsnos++
		}
	}
	delodd = delodd[:podsnos]

	return delodd
}
