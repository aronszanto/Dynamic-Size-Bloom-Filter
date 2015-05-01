package main

import "fmt"
import "os"
import "bufio"

//import "testing"
import "github.com/aszanto9/Blumo/scalablefilterpartition"

//import "github.com/aszanto9/Blumo/staticfilter"
//import "github.com/aszanto9/Blumo/staticfilterpartition"

func main() {
	dict := init_dict()
	not_in_dict := init_not_dict()

	test_scalable_filter(dict, not_in_dict)

}

func init_dict() []string {
	var d []string
	file, err := os.Open("../Dictionaries/1149891.txt")
	fmt.Printf("Attempting to open dictionary file...\n")
	if err != nil {
		fmt.Printf("File open failed.\n")
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d = append(d, scanner.Text())
	}
	file.Close()
	return d
}

func init_not_dict() []string {
	var nd []string
	file, err := os.Open("../Dictionaries/not_in_dict.txt")
	fmt.Printf("Attempting to open dictionary file...\n")
	if err != nil {
		fmt.Printf("File open failed.\n")
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nd = append(nd, scanner.Text())
	}
	file.Close()
	return nd
}

func test_scalable_filter(dict, not_in_dict []string) {
	fmt.Printf("Creating scalable, partitioned filter with error bound 0.1 percent...\n")
	filter := ScalableFilterPartition.NewFilter(.001)
	fmt.Printf("Filter created. Testing filter...\n")
	fmt.Printf("Inserting dictionary into scalable filter...\n")
	for i := range dict {
		filter.Insert([]byte(dict[i]))
	}
	fmt.Printf("Insert complete.\nTesting false positive rate...\n")

	false_pos := 0
	for i := range not_in_dict {
		if filter.Lookup([]byte(not_in_dict[i])) {
			false_pos++
		}
	}

	fmt.Printf(fmt.Sprint("Number of false positives: ", false_pos,
		"\nRate of false positives: ", (float64(false_pos)/float64(len(not_in_dict)))*100, " percent\n"))

}
