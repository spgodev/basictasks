package main

// 1. Используй camelCase
// 2. even не подходит, тк у тебя могут в функцию передаваться не только четные чисда
func counteven(even []int64) int64 {
	var (
		count int64
	)
	{
		for _, v := range even {
			if v%2 == 0 {
				count++
			}
		}
	}
	return count
}
