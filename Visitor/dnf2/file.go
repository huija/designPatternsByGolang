package dnf2

import (
	"designPatternsByGolang/Visitor/entry2"
	"designPatternsByGolang/Visitor/visitor"
)

type File struct {
	*entry2.EntrySuper
	name string
	size int
}

func NewFile(name string, size int) *File {
	file := &File{name: name, size: size}
	newEntry := entry2.NewEntry(file)
	file.EntrySuper = newEntry
	return file
}

func (F *File) GetName() string {
	return F.name
}

func (F *File) GetSize() int {
	return F.size
}

// TODO 访客模式, 授权访问
func (F *File) Accept(v visitor.Visitor) {
	v.Visit(F)
}
