package factory

type AbstractModel interface {
	CreateProduct(ownerName string) Product
}

type Factory struct {
	AbstractModel
}

func (F Factory) GetProduct(ownerName string) Product {
	return F.CreateProduct(ownerName)
}
