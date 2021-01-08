package main

func main() {

}

type animal interface {
	respirar()
	comer()
	especie() string
}

type perro struct {
	cantPatas  int
	respirando bool
	comiendo   bool
	ladrando   bool
}

type gato struct {
	cantPatas  int
	respirando bool
	comiendo   bool
	maullando  bool
}

func (this *perro) respirar() { this.respirando = true }
func (this *perro) comer()    { this.comiendo = true }
func (this *perro) especie()  { return "perro" }

func (this *gato) respirar()       { this.respirando = true }
func (this *gato) comer()          { this.comiendo = true }
func (this *gato) especie() string { return "gato" }
