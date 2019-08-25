package adapter

import "designPatternsByGolang/Adapter/old"

type MyAdapter struct {
	old.StringDemo
}

func (M MyAdapter) StartTarget() {
	M.Start()
	M.Start()
}
func (M MyAdapter) ContentTarget() {
	M.Content()
	M.Content()
}
func (M MyAdapter) EndingTarget() {
	M.Ending()
	M.Ending()
}
