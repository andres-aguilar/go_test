package main

import "fmt"

/*
 * Manejo básico de punteros
 */

func main() {
  var x, y *int

  numero := 5

  x = &numero
  y = &numero

  fmt.Println(x, y)  // Misma dirección de memoria
  fmt.Println(*x, *y)

  *x = 8
  fmt.Println(*x, *y)
}
