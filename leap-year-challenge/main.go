package main

import (
	"fmt"
	"math"
)

var year int

func main() {
	fmt.Scanf("%d", &year)     //Método que me permite tomar un valor por consola
	result := isLeapYear(year) //Guardo el resultado de la función en una variable
	showResult(result)
}

//Función que calcula si un año es biciesto o no
func isLeapYear(year int) bool {
	f := float64(year)
	if math.Mod(f, 4) == 0.0 && math.Mod(f, 100) != 0.0 {
		return true
	} else if math.Mod(f, 4) == 0.0 && math.Mod(f, 100) == 0.0 && math.Mod(f, 400) == 0.0 {
		return true
	}
	return false
}

//Función que muestra el resultado por consola
func showResult(result bool) {
	if result == true {
		fmt.Println("El año es biciesto")
	} else {
		fmt.Println("El año no es biciesto")
	}

}
