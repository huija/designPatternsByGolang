package main

import (
	"designPatternsByGolang/Composite/dnf"
	"fmt"
)

// 组合模式, 就是文件夹与文件都看作文件: 内容与容器一致
func main() {
	fmt.Println("Making root entries...")
	rootdir := dnf.NewDir("root")
	bindir := dnf.NewDir("bin")
	tmpdir := dnf.NewDir("tmp")
	usrdir := dnf.NewDir("usr")
	rootdir.Add(bindir)                  // /root/bin
	rootdir.Add(tmpdir)                  // /root/tmp
	rootdir.Add(usrdir)                  // /root/usr
	bindir.Add(dnf.NewFile("vi", 10000)) // /root/bin/...
	bindir.Add(dnf.NewFile("latex", 20000))
	rootdir.PrintList()

	fmt.Println("")
	fmt.Println("Making user entries...")
	yuki := dnf.NewDir("yuki")
	hanako := dnf.NewDir("hanako")
	tomura := dnf.NewDir("tomura")
	usrdir.Add(yuki) // /root/usr/...
	usrdir.Add(hanako)
	usrdir.Add(tomura)
	yuki.Add(dnf.NewFile("diary.html", 100))
	yuki.Add(dnf.NewFile("Composite.java", 200))
	hanako.Add(dnf.NewFile("memo.tex", 300))
	tomura.Add(dnf.NewFile("game.doc", 400))
	tomura.Add(dnf.NewFile("junk.mail", 500))
	rootdir.PrintList()

}
