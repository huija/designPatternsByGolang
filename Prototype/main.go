package main

import (
	"designPatternsByGolang/Prototype/fastfood"
	"fmt"
	"reflect"
)

func main() {
	manager := fastfood.NewManager()
	hamburger := fastfood.NewHamburger(2, 1, 2)
	milkyTea := fastfood.NewMilkyTea(50, 150, 350)
	manager.Register("hamburger", hamburger)
	manager.Register("milkyTea", milkyTea)
	// 快餐店开始营业
	p1 := manager.Create("hamburger") //点餐
	p2 := manager.Create("milkyTea")
	p1.Cost() //消耗材料制作p1
	p2.Cost()
	// 比较, 深度复制
	fmt.Println(&hamburger)
	fmt.Println(&p1)
	fmt.Println(reflect.DeepEqual(hamburger, p1))
}
