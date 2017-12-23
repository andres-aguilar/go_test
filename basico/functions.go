package main

import "fmt"

var variable_global int32

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func main() {
	variable_global = 50
	fmt.Println(variable_global)

	saludar("Andres")
}
