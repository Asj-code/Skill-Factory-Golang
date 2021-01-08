package main

import "fmt"

func main() {
	fmt.Println(uno(8))
	num, estado := dos(9)
	fmt.Println(num)
	fmt.Println(estado)
}

func uno(numero int) int { //el tipo de dato que devuelve la funcion
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 5 {
		return 5, true
	} else {
		return 10, false
	}
}
