package inter

// MyCollection 需要对外暴露获取迭代器对象的方法
type MyCollection interface {
	Iterator() Iterator
}
