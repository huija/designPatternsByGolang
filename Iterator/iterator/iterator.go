package iterator

// Iterator 迭代器接口, 实现者暴露的迭代方法
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
