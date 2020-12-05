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

// row is the first index and col is the second index
func RemoveSeat(array [][]string, row int, col int) [][]string {
	d := strconv.Itoa(col)
	array[row][0] = strings.Replace(array[row][0], d, "", -1)
	return array
}

func GetSeatId(seatSpecification string, seatMap [][]string) (int, [][]string) {
	m := make(map[string]int)
	m["rStart"] = 0
	m["rStop"] = 127
	m["cStart"] = 0
	m["cStop"] = 7
	for _, spec := range seatSpecification {
		m = BinaryPartition(string(spec), m)
	}
	seatMap = RemoveSeat(seatMap, m["rStart"], m["cStart"])
	return m["rStart"]*8 + m["cStart"], seatMap
}

func BinaryPartition(split string, m map[string]int) map[string]int {
	newRow := m["rStart"] + (m["rStop"]-m["rStart"])/2
	newCol := m["cStart"] + (m["cStop"]-m["cStart"])/2
	switch f := split; f {
	case "F":
		m["rStop"] = newRow
	case "B":
		m["rStart"] = newRow + 1
	case "R":
		m["cStart"] = newCol + 1
	case "L":
		m["cStop"] = newCol
	}
	return m
}

func main() {
	filename := flag.String("f", "test_input", "input file")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	data := string(b)

	lineData := strings.Split(data, "\n")
	maxSeatId := 0
	seatId := 0
	var seatMap [][]string
	for i := 0; i < 128; i++ {
		seatMap = append(seatMap, []string{"01234567"})
	}
	for _, line := range lineData {
		seatId, seatMap = GetSeatId(line, seatMap)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	fmt.Println("maxSeatId", maxSeatId)
	fmt.Println(seatMap)
	for i, seats := range seatMap {
		if len(seats[0]) == 1 {
			fmt.Println("my row is", i)
			fmt.Println("my col is", seats)
		}
	}
}
