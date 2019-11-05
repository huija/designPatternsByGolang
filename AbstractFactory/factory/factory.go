package factory

import "designPatternsByGolang/AbstractFactory/items"

type Factory interface {
	CreateLink(caption, url string) items.Item
	CreateTray(caption string) items.Item
	CreatePage(title, author string) items.Item
}
