package main

func zapros() {
	names := []string{"Max", "Dmitriy", "Alexander", "Sergey", "Anthony", "Tolyan"}
	var (
		imput  string = "Dmitriy" //  Строка для поиска
		target bool   = false
	)
	for _, v := range names {
		if v == imput {
			target = true
			break
		}
	}
	println(target)
}
