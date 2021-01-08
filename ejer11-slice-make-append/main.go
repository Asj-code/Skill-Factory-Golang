package main

import "fmt"

var miSlice []int

func main() {
	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{2, 5, 6, 8, 9} //array de 5 elementos
	porcion := elementos[1:]
	fmt.Println(elementos)
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 6, 20) //tipo de dato, longitud y capacidad
	fmt.Println(elementos)
	fmt.Printf("Largo: %d Capacidad: %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // append()agrega un elemento al slice y si su capacidad es exedida la extiende
	}
	fmt.Println(nums)
	fmt.Printf("Largo: %d Capacidad: %d\n", len(nums), cap(nums))
}

//la última posición de un arreglo es len(arreglo) - 1

//una matriz es un array que contiene arrays

/*
func main() {
	var matriz [3][2]int

	matriz[0][1] = 1   //dentro del primer array[0] el segundo valor[1] vale 1
}

*/
