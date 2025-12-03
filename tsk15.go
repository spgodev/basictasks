package main

import (
	"strings"
)

func Grades(tableGrades []int) string {
	n := len(tableGrades)
	results := make([]string, n) // заранее создаём срез нужной длины

	for i, grade := range tableGrades {
		switch grade {
		case 1:
			results[i] = "Very Bad"
		case 2:
			results[i] = "Bad"
		case 3:
			results[i] = "Average"
		case 4:
			results[i] = "Good"
		case 5:
			results[i] = "Excellent"
		default:
			results[i] = "Unknown grade"
		}
	}

	return strings.Join(results, ", ")
}
