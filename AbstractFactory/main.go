package main

import (
	"designPatternsByGolang/AbstractFactory/factory"
	"designPatternsByGolang/AbstractFactory/items"
)

func main() {
	f := factory.ListFactory{}
	people := f.CreateLink("人民日报", "http://www.people.com.cn/")
	gmw := f.CreateLink("光明日报", "http://www.gmw.cn/")

	us_yahoo := f.CreateLink("Yahoo!", "http://www.yahoo.com/")
	jp_yahoo := f.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp/")
	excite := f.CreateLink("Excite", "http://www.excite.com/")
	google := f.CreateLink("Google", "http://www.google.com/")
	//	创建Tray
	traynews := f.CreateTray("日报")
	traynews.(items.Tray).Add(people)
	traynews.(items.Tray).Add(gmw)

	trayyahoo := f.CreateTray("Yahoo!")
	trayyahoo.(items.Tray).Add(us_yahoo)
	trayyahoo.(items.Tray).Add(jp_yahoo)

	traysearch := f.CreateTray("检索引擎")
	traysearch.(items.Tray).Add(trayyahoo)
	traysearch.(items.Tray).Add(excite)
	traysearch.(items.Tray).Add(google)
	// 创建Pages
	page := f.CreatePage("LinkPage", "zhuanghj")
	page.(items.Page).Add(traynews)
	page.(items.Page).Add(traysearch)
	page.(items.Page).Output()
}
