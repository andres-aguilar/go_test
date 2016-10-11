package main

import "fmt"

/*
 * Manejo de variables 
 * Al crear una variable con var, se le asigna un valor pordefault
 * 0 para enteros, vacio para cadenas
 * notar que se usa la palabra reservada 'var' seguido del nombre o nombres
 * de las variables y después viene el tipo de dato que van a almacenar
 */

func main() {
    var nombre string 
    var edad int

    fmt.Println(nombre)
    fmt.Println(edad)

    // tambien se pueden crear variables e inicializarlas en la misma
    // linea con el operador :=
    nombre2 := "Andres Aguilar"
    fmt.Println(nombre2)
}