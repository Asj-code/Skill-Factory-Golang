package main

import "fmt"

func main() {
	fmt.Println(Calculo(34, 12))
	fmt.Println(Calculo(34))
}

//Calculo es un función que suma (se comenta para que no lo tome como un packete al compilar porque esta escrita en mayusculas)
func Calculo(numero ...int) int {

	total := 0
	for _, valor := range numero {
		total = total + valor
		// fmt.Printf("\n Posic [%d]: %d\n", posic, valor)
	}
	return total
}
