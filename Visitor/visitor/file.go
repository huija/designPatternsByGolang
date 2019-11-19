package visitor

type File struct {
	*EntrySuper
	name string
	size int
}

func NewFile(name string, size int) *File {
	file := &File{name: name, size: size}
	newEntry := NewEntry(file)
	file.EntrySuper = newEntry
	return file
}

func (F *File) GetName() string {
	return F.name
}

func (F *File) GetSize() int {
	return F.size
}

// TODO 访客模式, 授权访问
func (F *File) Accept(v Visitor) {
	v.Visit(F)
}
