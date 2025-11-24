package main

func swap(rot []int64) []int64 {
	var (
		temp int64
		sum  int64
		half int64
	)
	for range rot {
		sum++
	}
	half = sum / 2
	for i := 0; i < int(half); i++ {
		temp = rot[i]
		rot[i] = rot[int(sum)-1-i]
		rot[int(sum)-1-i] = temp
	}
	return rot
}
