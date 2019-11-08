package main

import (
	"designPatternsByGolang/Iterator/entity"
	"designPatternsByGolang/Iterator/impl"
	"fmt"
)

func main() {
	books := make([]*entity.Book, 0)
	bookShelf := impl.NewBookShelf(books)
	bookShelf.AppendBook(entity.NewBook("十万个为什么"))
	bookShelf.AppendBook(entity.NewBook("西游记"))
	bookShelf.AppendBook(entity.NewBook("红楼梦"))
	iterator := bookShelf.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next().(*entity.Book).GetName())
	}
}
