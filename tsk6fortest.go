package main

func stringsfortest(names []string) string {
	if len(names) == 0 {
		return ""
	}

	maxstring := names[0]
	maxlen := len(maxstring)

	for _, v := range names {
		if len(v) > maxlen {
			maxstring = v
			maxlen = len(v)
		}
	}

	return maxstring
}
