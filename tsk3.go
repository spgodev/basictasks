package main

// поправь замечания из предыдущиз задач
func countodd(odd []int64) int64 {
	var (
		count int64
	)
	{
		for _, v := range odd {
			if v%2 != 0 {
				count++
			}
		}
	}
	return count
}
