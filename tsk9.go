package main

func zapros(names []string, imput string) bool {

	for _, v := range names {
		if v == imput {
			return true
		}
	}
	return false
}
