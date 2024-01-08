package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  var result string
  meep := make(map[int][]int)

  file, _ := os.ReadFile(os.Args[1])

  input := strings.Split(string(file), "\n")
  input = input[:len(input) - 1]

  for _, line := range input { fmt.Println(line) }

  for x, line := range input {
    tiles := strings.Split(line, "|")
    for _, tile := range tiles {
      work := strings.ReplaceAll(tile, " ", "")
      if work == "X" {
        meep[x] = append(meep[x], 99)
        continue
      }
      w, err := strconv.Atoi(work)
      if err != nil { continue }
      meep[x] = append(meep[x], w)
    }
  }

  //correct location = x * 5 + 1 + y

  fmt.Println(meep[3][3])
  fmt.Println(result)
}
