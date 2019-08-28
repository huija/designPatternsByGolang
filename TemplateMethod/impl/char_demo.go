package impl

import (
	"designPatternsByGolang/TemplateMethod/template"
	"fmt"
)

type CharDemo struct {
	*template.Template
}

func (C CharDemo) Start() {
	fmt.Print("<<")
}

func (C CharDemo) Content() {
	fmt.Print("h")
}

func (C CharDemo) Ending() {
	fmt.Print(">>")
}
