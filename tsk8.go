package main

func RemoveOdd(nums []int64) []int64 {
	removeCount := int64(0)

	for _, v := range nums {
		if v%2 != 0 {
			nums[removeCount] = v
			removeCount++
		}
	}
	nums = nums[:removeCount]

	return nums
}
