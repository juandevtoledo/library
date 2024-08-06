package animal

import (
	"fmt"
)

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " Hace gua guau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " Hace miau miau")
}

func HaceSonido(animal Animal) {
	animal.Sonido()
}
