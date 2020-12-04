package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "strings"
    "strconv"
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
    password_1 := strings.TrimSpace(split_line[1])
    password := split_line[1]
    if (password_1 != password) {
      fmt.Printf("%q %q\n", password_1, password)
    }
    ind_1, _ := strconv.Atoi(policy_numbers[0])
    ind_2, _ := strconv.Atoi(policy_numbers[1])
    ind_1 -= 1
    ind_2 -= 1
    cond1 := (string(password[ind_1]) == policy[1]) && (string(password[ind_2]) != policy[1])
    cond2 := (string(password[ind_1]) != policy[1]) && (string(password[ind_2]) == policy[1])
    if (cond1 || cond2) {
       n_valid++
       // fmt.Println(cond1, cond2, password, string(password[ind_1]), string(password[ind_2]), policy[1])
    }
  }
  fmt.Println(n_valid)
}
