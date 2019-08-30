package main

import (
	"designPatternsByGolang/Factory/factory"
	"designPatternsByGolang/Factory/impl"
	"fmt"
)

func main() {
	concreteFactory := NewFactory()
	product1 := concreteFactory.GetProduct("小明")
	product2 := concreteFactory.GetProduct("小红")
	fmt.Println("-------1年后-------")
	product1.Destroy()
	product2.Destroy()
}

func NewFactory() *impl.ConcreteFactory {
	f := &impl.ConcreteFactory{}
	f2 := &factory.Factory{f}
	f.Factory = f2
	return f
}
