package old

import "fmt"

type StringDemo struct {
}

func (S *StringDemo) Start() {
	fmt.Println("+-----+")
}
func (S *StringDemo) Content() {
	fmt.Println("greetings")
}
func (S *StringDemo) Ending() {
	fmt.Println("+-----+")
}
