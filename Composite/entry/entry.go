package entry

import (
	"errors"
	"fmt"
)

// entry 统一抽象
type Entry interface {
	GetName() string
	GetSize() int
	PrintData(prefix string)
}

type EntrySuper struct {
	Entry
}

func NewEntry(entry Entry) *EntrySuper {
	return &EntrySuper{entry}
}

func (E *EntrySuper) PrintList() {
	E.PrintData("")
}

func (E *EntrySuper) Add(entry Entry) error {
	return errors.New("file can not add file")
}

func (E *EntrySuper) String() string {
	return E.GetName() + "(" + fmt.Sprintf("%d", E.GetSize()) + ")"
}
