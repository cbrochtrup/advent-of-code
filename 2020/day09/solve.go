package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filename = flag.String("f", "input-data", "path of file to load")
)

func splitFunc(c rune) bool {
	return c == '\n'
}

func readDataArr(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	strData := string(b)
	return strings.FieldsFunc(strData, splitFunc)
}

func convertToInt(data []string) []int {
	var retSlice []int
	for _, s := range data {
		toAppend, _ := strconv.Atoi(s)
		retSlice = append(retSlice, toAppend)
	}
	return retSlice
}

func updateCipher(input []int, size int, index int) ([]int, int, int) {
	return input[index : index+size], index + 1, input[index+size]
}

func checkValid(sumVals []int, toCheck int) bool {
	for i, v1 := range sumVals {
		for j, v2 := range sumVals {
			if i != j && v1+v2 == toCheck {
				return true
			}
		}
	}
	return false
}

func solve1(data []int) int {
	// get preamble
	size := 25
	for i := 0; i < len(data)-size-1; i++ {
		curNums, _, toCheck := updateCipher(data, size, i)
		// check next value
		if !checkValid(curNums, toCheck) {
			fmt.Printf("%d cannot be made from %s\n", toCheck, curNums)
			return toCheck
		}
	}
	return -1
}

func cdfCalc(nums []int) []int {
	retCdf := []int{nums[0]}
	for _, val := range nums[1:] {
		retCdf = append(retCdf, retCdf[len(retCdf)-1]+val)
	}
	return retCdf
}

func checkSums(cdf []int, rawData []int, size int, num int) []int {
	for i := 0; i < len(cdf)-size; i++ {
		if cdf[i+size]-cdf[i] == num {
			return rawData[i : i+size+1]
		}
	}
	return []int{-1}
}

func sumMinMax(vals []int) int {
	max := 0
	min := 999999999999
	for _, v := range vals {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max + min
}

func solve2(data []int, magicNum int) {
	cdf := cdfCalc(data)
	// loop through contiguous ranges to check their sums via cdf differences
	for size := 2; size < len(cdf)-25; size++ {
		sumVals := checkSums(cdf, data, size, magicNum)
		if sumVals[0] > -1 {
			fmt.Println("Cypher key is", sumMinMax(sumVals))
		}
	}
}

func main() {
	flag.Parse()
	lineData := readDataArr(*filename)
	intData := convertToInt(lineData)

	invalidNum := solve1(intData)
	solve2(intData, invalidNum)
}
