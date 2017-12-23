package main

import "fmt"

/*
 * Definición de estructuras como nuevo tipo de dato
 * y métodos para las estructuras
 */

type Usuario struct {
  nombre, apellido string
  edad int
}

func (user Usuario) saludar() string {
  return "Hola, me llamo " + user.nombre + " " + user.apellido
}

/*
 * los parametros se pasan por valor no por referencia
 * asi que, se crea una copia del 'objeto' para evitar
 * perder el 'objeto' se pasa como puntero
*/
func (user *Usuario) set_name(nombre string) {
  user.nombre = nombre
}

func main() {
  user := Usuario{"Andres", "Aguilar", 27}
  /* Otra forma de crearlo
  user := new(Usuario)  // retorna un puntero a la estructura
  user.nombre = "Andres"
  user.apellido = "Aguilar"
  user.edad = 27
  */

  user.set_name("Yosh")

  fmt.Println(user.saludar())
}
