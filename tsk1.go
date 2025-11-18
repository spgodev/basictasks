package main

import "fmt"

func main() {
	summa := []float64{1.1, 1.3, 1.8, 2.0, 3.6}

	var plus float64 = 0

	for _, v := range summa {
		plus += v
	}
	fmt.Println("Сумма элементов:", plus)
}
