package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	exec := leer_archivo()
	println(exec)
}

func leer_archivo() bool{
	error := false
	archivo, err := os.Open("texto.txt")

	// el defer funciona como un finally en java
	// se asegura que lo contenido en la función se ejecute siempre
	defer func(){ archivo.Close() }()
	
	if err != nil {
		fmt.Println("ERROR: no se puede leer el archivo!")
		error = true
	}

	scanner := bufio.NewScanner(archivo)

	var i int
	for scanner.Scan(){
		i ++
		line := scanner.Text()
		fmt.Print(i)
		fmt.Println(line)
	}

	return error
}
