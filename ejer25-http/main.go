package main

import "net/http"

func main() {
	http.HandleFunc("/", home) //manejador de funciones
	//lo primero es la ruta(/), nombre de la funcion (response)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":3000", nil) //escucha y devuelve la pagina que pido
	//primer parametro es el puerto normal
	//el segundo parametro sera un manejador que aca no lo necesito(nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html") //sive el archivo html(voy a tener solicitud(w) y respuesta(r))
	//y la ruta de que es lo que queremos que abra

}
