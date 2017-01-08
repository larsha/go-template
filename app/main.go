package main

import(
  "os"
  "fmt"
  "log"
  "net/http"
  "runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func httpHandler(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  port := os.Getenv("PORT")
  handler := httpHandler(http.DefaultServeMux)

  // Routes
  http.HandleFunc("/", indexHandler)

  http.ListenAndServe(":" + port, httpHandler)
}
