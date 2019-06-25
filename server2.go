// Минимальный echo-сервер со счетчиком запросов.
package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

var mutex sync.Mutex
var count int

func main() {
  http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработки возвращает компонент пути из URL запроса.
func handler(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  count++
  mutex.Unlock()
  fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
  mutex.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mutex.Unlock()
}
