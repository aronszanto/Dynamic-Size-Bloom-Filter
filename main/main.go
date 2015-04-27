package main

import "github.com/aszanto9/Blumo/scalablefilter"
import "github.com/aszanto9/Blumo/staticfilter"
import "bufio"
import "os"

import "fmt"

func main() {
	filter := ScalableFilter.New(.001)

	inserted := InsertLines(filter, "Dictionaries/bigdict.txt")
	//lines := []string{"aron", "grace", "joe", "joseph", "kai ri"}

	test := fmt.Sprint("Inserted ", inserted, " entries.\n Looking for grace: ", filter.Lookup([]byte("grace")), "\n\nLooking for azazaz: ", filter.Lookup([]byte("azazaz")), "\n\n")
	fmt.Printf(test)
}

// referenced http://stackoverflow.com/questions/5884154
func InsertLines(filter *StaticFilter.Filter, path string) int {
	file, err := os.Open(path)
	fmt.Printf("Attempting to open file...\n")
	if err != nil {
		fmt.Printf("File open failed.\n")
		panic(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		//op := fmt.Sprint("Inserting ", lines[i], "...\n")
		//fmt.Printf(op)
		filter.Insert([]byte(scanner.Text()))
		count++
	}
	return count
}
