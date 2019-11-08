package factory

import "designPatternsByGolang/AbstractFactory/items"

type Factory interface {
	CreateLink(caption, url string) items.Link
	CreateTray(caption string) items.Tray
	CreatePage(title, author string) items.Page
}
