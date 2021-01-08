package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int { //funcion publica que puedo acceder desde cualquier paquete porque esta con mayuscula
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 8 = %d\n", Calculo(5, 8))

	Calculo = func(num1 int, num2 int) int { //funcion anonima
		return num1 - num2
	}
	fmt.Printf("Resto 17 - 6 = %d\n", Calculo(17, 6))
}
