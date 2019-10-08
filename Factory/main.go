package main

import (
	"designPatternsByGolang/Factory/factory"
	"designPatternsByGolang/Factory/impl"
	"fmt"
)

func main() {
	concreteFactory := impl.NewFactory()
	product1 := concreteFactory.GetProduct("小明")
	product2 := concreteFactory.GetProduct("小红")
	fmt.Println("-------1年后-------")
	product1.Destroy()
	product2.Destroy()
	fmt.Println("我觉得下面的写法,更加清晰")
	concreteFactory = &impl.ConcreteFactory{}
	product1 = factory.GetProduct(concreteFactory, "王二")
	product2 = factory.GetProduct(concreteFactory, "麻子")
	fmt.Println("-------1年后-------")
	product1.Destroy()
	product2.Destroy()
}
