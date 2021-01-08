package main

import "fmt"

func main() {
	potencia(2)
}

func potencia(numero int) {
	if numero < 100000 { //salida de la función
		return
	}
	fmt.Println(numero)
	potencia(numero * 2) //llamada recursiva
}
