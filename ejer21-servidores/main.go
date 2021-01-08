package main

import (
	"log"
	"os"
)

func ejemploPanic() {

	f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	for i := 0; i <= 10; i++ {
		log.Printf("Error línea %v", i)
	}

	// defer func() {
	// 	reco := recover()
	// 	if reco != nil {
	// 		log.Fatalf("Ocurrió un error que generó panic \n%v", reco)
	// 	}
	// }()

	// a := 1
	// if a == 1 {
	// 	panic("Se encontró el valor 1")
	// }
}

func main() {
	ejemploPanic()

}
