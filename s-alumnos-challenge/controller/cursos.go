package controller

import (
	"fmt"
)

//Estructura con los datos que tendrá cada curso
type Curso struct {
	nombre  string
	alumnos []Alumno
}

//Slice que alamcena los cursos
var CursoArray []Curso

//Función que agrega un nuevo curso
func AgregarCurso(nombre string) {
	//Primero valído que exista
	ifExist := ifExistCurso(nombre)

	if ifExist == true {
		fmt.Println("Ese curso ya existe")
	} else {
		//Agregamos curso al array
		CursoArray = append(CursoArray, Curso{
			nombre: nombre,
		})
	}
}

//Función que valída si el curso ya existe o no
func ifExistCurso(nombre string) bool {
	ifExist := false

	for i, _ := range CursoArray { //Recorro el array para chequear si existe o no el curso
		if CursoArray[i].nombre == nombre {
			ifExist = true
		}
	}
	return ifExist
}

//Función que agrega un alumno a un curso
func AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {
	ifExist := ifExistCurso(nombreCurso)

	ifExistCurso := false

	if ifExist == false {
		fmt.Println("El curso que desea agregar no existe")
	} else {
		ifExistAlumno := IfExistByDni(alumno.dni)

		if ifExistAlumno == true {
			for i, _ := range CursoArray {
				if CursoArray[i].nombre == nombreCurso {
					for j, _ := range CursoArray[i].alumnos {
						if CursoArray[i].alumnos[j].dni == alumno.dni {
							fmt.Println("El alumno", CursoArray[i].alumnos[j].nombre, "ya se encuentra cargado en el curso", CursoArray[i].nombre)
							ifExistCurso = true
						}
					}
					if ifExistCurso != true {
						CursoArray[i].alumnos = append(CursoArray[i].alumnos, Alumno{
							nombre:   alumno.nombre,
							apellido: alumno.apellido,
							dni:      alumno.dni,
						})
					}
				}
			}
		} else {
			fmt.Println("El alumno que desea agregar no se encuentra")
		}
	}
}
