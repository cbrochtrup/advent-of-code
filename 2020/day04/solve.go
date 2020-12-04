package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func field_info(field string) (string, string) {
	field_info := strings.Split(field, ":")
	// field name, value
	return field_info[0], field_info[1]
}

func remove(str_arr []string, to_remove string) []string {
	for i, s := range str_arr {
		if s == to_remove {
			str_arr[len(str_arr)-1], str_arr[i] = s, string(str_arr[len(str_arr)-1])
			return str_arr[:len(str_arr)-1]
		}
	}
	return str_arr
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
	if err != nil {
		log.Fatal(err)
	}
	data := string(b)

	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// optional: cid

	line_data := strings.Split(data, "\n")
	req_fields := required_fields
	n_valid := 0
	for _, line := range line_data {
		info := strings.Fields(line)
		for _, item := range info {
			// fmt.Printf("%q, %q\n", req_fields, strings.Split(item, ":")[0])
			req_fields = remove(req_fields, strings.Split(item, ":")[0])
		}

		// Passport info ends on blank lines
		if len(info) < 1 {
			if len(req_fields) == 0 {
				n_valid++
			}
			// Reset checklist for next passport
			req_fields = required_fields
		}
	}
	fmt.Println(n_valid)
}
