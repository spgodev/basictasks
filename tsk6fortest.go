package main

// название ни о чем не говорит
func stringsfortest(names []string) string {
	if len(names) == 0 {
		return ""
	}

	// camelCase
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
