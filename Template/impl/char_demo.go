package impl

import (
	"designPatternsByGolang/Template/template"
	"fmt"
)

// CharDemo 字符实现, 需要反过来再包含一下(golang还是挺奇怪的)
type CharDemo struct {
	*template.Template
}

func (C *CharDemo) Start() {
	fmt.Print("<<")
}

func (C *CharDemo) Content() {
	fmt.Print("h")
}

func (C *CharDemo) Ending() {
	fmt.Print(">>")
}

// NewCharDemo 模仿构造器, 返回一个实现类
func NewCharDemo() *CharDemo {
	demo := &CharDemo{}
	newTemplate := template.NewTemplate(demo)
	demo.Template = newTemplate
	return demo
}
