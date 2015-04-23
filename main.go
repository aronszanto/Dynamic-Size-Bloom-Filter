package main

import "StaticFilter"

func main() {
	filter = StaticFilter.NewFilter(.01, 5)
	lines, err = ReadLines("Dictonaries/tinydict.txt")
	for i := 0; i < len(lines); i++ {
		filter.Insert()
	}
}
