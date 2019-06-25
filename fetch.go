// Выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
  "strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
    if strings.HasPrefix(url, prefix) != true {
      url = prefix + url
    }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
      fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
    }
		resp.Body.Close()
    fmt.Printf("fetch: чтение %s: статус - %s\n", url, resp.Status)
	}
}
