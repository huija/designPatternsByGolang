package singleton

import (
	"fmt"
	"sync"
)

// 私有化结构体, 无法直接获取, 只能通过构造器获取
type singleTonDemo1 struct {
	count int
}

// 饿汉
var singleTon1 = &singleTonDemo1{1}

func NewSingleTon1() *singleTonDemo1 {
	if singleTon1.count == 1 {
		fmt.Println("singleTonDemo1 created")
	}
	singleTon1.count++
	return singleTon1
}

type singleTonDemo2 struct {
	count int
}

// 懒汉,存在线程安全
var singleTon2 *singleTonDemo2

func NewSingleTon2() *singleTonDemo2 {
	if singleTon2 == nil {
		singleTon2 = &singleTonDemo2{1}
		fmt.Println("singleTonDemo2 created")
		return singleTon2
	}
	singleTon2.count++
	return singleTon2
}

type singleTonDemo3 struct {
	count int
}

// 双重锁,杜绝懒汉的线程安全问题的同时, 也只有第一次才会上锁
var singleTon3 *singleTonDemo3
var mutex sync.Mutex
var channel = make(chan int, 1)

func NewSingleTon3() *singleTonDemo3 {
	if singleTon3 == nil {
		channel <- 1
		//mutex.Lock()
		if singleTon3 == nil {
			singleTon3 = &singleTonDemo3{1}
			fmt.Println("singleTonDemo3 created")
		}
		<-channel
		//mutex.Unlock()
		return singleTon3
	}
	singleTon3.count++
	return singleTon3
}

type singleTonDemo4 struct {
	count int
}

// 使用sync.Once来实现, 与java使用内部类编译异曲同工
var singleTon4 *singleTonDemo4
var once sync.Once

func NewSingleTon4() *singleTonDemo4 {
	once.Do(func() {
		singleTon4 = &singleTonDemo4{1}
		fmt.Println("singleTonDemo4 created")
	})
	if singleTon4.count != 1 {
		singleTon4.count++
	}
	return singleTon4
}
