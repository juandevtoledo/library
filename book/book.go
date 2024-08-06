package book

import "fmt"

type Book struct {
	//Mayuscula publico o exportado, minuscula privado (Encapsulación)
	title string
	autor string
	pages int
}

// Constructor qué retorna un puntero
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title: title,
		autor: author,
		pages: pages,
	}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s , Autor %s , Pages %d \n", b.title, b.autor, b.pages)
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}
