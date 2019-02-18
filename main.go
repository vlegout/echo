package main

import (
  "encoding/json"
  "net/http"
  "strings"
)

type Response struct {
  Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")

  js, _ := json.Marshal(Response{message})

  w.Write(js)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8888", nil)
}
