package template

type AbstractModel interface {
	Start()
	Content()
	Ending()
}

type Template struct {
	AbstractModel
}

func (T Template) Display() {
	T.Start()
	for i := 0; i < 3; i++ {
		T.Content()
	}
	T.Ending()
}
