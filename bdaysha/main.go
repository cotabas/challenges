package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
  //bday := "b189e135d9102763d0316d8ab0e7f12445405b2ac57c52ea5fde6984c75e9328"
  bday := "6c833595c7502119465895442b340b3118a8c1aec222882d35e364f75e57b268"
  test := sha256.New()

  for y := 1900; y <= 2023; y++ {
    yyyy := strconv.Itoa(y)
    for m := 1; m <= 12; m++ {
      mm := strconv.Itoa(m)
      if len(mm) == 1 { mm = "0" + mm }

      for d := 1; d <= 31; d++ {
        dd := strconv.Itoa(d)
        if len(dd) == 1 { dd = "0" + dd }
        date := mm + "/" + dd + "/" + yyyy
        test.Write([]byte(date))
        huh := test.Sum(nil)
        hext := make([]byte, hex.EncodedLen(len(huh)))
        hex.Encode(hext, huh)

        if string(hext) == bday { fmt.Println(date) }
        test = sha256.New()
        //fmt.Println(string(hext))
      }
    }
  }
}
