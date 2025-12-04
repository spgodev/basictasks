package main

func FoundIndex(nums []int64, input int64) int {
	//зачем этот var?
	for i, v := range nums {
		if v == input {
			return i
		}
	}
	return -1
}
