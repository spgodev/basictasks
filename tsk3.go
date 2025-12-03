package main

// поправь замечания из предыдущиз задач
func CountOdd(nums []int64) int {
	count := 0
	{
		for _, v := range nums {
			if v%2 != 0 {
				count++
			}
		}
	}
	return count
}
