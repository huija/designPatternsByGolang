package dnf

import (
	"designPatternsByGolang/Composite/entry"
	"fmt"
	"github.com/gokangaroo/common/utils"
)

type Dir struct {
	*entry.EntrySuper
	name    string
	entries []entry.Entry
}

func NewDir(name string) *Dir {
	dir := &Dir{name: name}
	newEntry := entry.NewEntry(dir)
	dir.EntrySuper = newEntry
	dir.entries = make([]entry.Entry, 0)
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

func (D *Dir) Add(entry entry.Entry) {
	D.entries = append(D.entries, entry)
}

func (D *Dir) PrintData(prefix string) {
	fmt.Println(prefix + "/" + utils.ToString(D))
	for _, value := range D.entries {
		value.PrintData(prefix + "/" + value.GetName())
	}
}
