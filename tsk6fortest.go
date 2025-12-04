package main

// название ни о чем не говорит
func SearchMaxLenString(names []string) string {
	if len(names) == 0 {
		return ""
	}

	// camelCase
	maxString := names[0]
	maxLen := len(maxString)

	for _, v := range names {
		if len(v) > maxLen {
			maxString = v
			maxLen = len(v)
		}
	}

	return maxString
}
