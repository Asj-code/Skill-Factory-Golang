package main

import (
	"fmt"
	"time"

	us "github.com/766b/go-outliner/skill-factory-go/ejer14/user"
)

type alumno struct {
	us.Usuario
}

func main() {
	user := new(alumno)
	user.AltaUsuario(1, "Juan Martin", time.Now(), true)
	fmt.Println(user.Usuario)
}
