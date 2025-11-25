package main

func strings1(names []string) string {
	var (
		maxstring       string = names[0]
		maxcountsymbols int64  = 0
		currentlens     int64  = 0
	)
	for _, v := range names {
		for range v {
			currentlens++

		}
		if currentlens > maxcountsymbols {
			maxcountsymbols = currentlens
			maxstring = v
		}
		currentlens = 0
	}
	return maxstring
}
