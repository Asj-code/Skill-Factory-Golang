package main

import "fmt"

func main() {

	var stringInput string //Guardo el valor en una variable

	fmt.Scanf("%s", &stringInput) //Método que me permite tomar un valor por consola

	// Devuelvo el string invertido y lo muestro por consola
	stringReversed := reverseString(stringInput)
	fmt.Println(stringInput)
	fmt.Println(stringReversed)
}

//Función que invierte el orden de los caracteres de un string
func reverseString(stringInput string) (result string) {
	for _, value := range stringInput {
		result = string(value) + result
	}
	return
}
