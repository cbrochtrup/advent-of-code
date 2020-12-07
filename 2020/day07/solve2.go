package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"github.com/golang-collections/collections/set"
)

        // "github.com/heimdalr/dag"
var (
	filename = flag.String("f", "input-data", "path of file to load")
	dag      = make(map[string]*set.Set)
	dagInt   = make(map[string]map[string]float64)
)

func nextRecord(data []byte, delim string) (string, []byte) {
	if len(data) == 0 {
		return "", nil
	}
	delimIdx := bytes.Index(data, []byte(delim))
	boundaryIdx := delimIdx + len(delim)
	if delimIdx == -1 {
		boundaryIdx = len(data)
	}
	record := data[:boundaryIdx]
	return string(record), data[boundaryIdx:]
}

func addToGraph(dag map[string]*set.Set, dagI map[string]map[string]float64, bagType string, contained []string, nCont []int) (map[string]*set.Set, map[string]map[string]float64) {
	for i, c := range contained {
		if _, ok := dag[bagType]; ok {
			dag[bagType].Insert(c)
			dagI[bagType][c] = float64(nCont[i])
		} else {
			dag[bagType] = set.New(c)
			newVer := make(map[string]float64)
			newVer[c] = float64(nCont[i])
			dagI[bagType] = newVer
		}
	}
	return dag, dagI
}

func getContainingBags(bagString string) ([]string, []int) {
	var retBags []string
	var retInts []int
	bs := strings.Split(bagString, ",")
	for _, s := range bs {
		flds := strings.Fields(s)
		num, err := strconv.Atoi(flds[0])
		if err != nil {
			fmt.Println("no bags inside", s)
			return []string{""}, []int{0}
		}
		retBags = append(retBags, strings.Join(flds[1:3], " "))
		retInts = append(retInts, num)
	}
	return retBags, retInts
}

func totalBags(dagNode map[string]int) int {
  totalBags := 0
  for _, bags := range dagNode {
    totalBags += bags
  }
  return totalBags
}


func sumDag(dagI map[string]map[string]float64, curBag string, initBags float64, finalBags float64) (map[string]map[string]float64, float64) {
  // Make this will fix numerical instability? no...
  var dagI2 = make(map[string]map[string]float64)
  for k, v := range dagI {
    dagI2[k] = v
  }
  for bag, degree := range dagI[curBag] {
    fmt.Printf("start bags from %s in %s is %f. Adding %f\n", curBag, bag, dagI[curBag][bag], degree)
    dagI[curBag][bag] = degree * initBags
    finalBags += dagI[curBag][bag]
    fmt.Printf("final bags from %s in %s is %f\n", curBag, bag, dagI[curBag][bag])
    dagI2, finalBags = sumDag(dagI, bag, dagI[curBag][bag], finalBags)
  }
  return dagI2, finalBags
}

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("error opening data file: %v\n", err)
		os.Exit(1)
	}
	remaining := data[:]
	var line string
	for len(remaining) > 0 {
		line, remaining = nextRecord(remaining, "\n")
		split := strings.Split(strings.TrimSpace(line), "contain")
		bagType := strings.Join(strings.Fields(split[0])[:2], " ")
		retBags, numBags := getContainingBags(split[1])
		dag, dagInt = addToGraph(dag, dagInt, bagType, retBags, numBags)
	}
	var finalBags float64
	dagInt, finalBags = sumDag(dagInt, "shiny gold", 1, 0.0)
	fmt.Println(finalBags)
}
