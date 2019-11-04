package main

import (
	"designPatternsByGolang/Builder/builder"
	"designPatternsByGolang/Builder/director"
	"fmt"
)

func main() {
	htmlBuilder := &builder.HTMLBuilder{}
	director := &director.Director{Builder: htmlBuilder}
	director.Construct()
	filename := htmlBuilder.GetFileName()
	fmt.Println(filename + "文件编写完成,存放在项目根目录.")
}
