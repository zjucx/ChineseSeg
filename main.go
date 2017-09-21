package main

import (
  "fmt"
)

func main() {
  str := "我们是中国人"
  r := []rune(str)
  fmt.Println("rune=", string(r[0]))
  //fmt.Printf(str[0])
}
