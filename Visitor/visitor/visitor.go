package visitor

import (
	"fmt"
	"github.com/gokangaroo/common/utils"
)

type Element interface {
	Accept(v Visitor)
}

type Visitor interface {
	Visit(Element)
}

type ListVisitor struct {
	currentDir string //当前访问的文件夹名称
}

func (L *ListVisitor) Visit(E Element) {
	switch E.(type) {
	case *File:
		fmt.Println(L.currentDir + "/" + E.(*File).GetName() + "(" + utils.ToString(E.(*File).GetSize()) + ")")
	case *Dir:
		fmt.Println(L.currentDir + "/" + E.(*Dir).GetName() + "(" + utils.ToString(E.(*Dir).GetSize()) + ")")
		saveDir := L.currentDir
		L.currentDir = L.currentDir + "/" + E.(*Dir).GetName()
		entries := E.(*Dir).entries
		for _, value := range entries {
			// 这是一个递归
			value.Accept(L)
		}
		L.currentDir = saveDir
	}
}
