package fastfood

import (
	"fmt"
)

type MilkyTea struct {
	pearl int64
	milk  int64
	water int64
}

func NewMilkyTea(pearl, milk, water int64) *MilkyTea {
	return &MilkyTea{pearl, milk, water}
}

func (m MilkyTea) Cost() {
	fmt.Println("======================")
	fmt.Println("你有新的饿了么订单,顾客点了一杯奶茶!")
	fmt.Printf("使用%d克的珍珠\n", m.pearl)
	fmt.Printf("使用%dml的水\n", m.water)
	fmt.Printf("使用%dml的牛奶\n", m.milk)
	fmt.Println("铛铛铛铛,一杯奶茶完成啦!")
	fmt.Println("======================")
}

func (m MilkyTea) CreateClone() Product {
	return Clone(m)
}
