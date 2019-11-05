package items

import (
	"io"
	"os"
)

// S2H 从结构体转为html的抽象
type Item interface {
	MakeHtml() string
}

// Item 抽象零件, 标题
type ItemWrap struct {
	Item
	caption string
}

func NewItemWrap(caption string, item Item) ItemWrap {
	return ItemWrap{item, caption}
}

// Link 超链接
type Link struct {
	ItemWrap
	url string
}

func NewLink(caption, url string) Link {
	link := Link{url: url}
	link.ItemWrap = NewItemWrap(caption, link)
	return link
}

// Tray 超链接的外包装
type Tray struct {
	ItemWrap
	tray []Item
}

func NewTray(caption string) Tray {
	tray := Tray{tray: make([]Item, 0)}
	tray.ItemWrap = NewItemWrap(caption, tray)
	return tray
}

func (T Tray) Add(i Item) {
	T.tray = append(T.tray, i)
}

// Page 抽象产品, 最终的页面
type Page struct {
	ItemWrap
	title   string
	author  string
	content []Item
}

func NewPage(title, author string) Page {
	page := Page{title: title, author: author, content: make([]Item, 0)}
	page.ItemWrap = NewItemWrap("", page)
	return page
}

func (P Page) Add(i Item) {
	P.content = append(P.content, i)
}

func (P Page) Output() {
	fileName := P.title + ".html"
	if file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		_, _ = io.WriteString(file, P.MakeHtml())
	}
}
