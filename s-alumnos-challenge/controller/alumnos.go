package controller

import (
	"fmt"
)

//Estructura con los datos que tendra cada alumno inicializado
type Alumno struct {
	nombre   string
	apellido string
	dni      int32
}

//Slice donde se guardarán los alumnos inicializados
var AlumnoArray []Alumno

//Función que crea un alumno
func NewAlumno(nombre string, apellido string, dni int32) Alumno {
	return Alumno{
		nombre:   nombre,
		apellido: apellido,
		dni:      dni,
	}
}

//Función para agregar un alumno al slice
func AgregarAlumno(a Alumno) {
	//Antes de hacer un append al slice hay que validar que el alumno no exista

	ifExist := IfExistByDni(a.dni) //Guardo el resultado de la función en una variable

	//Valído
	if ifExist == true {
		fmt.Println("Ese alumno ya se encuentra cargado")
	} else {
		AlumnoArray = append(AlumnoArray, Alumno{
			nombre:   a.nombre,
			apellido: a.apellido,
			dni:      a.dni,
		})
	}

}

//Función que valída si un alumno ya existe o no por su dni
func IfExistByDni(dni int32) bool {
	ifExist := false

	for i, _ := range AlumnoArray { //con range recorro el array
		if AlumnoArray[i].dni == dni {
			ifExist = true
		}
	}

	return ifExist //retorna si existe o no el alumno
}

//Función que elimina un alumno del slice por dni
func EliminarAlumno(dni int32) {
	ifExist := IfExistByDni(dni)

	//Valído
	if ifExist == true {
		for i := 0; i < len(AlumnoArray); i++ {
			if AlumnoArray[i].dni == dni {
				AlumnoArray = append(AlumnoArray[:i], AlumnoArray[i+1:]...) /*Recibe la posición donde está el
				alumno que quiero eliminar
				y mueve todos los elemento una posición para
				que no me quede el espacio vacío*/
			}
		}
	} else {
		fmt.Println("Ese alumno no ya se encuentra cargado en el slice")

	}

}
