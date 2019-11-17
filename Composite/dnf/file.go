package dnf

import (
	"designPatternsByGolang/Composite/entry"
	"fmt"
)

type File struct {
	*entry.EntrySuper
	name string
	size int
}

func NewFile(name string, size int) *File {
	file := &File{name: name, size: size}
	newEntry := entry.NewEntry(file)
	file.EntrySuper = newEntry
	return file
}

func (F *File) GetName() string {
	return F.name
}

func (F *File) GetSize() int {
	return F.size
}

func (F *File) PrintData(prefix string) {
	fmt.Println(prefix + "/" + F.ToString())
}
