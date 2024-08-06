package book

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}
