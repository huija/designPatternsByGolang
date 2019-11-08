package fastfood

type Manager struct {
	showcase map[string]interface{}
}

func (m *Manager) Register(name string, product Product) {
	m.showcase[name] = product
}

func (m *Manager) Create(name string) Product {
	i := m.showcase[name].(Product)
	return i.CreateClone()
}

func NewManager() *Manager {
	i := make(map[string]interface{})
	manager := &Manager{i}
	return manager
}
