package visitor

type Dir struct {
	*EntrySuper
	name    string
	entries []Entry
}

func NewDir(name string) *Dir {
	dir := &Dir{name: name}
	newEntry := NewEntry(dir)
	dir.EntrySuper = newEntry
	dir.entries = make([]Entry, 0)
	return dir
}

func (D *Dir) GetName() string {
	return D.name
}

func (D *Dir) GetSize() int {
	size := 0
	for _, value := range D.entries {
		size += value.GetSize()
	}
	return size
}

func (D *Dir) Add(entry Entry) {
	D.entries = append(D.entries, entry)
}

// TODO 访客模式, 授权访问
func (D *Dir) Accept(v Visitor) {
	v.Visit(D)
}
