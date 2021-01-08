package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Lllegué hasta aquí")
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(miCanal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000; i++ {
	}
	final := time.Now()
	miCanal <- final.Sub(inicio) //me da la duración que tengo que devolver en mi variable miCanal
}
