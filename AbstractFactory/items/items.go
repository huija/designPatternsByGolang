package items

import (
	"io"
	"os"
)

// Item 零件抽象
type Item interface {
	MakeHtml() string
}

type Page interface {
	Add(item Item)
	Output()
	MakeHtml() string
}

type Tray interface {
	Add(item Item)
	MakeHtml() string
}

type Link interface {
	MakeHtml() string
}

// LinkImpl 超链接
type LinkImpl struct {
	Link
	caption string
	url     string
}

func NewLinkImpl(caption, url string, l Link) *LinkImpl {
	link := LinkImpl{caption: caption, url: url}
	link.Link = l
	return &link
}

// TrayImpl 超链接的外包装
type TrayImpl struct {
	Tray
	caption string
	tray    []Item
}

func NewTrayImpl(caption string, t Tray) *TrayImpl {
	tray := TrayImpl{caption: caption, tray: make([]Item, 0)}
	tray.Tray = t
	return &tray
}

func (T *TrayImpl) Add(i Item) {
	T.tray = append(T.tray, i)
}

// PageImpl 抽象产品, 最终的页面
type PageImpl struct {
	Page
	title   string
	author  string
	content []Item
}

func NewPageImpl(title, author string, pi Page) *PageImpl {
	page := PageImpl{title: title, author: author, content: make([]Item, 0)}
	page.Page = pi
	return &page
}

func (P *PageImpl) Add(i Item) {
	P.content = append(P.content, i)
}

func (P *PageImpl) Output() {
	fileName := P.title + ".html"
	if file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		_, _ = io.WriteString(file, P.MakeHtml())
	}
}
