package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func field_info(field string) (string, string) {
	f := strings.Split(field, ":")
	// field name, data
	return f[0], f[1]
}

func valid_data(field string, data string) bool {
	switch f := field; f {
	case "byr":
		d, _ := strconv.Atoi(data)
		return d >= 1920 && d <= 2003
	case "iyr":
		d, _ := strconv.Atoi(data)
		return d >= 2010 && d <= 2020
	case "eyr":
		d, _ := strconv.Atoi(data)
		return d >= 2020 && d <= 2030
	case "hgt":
		units := string(data[len(data)-2:])
		if units == "cm" {
			d, _ := strconv.Atoi(string(data[:len(data)-2]))
			return d >= 150 && d <= 193
		} else if units == "in" {
			d, _ := strconv.Atoi(string(data[:len(data)-2]))
			return d >= 59 && d <= 76
		} else {
			return false
		}
	case "hcl":
		ret, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, data)
		return ret
	case "ecl":
		valid_colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, c := range valid_colors {
			if c == data {
				return true
			}
		}
		return false
	case "pid":
		ret, _ := regexp.MatchString(`^[0-9]{9}$`, data)
		return ret
	default:
		return false
	}
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

	line_data := strings.Split(data, "\n")
	req_fields := required_fields
	n_valid := 0
	for _, line := range line_data {
		info := strings.Fields(line)
		for _, item := range info {
			field, data := field_info(item)
			if valid_data(field, data) {
				req_fields = remove(req_fields, strings.Split(item, ":")[0])
			}
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
