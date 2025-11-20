package main

import (
	"fmt"
)

func grades() {
	grades := []int{1, 2, 3, 4, 5}

	for _, grade := range grades {
		switch grade {
		case 1:
			fmt.Println("Very Bad")
		case 2:
			fmt.Println("Bad")
		case 3:
			fmt.Println("Average")
		case 4:
			fmt.Println("Good")
		case 5:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown grade")
		}
	}
}
