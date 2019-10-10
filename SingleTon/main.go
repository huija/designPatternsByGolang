package main

import (
	"designPatternsByGolang/SingleTon/singleton"
	"fmt"
	"reflect"
)

func main() {
	s1 := singleton.NewSingleTon1()
	s2 := singleton.NewSingleTon1()
	fmt.Println(reflect.DeepEqual(s1, s2))

	s3 := singleton.NewSingleTon2()
	s4 := singleton.NewSingleTon2()
	fmt.Println(reflect.DeepEqual(s3, s4))

	s5 := singleton.NewSingleTon3()
	s6 := singleton.NewSingleTon3()
	fmt.Println(reflect.DeepEqual(s5, s6))

	s7 := singleton.NewSingleTon4()
	s8 := singleton.NewSingleTon4()
	fmt.Println(reflect.DeepEqual(s7, s8))
}
