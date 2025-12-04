package main

// input
func Containes(names []string, input string) bool {

	for _, v := range names {
		if v == input {
			return true
		}
	}
	return false
}
