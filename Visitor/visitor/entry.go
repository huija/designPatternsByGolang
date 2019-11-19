package visitor

import (
	"errors"
	"github.com/gokangaroo/common/utils"
)

// entry2 统一抽象
type Entry interface {
	Element
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

func (E *EntrySuper) Add(entry Entry) error {
	return errors.New("file can not add file")
}

func (E *EntrySuper) ToString() string {
	return E.GetName() + "(" + utils.ToString(E.GetSize()) + ")"
}