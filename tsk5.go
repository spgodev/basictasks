package main

// 1. название функции не отражет поведение функции
// 2. В функции есть ошибка.
// Написал тесты, которые помогут тебе ее найти
func bigone(num []int64) int64 {
	var (
		maxvalue int64 = num[0]
	)

	for _, v := range num {
		//camelCase
		if v > maxvalue {
			maxvalue = v
		}
	}
	return maxvalue
}
