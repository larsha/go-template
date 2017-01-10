package main

import(
  "fmt"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := Set("foo", "bar")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  val, err := Get("foo")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", val)
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", middlewareHandler(http.DefaultServeMux))
}
