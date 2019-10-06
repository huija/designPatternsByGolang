package main

import (
	"designPatternsByGolang/Template/Impl"
	"designPatternsByGolang/Template/template"
	"fmt"
)

func main() {
	c := Impl.NewCharDemo()
	c.Display()
	fmt.Println()
	s := Impl.NewStringDemo()
	s.Display()
	// 简单实现
	fmt.Println("下面我认为是golang实现模板模式的正确方式, 而上面则是为了实现而实现")
	c2 := &Impl.CharDemo{}
	template.Display(c2)
	s2 := Impl.StringDemo{}
	template.Display(s2)
}
