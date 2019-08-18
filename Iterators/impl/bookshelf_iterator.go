package impl

// BookShelfIterator 迭代器实现, 迭代书架
type BookShelfIterator struct {
	index     int // 迭代指针
	bookShelf BookShelf
}

// NewBookShelfIterator 返回指定书架的迭代器对象
func NewBookShelfIterator(bookShelf BookShelf) *BookShelfIterator {
	return &BookShelfIterator{0, bookShelf}
}

// HasNext 是否还有后续
func (bookShelfIterator *BookShelfIterator) HasNext() bool {
	if bookShelfIterator.index < bookShelfIterator.bookShelf.GetLength() {
		return true
	}
	return false
}

// Next 迭代下一个对象
func (bookShelfIterator *BookShelfIterator) Next() interface{} {
	bookShelfIterator.index++
	return bookShelfIterator.bookShelf.GetBook(bookShelfIterator.index - 1)
}
