package main

import "fmt"

var numero int

func main() {
	// numero = 1

	//switch case
	switch numero = 4; numero { //se puede asignar el valor en la misma linea
	case 1: //en el caso que numero valga 1
		fmt.Println("Vale 1")
	case 2: //en el caso que numero valga 2
		fmt.Println("Vale 2")
	default:
		fmt.Println("Vale más que 2")
	}

}
