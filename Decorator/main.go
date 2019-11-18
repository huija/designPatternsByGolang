package main

import (
	"designPatternsByGolang/Decorator/decorator"
	"designPatternsByGolang/Decorator/display"
)

// Decorator 装饰物与被被装饰物的一致性
func main() {
	b := decorator.NewFullBorder(
		decorator.NewFullBorder(
			display.NewStringDisplay(
				"你好世界!")))
	b.Show()
}
