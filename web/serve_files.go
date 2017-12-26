package main

import "net/http"

func main()  {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
  })

  staticFiles := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

  http.ListenAndServe(":8080", nil)
}
