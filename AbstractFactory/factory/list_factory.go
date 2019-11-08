package factory

import "designPatternsByGolang/AbstractFactory/items"

type ListFactory struct {
}

func (L *ListFactory) CreateLink(caption, url string) items.Link {
	return items.NewListLink(caption, url)
}

func (L *ListFactory) CreateTray(caption string) items.Tray {
	return items.NewListTray(caption)
}

func (L *ListFactory) CreatePage(title, author string) items.Page {
	return items.NewListPage(title, author)
}
