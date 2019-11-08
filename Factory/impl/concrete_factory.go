package impl

import (
	"designPatternsByGolang/Factory/factory"
	"fmt"
)

type ConcreteFactory struct {
	*factory.Factory
}

func (C *ConcreteFactory) CreateProduct(ownerName string) factory.Product {
	fmt.Println(ownerName + "来购买了此产品")
	return &ConcreteProduct{ownerName}
}

func NewFactory() *ConcreteFactory {
	f := &ConcreteFactory{}
	f2 := &factory.Factory{f}
	f.Factory = f2
	return f
}
