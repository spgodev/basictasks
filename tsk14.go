package main

func foundindex(indexf []int64, zapros int64) int {
	//зачем этот var?
	var ()
	for i, v := range indexf {
		if v == zapros {
			return i
		}
	}
	return -1
}
