package main

func RemoveDuplicate(input []int64) []int64 {
	if len(input) == 0 {
		return input
	}
	remover := 1
	for i := 1; i < len(input); i++ {
		if input[i] != input[remover-1] {
			input[remover] = input[i]
			remover++
		}
	}
	input = input[:remover]
	return input
}
