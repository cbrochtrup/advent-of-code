package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
        "strings"
        "strconv"
        "github.com/golang-collections/collections/set"
)


var (
	filename  = flag.String("f", "input-data", "path of file to load")
        dag       = make(map[string]*set.Set)
        dagInt    = make(map[string]map[string]int)
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

func addToGraph(dag map[string]*set.Set, dagI map[string]map[string]int, bagType string, contained []string, nCont []int) (map[string]*set.Set, map[string]map[string]int) {
  for i, c := range contained {
    if _, ok := dag[bagType]; ok {
      dag[bagType].Insert(c)
      dagI[bagType][c] = nCont[i]
    } else {
      dag[bagType] = set.New(c)
      newVer := make(map[string]int)
      newVer[c] = nCont[i]
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
      fmt.Println("no bads inside", s)
      return  []string{""}, []int{0}
    }
    retBags = append(retBags, strings.Join(flds[1:3], " "))
    retInts = append(retInts, num)
  }
  return retBags, retInts
}

func searchDag(dagI map[string]map[string]int, contains *set.Set, otherList []string, targetBag string) {
  curLen := contains.Len()
  for bag, inside := range dagI {
    if val, ok := inside[targetBag]; ok {
      contains.Insert(bag)
      otherList = append(otherList, bag)
      nBags += val
    }
  }
  if curLen < contains.Len() {
    for _, k := range otherList {
      searchDag(dagI, contains, otherList, k)
    }
  } else {
    fmt.Printf("no new bags, total %d\n", contains.Len())
  }
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
	searchDag(dagInt, set.New(nil), []string{}, "shiny gold")
}
