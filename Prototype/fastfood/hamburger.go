package fastfood

import (
	"fmt"
)

type Hamburger struct {
	bread     int64
	meat      int64
	vegetable int64
}

func NewHamburger(bread, meat, vegetable int64) *Hamburger {
	return &Hamburger{bread, meat, vegetable}
}

func (h *Hamburger) Cost() {
	fmt.Println("======================")
	fmt.Println("你有新的美团订单,顾客点了一个汉堡包!")
	fmt.Printf("使用%d片面包\n", h.bread)
	fmt.Printf("使用%d片肉片\n", h.meat)
	fmt.Printf("使用%d片菜叶\n", h.vegetable)
	fmt.Println("铛铛铛铛,一个汉堡包完成啦!")
	fmt.Println("======================")
}

func (h *Hamburger) CreateClone() Product {
	return Clone(h)
}
