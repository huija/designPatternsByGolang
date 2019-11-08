package fastfood

import "reflect"

type Product interface {
	Cost()
	CreateClone() Product
}

// Clone 深度复制
// 简单实现,Product直接做参数实现多态, 没将Product放到struct内来写
func Clone(p Product) Product {
	switch p.(type) {
	case *Hamburger:
		v := reflect.ValueOf(*(p.(*Hamburger)))
		return &Hamburger{v.FieldByName("bread").Int(), v.FieldByName("meat").Int(), v.FieldByName("vegetable").Int()}
	case *MilkyTea:
		v := reflect.ValueOf(*(p.(*MilkyTea)))
		return &MilkyTea{v.FieldByName("pearl").Int(), v.FieldByName("milk").Int(), v.FieldByName("water").Int()}
	default:
		panic("switch type error")
	}
}
