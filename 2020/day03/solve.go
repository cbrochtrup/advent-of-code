package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "os"
    "log"
    "strings"
)

func main() {
  x := flag.Int("x", 3, "x")
  y := flag.Int("y", 1, "y")
  flag.Parse()

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
  current_pos := int(0)
  n_trees := 0
  for i, line := range line_data {
    if i % *y != 0 {
      continue
    }
    line = strings.TrimSpace(line)
    if (len(line) < 1) {
      continue;
    }
    if (string(line[current_pos]) == "#") {
      n_trees++
    }
    current_pos = (current_pos + *x) % len(line)
  }
  fmt.Println(n_trees)
}
