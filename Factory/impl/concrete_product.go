package impl

import "fmt"

type ConcreteProduct struct {
	ownerName string
}

func (C *ConcreteProduct) Destroy() {
	fmt.Println(C.ownerName + "的产品坏了,请维修人员进行维修")
}
