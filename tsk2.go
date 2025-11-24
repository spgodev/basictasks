package main

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
