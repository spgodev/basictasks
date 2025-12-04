package main

func RemoveElement(nums []int64, idx int) []int64 {
	if idx < 0 || idx >= len(nums) {
		return nums
	}
	for i := idx; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums = nums[:len(nums)-1]
	return nums
}
