package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
  "suffixtree"
)

func main() {
  //fmt.Printf(str[0])
  suffix := suffixtree.New()
  //suffix.Build("中国人中国狗中国人的")
  //suffix.Build("今天天气不错我请假了今天")

  ReadLine("../PoetryGen/poetry", suffix.Build)

  //suffix.Build(str)
  //fmt.Println(str)
}

func ReadLine(fileName string, handler func(string)) error {
    f, err := os.Open(fileName)
    if err != nil {
        return err
    }
    buf := bufio.NewReader(f)
    for {
        line, err := buf.ReadString('\n')
        // preprocess for each line
        line = strings.TrimSpace(line)
        line = strings.Replace(line, " ", "", -1)
        line = strings.Replace(line, "，", "", -1)
        line = strings.Replace(line, "。", "", -1)
        line = strings.Replace(line, "　", "", -1)
        if line != "" {
          fmt.Println(line)
          handler(line)
        }
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
    return nil
}
