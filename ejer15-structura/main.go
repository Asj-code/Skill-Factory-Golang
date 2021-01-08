package main

import (
	"fmt"
	"time"
)

type usuario struct { //el equivalente  a clases en js
	ID        int
	Nombre    string
	FechaAlta time.Time
	status    bool
}

func main() {
	User := new(usuario)
	User.ID = 10
	User.Nombre = "Pablo"
	User.FechaAlta = time.Now()
	fmt.Println(User)
}
