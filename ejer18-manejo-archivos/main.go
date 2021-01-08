package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// leoArchivo()
	// leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt") //con readfile no hace falta abrir y cerrar
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo)) //no usamos un buffer por que lo hace ioutil
		//convertimos a string un array de byte
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt") //hacemos un Open
	if err != nil {
		fmt.Println("Hubo un error") //manejo del error
	} else {
		scanner := bufio.NewScanner(archivo) // objeto de tipo scanner
		for scanner.Scan() {                 //hace un scanner linea por linea
			registro := scanner.Text()               //guardamos en una variable cada scanneo convertido a texto(string)
			fmt.Printf("Linea > " + registro + "\n") //imprimimos el registro linea por linea
		}
	}
	archivo.Close() //siempre que abrimos con Open hay que cerrar
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt") //crea un uevo archivo
	if err != nil {
		fmt.Println("Hubo un error") //manejo del error
		return                       //para que corte
	}
	fmt.Fprintln(archivo, "Esta es la linea nueva") //le pasamos el texto que queremos agregar al nuevo archivo
	archivo.Close()
}

func graboArchivo2() {
	nombreArchivo := "./nuevoArchivo.txt"
	if Append(nombreArchivo, "\nEsta es la segunda linea") == false { //agregamos el texto al archivo
		fmt.Println("error al agregar la segunda linea") //en caso de que sea false muestar el error
	}

}

//Append is a function
func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto) //devuelve dos valores
	if err != nil {
		fmt.Println("Hubo un error en el WriteString")
		return false
	}
	arch.Close()
	return true
}
