package main

import "fmt"

func main()  {
  channel := make(chan string)

  go func (channel chan string)  {
    for {
      var nombre string
      fmt.Scanln(&nombre)
      channel <- nombre
    }
  } (channel)

  mensaje := <- channel
  fmt.Println("El mensaje es: " + mensaje)
}
