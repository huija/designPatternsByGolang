package display

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}

type DisplaySuper struct {
	Display
}

func NewDisplaySuper(D Display) *DisplaySuper {
	return &DisplaySuper{D}
}

func (D *DisplaySuper) Show() {
	for i := 0; i < D.GetRows(); i++ {
		fmt.Println(D.GetRowText(i))
	}
}
