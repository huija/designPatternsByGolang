package dnf2

import (
	"designPatternsByGolang/Visitor/entry2"
	"designPatternsByGolang/Visitor/visitor"
)

type Dir struct {
	*entry2.EntrySuper
	name    string
	entries []entry2.Entry
}

func NewDir(name string) *Dir {
	dir := &Dir{name: name}
	newEntry := entry2.NewEntry(dir)
	dir.EntrySuper = newEntry
	dir.entries = make([]entry2.Entry, 0)
	return dir
}

func (D *Dir) GetName() string {
	return D.name
}

func (D *Dir) GetSize() int {
	size := 0
	for _, value := range D.entries {
		size += value.GetSize()
	}
	return size
}

func (D *Dir) Add(entry entry2.Entry) {
	D.entries = append(D.entries, entry)
}

// TODO 访客模式, 授权访问
func (D *Dir) Accept(v visitor.Visitor) {
	v.Visit(D)
}

func (D *Dir) GetEntries() []entry2.Entry {
	return D.entries
}
