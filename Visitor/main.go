package main

import (
	"designPatternsByGolang/Visitor/dnf2"
	"designPatternsByGolang/Visitor/visitorImpl"
	"fmt"
)

// 访问者, 利用函数重载, 访问的类型不一样, 访问策略不一样
// 但是Golang不支持方法重载, 最后也只能是方法参数的多态的访问.
func main() {
	fmt.Println("golang不支持重载,只能限制在多态")
	// 这个模式还存在循环引用的问题, 所以, golang实在不适合这个模式, 切勿强行使用
	// TODO 这里我们复用原来的Composite代码, TODO只是为了高亮
	fmt.Println("Making root entries...")
	rootdir := dnf2.NewDir("root")
	bindir := dnf2.NewDir("bin")
	tmpdir := dnf2.NewDir("tmp")
	usrdir := dnf2.NewDir("usr")
	rootdir.Add(bindir)                   // /root/bin
	rootdir.Add(tmpdir)                   // /root/tmp
	rootdir.Add(usrdir)                   // /root/usr
	bindir.Add(dnf2.NewFile("vi", 10000)) // /root/bin/...
	bindir.Add(dnf2.NewFile("latex", 20000))
	rootdir.Accept(&visitorImpl.ListVisitor{})

	fmt.Println("")
	fmt.Println("Making user entries...")
	yuki := dnf2.NewDir("yuki")
	hanako := dnf2.NewDir("hanako")
	tomura := dnf2.NewDir("tomura")
	usrdir.Add(yuki) // /root/usr/...
	usrdir.Add(hanako)
	usrdir.Add(tomura)
	yuki.Add(dnf2.NewFile("diary.html", 100))
	yuki.Add(dnf2.NewFile("Composite.java", 200))
	hanako.Add(dnf2.NewFile("memo.tex", 300))
	tomura.Add(dnf2.NewFile("game.doc", 400))
	tomura.Add(dnf2.NewFile("junk.mail", 500))
	rootdir.Accept(&visitorImpl.ListVisitor{})
}
