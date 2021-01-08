package main

import "fmt"

/* Los MIDDLEWARES son INTERCEPTORES que permiten ejecutar
instrucciones comunes a varias funciones
que reciben y devuelven los mismos tupos de variables
*/

func main() {
	fmt.Println("Incio de programa")
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(8, 7)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(4, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		//Lo siguiente es lo que se aplica previo a realizar todas las funciones
		fmt.Println("Inicio de función")
		return f(a, b)
	}
}
