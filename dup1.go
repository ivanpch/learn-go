// Выводит текст каждой строки, которая появляется в стандартном вводе более
// одного раза, а также количество ее появлений.
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    counts[input.Text()]++
  }
  for line, count := range counts {
    if count > 1 {
      fmt.Printf("%d\t%s\n", count, line)
    }
  }
}
