package factory

import "designPatternsByGolang/AbstractFactory/items"

type ListFactory struct {
}

func (L *ListFactory) CreateLink(caption, url string) items.Item {
	return items.NewListLink(caption, url)
}

func (L *ListFactory) CreateTray(caption string) items.Item {
	return items.NewListTray(caption)
}

func (L *ListFactory) CreatePage(title, author string) items.Item {
	return items.NewListPage(title, author)
}
