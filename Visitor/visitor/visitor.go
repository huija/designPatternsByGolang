package visitor

// TODO 新增两个高层抽象
type Visitor interface {
	Visit(Element)
}

type Element interface {
	Accept(v Visitor)
}
