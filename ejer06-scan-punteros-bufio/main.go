package main

import "fmt"

var num1, num2 int

func main() {
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&num1)
	fmt.Println("Ingrese otro numero: ")
	fmt.Scanln(&num2)
	fmt.Printf("La suma es: %d", num1+num2)
}

//tambien se puede hacer importando las librerias "bufio" y "os"

/*
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}else {
		fmr.Println("Hola ",nombre)
	}
}

*/
