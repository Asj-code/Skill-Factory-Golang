package main

import "fmt"

func main() {
	asistencia() // Esta función me muestra los alumnos presentes y ausentes de la clase
}

func asistencia() {
	alumnos := map[string]bool{
		"Juan":  true,
		"Jose":  false,
		"Pedro": true,
		"Maria": true,
	}

	for alumno, asistencia := range alumnos {
		fmt.Printf("El alumno %s asistió: %t\n", alumno, asistencia)
	}
	asistencia, existe := alumnos["Benjamín"]
	fmt.Printf("El alumno %s y existe: %t\n", asistencia, existe)
}
