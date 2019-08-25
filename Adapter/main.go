package main

import (
	"designPatternsByGolang/Adapter/adapter"
	new2 "designPatternsByGolang/Adapter/new"
	"designPatternsByGolang/Adapter/old"
)

func main() {
	stringDemo := old.StringDemo{}
	stringDemo.Start()
	stringDemo.Content()
	stringDemo.Ending()
	var stringTargetDemo new2.StringTargetDemo = adapter.MyAdapter{}
	stringTargetDemo.StartTarget()
	stringTargetDemo.ContentTarget()
	stringTargetDemo.EndingTarget()
}
