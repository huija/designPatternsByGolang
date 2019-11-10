package function

import "designPatternsByGolang/Bridge/impl"

// CountDisplay 功能: 指定展示多少次
type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl impl.DisplayImpl) *CountDisplay {
	return &CountDisplay{NewDisplay(impl)}
}

// MultiDisplay 功能
func (CD *CountDisplay) MultiDisplay(times int) {
	CD.Open()
	for i := 0; i < times; i++ {
		CD.Print()
	}
	CD.Close()
}
