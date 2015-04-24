package main

import "github.com/josephwandile/Blumo/StaticFilter"
import "bufio"
import "os"

import "fmt"

func main() {
	filter := StaticFilter.NewFilter(1000, .01)

	//lines, _ := ReadLines("./Dictonaries/tinydict.txt")
	lines := []string{"aron", "grace", "joe", "joseph", "kai ri"}

	for i := 0; i < len(lines); i++ {
		op := fmt.Sprint("Inserting ", lines[i], "...\n")
		fmt.Printf(op)
		filter.Insert([]byte(lines[i]))
	}

	fmt.Printf("\n\n")
	test := fmt.Sprint("Looking for aron: ", filter.Lookup([]byte("aron")), "\n\nLooking for matt: ", filter.Lookup([]byte("matt")), "\n\n")
	fmt.Printf(test)
}

// http://stackoverflow.com/questions/5884154
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	fmt.Printf("Attempting to open file...\n")
	if err != nil {
		fmt.Printf("File open failed.\n")
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
