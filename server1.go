// Минимальный echo-сервер.
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработки возвращает компонент пути из URL запроса.
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}
