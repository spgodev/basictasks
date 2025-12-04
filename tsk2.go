package main

// 1. Используй camelCase
// 2. even не подходит, тк у тебя могут в функцию передаваться не только четные чисда
func countEven(nums []int64) int {
	count := 0
	{
		for _, v := range nums {
			if v%2 == 0 {
				count++
			}
		}
	}
	return count
}
