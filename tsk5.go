package main

// 1. название функции не отражет поведение функции
// 2. В функции есть ошибка.
// Написал тесты, которые помогут тебе ее найти
func MaxValueSearch(num []int64) int64 {
	if len(num) == 0 {
		return 0
	}

	maxValue := num[0]

	for _, v := range num {
		//camelCase
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
