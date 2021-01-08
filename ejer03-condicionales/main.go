package main

import "fmt"

var estado bool //false por defecto

func main() {
	// estado = true

	if estado = true; estado == true { //reasigna el valor en la misma linea que compara
		fmt.Println("Estado es true")
	} else if miBool = true; miBool == false {
		fmt.Println(miBool)
	} else {
		fmt.Println("Estado es false")
	}
}
