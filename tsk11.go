package main

func dellel(delel []int64, idx int) []int64 {
	if idx < 0 || idx >= len(delel) {
		return delel
	}
	for i := idx; i < len(delel)-1; i++ {
		delel[i] = delel[i+1]
	}
	delel = delel[:len(delel)-1]
	return delel
}
