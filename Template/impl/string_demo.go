package impl

import (
	"designPatternsByGolang/Template/template"
	"fmt"
)

// StringDemo 字符实现
type StringDemo struct {
	*template.Template
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

// NewStringDemo 模仿构造器, 返回一个实现类
func NewStringDemo() *StringDemo {
	demo := &StringDemo{}
	newTemplate := template.NewTemplate(demo)
	demo.Template = newTemplate
	return demo
}
