package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/golang-collections/collections/set"
	"io/ioutil"
	"os"
	"strings"
)

var (
	filename = flag.String("f", "input-data", "path of file to load")
	dag      = make(map[string]*set.Set)
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

func addToGraph(dag map[string]*set.Set, bagType string, contained []string) map[string]*set.Set {
	for _, c := range contained {
		if _, ok := dag[bagType]; ok {
			dag[bagType].Insert(c)
		} else {
			dag[bagType] = set.New(c)
		}
	}
	return dag
}

func getContainingBags(bagString string) []string {
	var retBags []string
	bs := strings.Split(bagString, ",")
	for _, s := range bs {
		flds := strings.Fields(s)
		retBags = append(retBags, strings.Join(flds[1:3], " "))
	}
	return retBags
}

func searchDag(dag map[string]*set.Set, contains *set.Set, otherList []string, targetBag string) {
	curLen := contains.Len()
	for bag, inside := range dag {
		if inside.Has(targetBag) {
			contains.Insert(bag)
			otherList = append(otherList, bag)
		}
	}
	if curLen < contains.Len() {
		for _, k := range otherList {
			searchDag(dag, contains, otherList, k)
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
		fmt.Printf("%q\n", getContainingBags(split[1]))
		dag = addToGraph(dag, bagType, getContainingBags(split[1]))
	}
	searchDag(dag, set.New(nil), []string{}, "shiny gold")
}
