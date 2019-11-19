package visitorImpl

import (
	"designPatternsByGolang/Visitor/dnf2"
	"designPatternsByGolang/Visitor/visitor"
	"fmt"
	"github.com/gokangaroo/common/utils"
)

type ListVisitor struct {
	currentDir string //当前访问的文件夹名称
}

func (L *ListVisitor) Visit(E visitor.Element) {
	switch E.(type) {
	case *dnf2.File:
		fmt.Println(L.currentDir + "/" + E.(*dnf2.File).GetName() + "(" + utils.ToString(E.(*dnf2.File).GetSize()) + ")")
	case *dnf2.Dir:
		fmt.Println(L.currentDir + "/" + E.(*dnf2.Dir).GetName() + "(" + utils.ToString(E.(*dnf2.Dir).GetSize()) + ")")
		saveDir := L.currentDir
		L.currentDir = L.currentDir + "/" + E.(*dnf2.Dir).GetName()
		entries := E.(*dnf2.Dir).GetEntries()
		for _, value := range entries {
			// 这是一个递归
			value.Accept(L)
		}
		L.currentDir = saveDir
	}
}
