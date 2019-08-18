package impl

import (
	"designPatternsByGolang/Iterators/entity"
	"designPatternsByGolang/Iterators/inter"
)

// BookShelf 包含很多书, 暴露迭代方法可以迭代到书籍
type BookShelf struct {
	index int //这个其实就是size了, 又或者叫endPoint
	books []entity.Book
}

// NewBookShelf 新建一个书架对象
func NewBookShelf(books []entity.Book) (bookShelf *BookShelf) {
	return &BookShelf{len(books), books}
}

// GetBook 获取指定下标的书籍 此index非彼index
func (bookShelf *BookShelf) GetBook(index int) entity.Book {
	return bookShelf.books[index]
}

// AppendBook 添加书籍
func (bookShelf *BookShelf) AppendBook(book entity.Book) {
	bookShelf.books = append(bookShelf.books, book)
	bookShelf.index++
}

// GetLength 获取书籍数量
func (bookShelf *BookShelf) GetLength() int {
	//return len(bookShelf.books)
	return bookShelf.index
}

// Iterator 返回当前书架的迭代器对象
func (bookShelf *BookShelf) Iterator() inter.Iterator {
	return NewBookShelfIterator(*bookShelf)
}
