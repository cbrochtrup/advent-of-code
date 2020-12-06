package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	questions = regexp.MustCompile("[a-z]")
	filename  = flag.String("f", "input-data", "path of file to load")
)

func nextRecord(data []byte) (int, []byte) {
	if len(data) == 0 {
		return 0, nil
	}
	delim := "\n\n"
	delimIdx := bytes.Index(data, []byte(delim))
	boundaryIdx := delimIdx + len(delim)
	if delimIdx == -1 {
		boundaryIdx = len(data)
	}
	record := data[:boundaryIdx]
	return parseGroup(string(record)), data[boundaryIdx:]
}

func parseGroup(record string) int {
	fmt.Println(removeDuplicateValues(questions.FindAllString(record, -1)))
	return len(removeDuplicateValues(questions.FindAllString(record, -1)))
}

func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("error opening data file: %v\n", err)
		os.Exit(1)
	}
	questionCount := 0
	remaining := data[:]
	nQuestions := 0
	for len(remaining) > 0 {
		nQuestions, remaining = nextRecord(remaining)
		questionCount += nQuestions
	}
	fmt.Printf("Unique questions from group: %d\n", questionCount)
}
