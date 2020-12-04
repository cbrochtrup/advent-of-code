package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "strings"
    "strconv"
    "regexp"
)

func main() {
  fmt.Println("hello world")
  file, err := os.Open("input.txt")
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
  line_data := strings.Split(data, "\n")
  n_valid := 0
  for _, line := range line_data {
    split_line := strings.Split(line, ":")
    policy := strings.Fields(split_line[0])
    if len(split_line) < 2 {
      fmt.Println("invalid line")
      continue
    }

    policy_numbers := strings.Fields(strings.Replace(policy[0], "-", " ", -1))
    reg := regexp.MustCompile(policy[1])
    password := split_line[1]
    matches := reg.FindAllString(password, -1)
    min_let, _ := strconv.Atoi(policy_numbers[0])
    max_let, _ := strconv.Atoi(policy_numbers[1])
    if (len(matches) >= min_let && len(matches) <= max_let) {
       n_valid++
    }
  }
  fmt.Println(n_valid)
}
