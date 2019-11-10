package main

import (
	"designPatternsByGolang/Bridge/function"
	"designPatternsByGolang/Bridge/impl"
)

// 桥接模式, 主要思想是, 实现与功能层次分开, 解耦合.
// 功能层次作为调用者到实现层的一座桥(高架桥,看路标转弯)

func main() {
	d1 := function.NewDisplay(impl.NewStringDisplayImpl("Hello, China."))
	d2 := function.NewCountDisplay(impl.NewStringDisplayImpl("Hello, World."))
	d1.FinalDisplay()
	d2.FinalDisplay()
	d2.MultiDisplay(5)
}
