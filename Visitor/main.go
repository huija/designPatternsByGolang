package main

import (
	"designPatternsByGolang/Visitor/visitor"
	"fmt"
)

// 访问者, 利用函数重载, 访问的类型不一样, 访问策略不一样
// 但是Golang不支持方法重载, 最后也只能是方法参数的多态的访问.
func main() {
	fmt.Println("golang不支持重载,只能限制在多态")
	c := &visitor.CustomerCol{}
	c.Add(visitor.NewEnterpriseCustomer("A company"))
	c.Add(visitor.NewEnterpriseCustomer("B company"))
	c.Add(visitor.NewIndividualCustomer("bob"))
	c.Accept(&visitor.ServiceRequestVisitor{})
	fmt.Println("==============================")
	c.Accept(&visitor.AnalysisVisitor{})
}
