package main

import (
	"designPatternsByGolang/TemplateMethod/impl"
	"designPatternsByGolang/TemplateMethod/template"
	"fmt"
)

func main() {
	c := NewCharDemo()
	c.Display()
	fmt.Println()
	s := NewStringDemo()
	s.Display()
}

func NewStringDemo() *impl.StringDemo {
	s := &impl.StringDemo{}
	template2 := &template.Template{AbstractModel: s}
	s.Template = template2
	return s
}

func NewCharDemo() *impl.CharDemo {
	c := &impl.CharDemo{}
	template1 := &template.Template{AbstractModel: c}
	c.Template = template1
	return c
}
