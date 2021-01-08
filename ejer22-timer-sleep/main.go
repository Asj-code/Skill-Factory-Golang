package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLentooo("Luciana") //arranca pero sigue haciendo lo otro
	//con go adelante la hacemos (asincrona)
	fmt.Println("Estoy aquí")
	go miNombreLentooo("patricio")
	fmt.Println("Holaaaaaaaa")
	var x string
	fmt.Scanln(&x)

}

func miNombreLentooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
