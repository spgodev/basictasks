package main

func Reverse(nums []int64) []int64 {
	// можно использовать len

	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	return nums
}
