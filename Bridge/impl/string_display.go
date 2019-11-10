package impl

import (
	"fmt"
	"unicode/utf8"
)

type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{str: str, width: utf8.RuneCountInString(str)}
}

func (S *StringDisplayImpl) RawOpen() {
	S.printLine()
}
func (S *StringDisplayImpl) RawPrint() {
	fmt.Println("|" + S.str + "|")
}
func (S *StringDisplayImpl) RawClose() {
	S.printLine()
}

func (S *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < S.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
