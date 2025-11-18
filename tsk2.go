package main

import (
	"fmt"
)

func main() {
	even := []int64{-7, -4, 1, 3, 4, 8}
	var count int64 = 0
	{
		for _, v := range even {
			if v%2 == 0 {
				count++
			}
		}
	}
	fmt.Println("Количество четных элементов:", count)
}
