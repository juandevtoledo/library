package book

import "fmt"

type TexBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, autor string, pages int, editorial, level string) *TexBook {
	return &TexBook{
		Book:      Book{title, autor, pages},
		editorial: editorial,
		level:     level,
	}
}

func (t *TexBook) PrintInfo() {
	fmt.Printf("Title %s , autor %s , page %v, editorial %s , level %s \n",
		t.title, t.autor, t.pages, t.editorial, t.level)
}
