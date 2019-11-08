package main

import (
	"designPatternsByGolang/AbstractFactory/factory"
)

func main() {
	t := "list_factory"
	var f factory.Factory
	switch t {
	case "list_factory":
		f = &factory.ListFactory{}
	}
	people := f.CreateLink("人民日报", "http://www.people.com.cn/")
	gmw := f.CreateLink("光明日报", "http://www.gmw.cn/")

	us_yahoo := f.CreateLink("Yahoo!", "http://www.yahoo.com/")
	jp_yahoo := f.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp/")
	excite := f.CreateLink("Excite", "http://www.excite.com/")
	google := f.CreateLink("Google", "http://www.google.com/")
	//	创建Tray
	traynews := f.CreateTray("日报")
	traynews.Add(people)
	traynews.Add(gmw)

	trayyahoo := f.CreateTray("Yahoo!")
	trayyahoo.Add(us_yahoo)
	trayyahoo.Add(jp_yahoo)

	traysearch := f.CreateTray("检索引擎")
	traysearch.Add(trayyahoo)
	traysearch.Add(excite)
	traysearch.Add(google)
	// 创建Pages
	page := f.CreatePage("LinkPage", "zhuanghj")
	page.Add(traynews)
	page.Add(traysearch)
	page.Output()
}
