package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//Inicializar Objeto (No funciona, por qué los metodos son privados)
	/*mybook := book.Book{
		Title: "Moby dick",
		Autor: "Herman Melville",
		Pages: 300,
	}*/

	//Inicializar por constructor custom
	mybook2 := book.NewBook("El corner", "gabo", 100)

	//mybook.PrintInfo()
	//mybook2.PrintInfo()
	mybook2.SetTitle("Cornel (Mdf)")
	fmt.Println(mybook2.GetTitle())

	myTextBook := book.NewTextBook("1922", "Charles dicken", 10, "Pinguin", "b1")
	myTextBook.PrintInfo()

	//myTextBook ya es una referencia, por qué el NewBook retorna una referencia
	book.Print(myTextBook)
	book.Print(mybook2)

	perro := animal.Perro{Nombre: "Dody"}
	gato := animal.Gato{Nombre: "Cat"}

	animal.HaceSonido(&perro)
	animal.HaceSonido(&gato)

	fmt.Println("***************************************")
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Peque"},
		&animal.Gato{Nombre: "Cat"},
		&animal.Gato{Nombre: "Catty"},
	}

	for _, value := range animales {
		value.Sonido()
	}

}
