package factory

type AbstractModel interface {
	CreateProduct(ownerName string) Product
}

type Factory struct {
	AbstractModel
}

func (F *Factory) GetProduct(ownerName string) Product {
	return F.CreateProduct(ownerName)
}

// GetProduct 没必要为了抽象类而抽象类
func GetProduct(A AbstractModel, ownerName string) Product {
	return A.CreateProduct(ownerName)
}
