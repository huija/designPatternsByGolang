package impl

import (
	"designPatternsByGolang/TemplateMethod/template"
	"fmt"
)

type StringDemo struct {
	*template.Template
}

func (S StringDemo) Start() {
	fmt.Println("+-----+")
}
func (S StringDemo) Content() {
	fmt.Println("greetings")
}
func (S StringDemo) Ending() {
	fmt.Println("+-----+")
}
