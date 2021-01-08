package main

import (
	"fmt"
	"strconv"
)

var numero, numAuxiliar, numTotal int
var text string
var Estado bool //variable global que llamo en la función de abajo

func main() {
	unTexto := "Hola!"
	fmt.Println(unTexto)
	numero = 4
	fmt.Println(numero)
	var numDelMain int
	unTexto = "Chau!"
	numero2, numero3, otroTexto := 8, 63, "Como estas?"
	numAuxiliar = 3
	fmt.Println(numAuxiliar)
	numTotal = numero + numAuxiliar
	fmt.Println(numTotal)
	fmt.Println(text)

	fmt.Println(numDelMain)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(otroTexto)

	var moneda float32
	numero = int(moneda)           //Para convertir de float a entero
	unTexto = strconv.Itoa(numero) //para convertir a numero un string importando la libreria strconv
	fmt.Println(unTexto)

	estado := true
	fmt.Println(estado)
	mostrarEstado()
}

func mostrarEstado() {
	fmt.Println(Estado) //llamada de la variable global
}
