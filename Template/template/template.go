package template

// Template struct包着接口, 然后实现多态(接口无法直接作为接收器)
type Template struct {
	Implement
}

// NewTemplate struct来实现接收器, 定义一些固定流程, interface则实现多态
// 当然, 完全可以不这么写, 可以将Interface作为Display参数传入(而不是接收器)亦可, 更加好理解.
// 这种情况下模板模式就是一个简单的Interface实现了, 与其余语言一样, 通过接口定义即可实现.
func NewTemplate(implement Implement) *Template {
	return &Template{implement}
}

// Implement 抽象模板实现类需要实现的方法
type Implement interface {
	Start()
	Content()
	Ending()
}

func (T *Template) Display() {
	T.Start()
	for i := 0; i < 3; i++ {
		T.Content()
	}
	T.Ending()
}

// Display 简单实现
func Display(implement Implement) {
	implement.Start()
	for i := 0; i < 3; i++ {
		implement.Content()
	}
	implement.Ending()
}
