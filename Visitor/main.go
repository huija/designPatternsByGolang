package main

import (
	"designPatternsByGolang/Visitor/visitor"
	"fmt"
)

// 访问者, 利用函数重载, 访问的类型不一样, 访问策略不一样
// 但是Golang不支持方法重载, 最后也只能是方法参数的多态的访问.
func main() {
	fmt.Println("golang不支持重载,只能限制在多态")
	// 这个模式还存在循环引用的问题, 所以, golang实在不适合这个模式, 切勿强行使用
	// TODO 这里我们复用原来的Composite代码, printList改为visitor
	fmt.Println("Making root entries...")
	rootdir := visitor.NewDir("root")
	bindir := visitor.NewDir("bin")
	tmpdir := visitor.NewDir("tmp")
	usrdir := visitor.NewDir("usr")
	rootdir.Add(bindir)                      // /root/bin
	rootdir.Add(tmpdir)                      // /root/tmp
	rootdir.Add(usrdir)                      // /root/usr
	bindir.Add(visitor.NewFile("vi", 10000)) // /root/bin/...
	bindir.Add(visitor.NewFile("latex", 20000))
	rootdir.Accept(&visitor.ListVisitor{})

	fmt.Println("")
	fmt.Println("Making user entries...")
	yuki := visitor.NewDir("yuki")
	hanako := visitor.NewDir("hanako")
	tomura := visitor.NewDir("tomura")
	usrdir.Add(yuki) // /root/usr/...
	usrdir.Add(hanako)
	usrdir.Add(tomura)
	yuki.Add(visitor.NewFile("diary.html", 100))
	yuki.Add(visitor.NewFile("Composite.java", 200))
	hanako.Add(visitor.NewFile("memo.tex", 300))
	tomura.Add(visitor.NewFile("game.doc", 400))
	tomura.Add(visitor.NewFile("junk.mail", 500))
	rootdir.Accept(&visitor.ListVisitor{})
}
