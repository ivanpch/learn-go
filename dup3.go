// Выводит текст каждой строки, которая появляется более одного раза. Чтение
// производится из списка файлов.
package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
      continue
    }
    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
  }
  for line, count := range counts {
    if count > 1 {
      fmt.Printf("%s\t%d\n", line, count)
    }
  }
}
