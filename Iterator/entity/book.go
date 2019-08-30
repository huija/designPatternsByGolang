package entity

// Book 数据, 被迭代的对象(可以是任何类型)
type Book struct {
	name string
}

// NewBook 模拟构造函数
func NewBook(name string) (book Book) {
	return Book{name: name}
}

// GetName 获取书籍名称, 这个与迭代器模式无关, 是这个单元自身的方法
func (book Book) GetName() string {
	return book.name
}
