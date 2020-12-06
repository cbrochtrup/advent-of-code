package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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
	record := data[:boundaryIdx-len(delim)]
	return parseGroup(string(record)), data[boundaryIdx:]
}

func parseGroup(record string) int {
	return countCommonValue(strings.Fields(record))
}

func countCommonValue(strSlice []string) int {
	keys := make(map[string]int)
	nFields := len(strSlice)
	for _, entry := range strSlice {
		sliceChars := removeDuplicateValues(questions.FindAllString(entry, -1))
		for _, char := range sliceChars {
                  keys[string(char)]++
                }
	}
	nCommon := 0
	for _, count := range keys {
		if count == nFields {
			nCommon++
		}
	}
	return nCommon
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
	fmt.Printf("Questions answered yes by all: %d\n", questionCount)
}
