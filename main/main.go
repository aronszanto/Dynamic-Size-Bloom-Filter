package main

//import "github.com/aszanto9/Blumo/scalablefilter"

/*
This exists for testing.
import "github.com/aszanto9/Blumo/staticfilter"
*/

import "bufio"
import "os"

//import "github.com/davecheney/profile"
import "github.com/aszanto9/Blumo/scalablefilterpartition"
import "fmt"

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	filter := ScalableFilterPartition.NewFilter(.001)

	inserted := InsertLines(filter, "../Dictionaries/tinydicttest.txt")
	//lines := []string{"aron", "grace", "joe", "joseph", "kai ri"}

	test := fmt.Sprint("Inserted ", inserted, " entries.\n Looking for grace: ", filter.Lookup([]byte("grace")), "\n\nLooking for qwertyuiop: ", filter.Lookup([]byte("qwertyuiop")), "\n\n")
	fmt.Printf(test)
}

// referenced http://stackoverflow.com/questions/5884154
func InsertLines(filter *ScalableFilterPartition.SBF, path string) int {
	file, err := os.Open(path)
	fmt.Printf("Attempting to open file...\n")
	if err != nil {
		fmt.Printf("File open failed.\n")
		panic(err)
		return 0
	}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		op := fmt.Sprint("Inserting ", l, "\n")
		fmt.Printf(op)
		filter.Insert([]byte(l))
		count++
	}
	file.Close()
	return count
}
