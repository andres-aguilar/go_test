package main

import (
  "net/http"
  "io"
)

func main() {
  http.HandleFunc("/", handler)

  http.HandleFunc("/hola", func (w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hola mundo desde otra URL")
  })

  http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
  io.WriteString(w, "Hola mundo")
}
