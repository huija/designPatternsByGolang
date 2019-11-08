package main

import (
	"designPatternsByGolang/Template/impl"
	"designPatternsByGolang/Template/template"
	"fmt"
)

func main() {
	c := impl.NewCharDemo()
	c.Display()
	fmt.Println()
	s := impl.NewStringDemo()
	s.Display()
	// 简单实现
	fmt.Println("我觉得下面的写法,更加清晰")
	c2 := &impl.CharDemo{}
	template.Display(c2)
	fmt.Println()
	s2 := &impl.StringDemo{}
	template.Display(s2)
}
