package main

import (
	"alumnosChallenge/controller"
	"fmt"
	/*
		Esto nos permit acceder a cualquier funcionalidad
		que tengamos pública en la controladora
		escrita en mayusculas
	*/)

func main() {

	//Así instanciamos un alumno
	alumno1 := controller.NewAlumno("Juan", "Perez", 35697890)
	alumno2 := controller.NewAlumno("Adrián", "García", 34567231)
	alumno3 := controller.NewAlumno("María", "Rodriguez", 36699861)
	alumno4 := controller.NewAlumno("Matías", "Gonzalez", 33695670)
	alumno5 := controller.NewAlumno("Matías", "Gonzalez", 33695670)

	//Mediante controller llamamos a la función que agrega un alumno y le pasamos el alumno instanciado
	controller.AgregarAlumno(alumno1)
	controller.AgregarAlumno(alumno2)
	controller.AgregarAlumno(alumno3)
	controller.AgregarAlumno(alumno4)
	controller.AgregarAlumno(alumno5)

	//Muestro por consola el slice de alumnos creado
	fmt.Println(controller.AlumnoArray)

	controller.EliminarAlumno(33695670)

	fmt.Println(controller.AlumnoArray)

	//Mediante el controller llamamos a la función que agrega un curso
	controller.AgregarCurso("Literatura")
	controller.AgregarCurso("Programación")
	controller.AgregarCurso("Matemática")
	controller.AgregarCurso("Matemática")

	//Muestro por consola el arreglo de cursos
	fmt.Println(controller.CursoArray)

	controller.AgregarAlumnoACurso(alumno1, "Literatura")

	fmt.Println(controller.CursoArray)
}
