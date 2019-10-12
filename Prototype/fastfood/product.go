package fastfood

import (
	"reflect"
)

type Product interface {
	Cost()
	CreateClone() Product
}

// Clone 深复制
func Clone(p Product) Product {
	switch p.(type) {
	case Hamburger:
		v := reflect.ValueOf(p)
		return &Hamburger{v.FieldByName("bread").Int(), v.FieldByName("meat").Int(), v.FieldByName("vegetable").Int()}
	case MilkyTea:
		v := reflect.ValueOf(p)
		return &MilkyTea{v.FieldByName("pearl").Int(), v.FieldByName("milk").Int(), v.FieldByName("water").Int()}
	default:
		panic("switch type error")
	}
}
