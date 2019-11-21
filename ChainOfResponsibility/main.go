package main

import (
	"designPatternsByGolang/ChainOfResponsibility/support"
	"designPatternsByGolang/ChainOfResponsibility/trouble"
)

func main() {
	// 责任链模式, 挨家挨户要饭吃
	jishu := support.NewOddSupport("小明")
	li := support.NewLimitSupport("小红", 100)
	jishu.SetNext(li.Support)
	jishu.SupportTrouble(trouble.NewTrouble(11))
	jishu.SupportTrouble(trouble.NewTrouble(12))
	jishu.SupportTrouble(trouble.NewTrouble(102))
}
