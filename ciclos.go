package main

import "fmt"

/*
 * Ciclos y condicionales
 */

func main() {
	fmt.Println("Comenzando a trabajar con ciclos\n")

	// Go solo tiene implementado el ciclo for en su nucleo
	
	numero1 := 10
	numero2 := 20

	if numero1 > numero2 {
		fmt.Printf("%d es mayor que %d \n", numero1, numero2)
	} else if numero1 < numero2 {
		fmt.Printf("%d es mayor que %d \n", numero2, numero1)
	} else {
	    fmt.Println("Los numeros son iguales")
	}

	for i:=0; i<=10; i++{
	    fmt.Println(i)
	}
}
