package main

//sistema operativo
//permite crear un paso a paso de lo que esta pasando

func ejemploPanic() {
	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1") //fuerza el error y aborta el programa
	}
}

func main() {
	ejemploPanic()
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo) //me devuelve un puntero al archivo(f)

	// defer fmt.Println("Llegamos al final del programa") //siempre se ejecuta antes de salir del programa donde esta

	// if err != nil {
	// 	fmt.Println("Error abriendo el archivo")
	// 	// os.Exit(1) si hay un exit no nos deja hacer el defer
	// 	//el número es para saber porque puerta salió
	// }

	// defer f.Close() //cierra el programa ya sea por error o por éxito
}
