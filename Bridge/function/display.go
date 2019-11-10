package function

import "designPatternsByGolang/Bridge/impl"

// Display 功能层抽象, 桥梁
type Display struct {
	impl impl.DisplayImpl
}

func NewDisplay(impl impl.DisplayImpl) *Display {
	return &Display{impl}
}

func (D *Display) Open() {
	D.impl.RawOpen()
}

func (D *Display) Print() {
	D.impl.RawPrint()
}

func (D *Display) Close() {
	D.impl.RawClose()
}

func (D *Display) FinalDisplay() {
	D.Open()
	D.Print()
	D.Close()
}
