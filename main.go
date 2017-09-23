package main

import (
  "suffixtree"
)

func main() {
  //fmt.Printf(str[0])
  suffix := suffixtree.New()
  suffix.Build("abcabxabcd")
  suffix.Print("abcabxabcd")
}
