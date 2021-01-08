package main

import "fmt"

var num, numSecreto int

func main() {
	numSecreto = 9
	var primeraVez bool = true

	for { //repite todo lo que contiene hasta que se haga un break

		fmt.Println("Ingrese un numero: ") //pide por pantalla
		fmt.Scanln(&num)                   //lee lo que ingresaste y se le asigna a la variable declarada num el valor ingresado

		if num == numSecreto {
			fmt.Println("Adivinaste el número!")
			if primeraVez == true {
				primeraVez = false
				continue
			} else {
				break //corta la iteración
			}

		} else {
			fmt.Println("No adivinaste el numero!")
		}
	}

}

//también se puede hacer el clásico for

/*
for i:=0; i > 10; i++ {

}

*/
