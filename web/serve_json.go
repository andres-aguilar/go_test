package main

import (
  "net/http"
  "encoding/json"
)

type Curso struct {
  /* Los atributos privados no serán convertidos en JSON*/
  Title string `json:"title"`
  NumVideos int `json:"num_videos"`
}

type Cursos []Curso

func main()  {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    cursos := Cursos{
      Curso{"Curso de GO", 28},
      Curso{"Curso de Ruby", 10},
      Curso{"Curso de Java SE", 30},
      Curso{"Curso de Java EE", 50},
    }

    json.NewEncoder(w).Encode(cursos)
  })
  http.ListenAndServe(":8081", nil)
}
