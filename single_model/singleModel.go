package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once // 内核信号

// 懒汉式
func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func main() {
	i1 := GetSingleton()
	i2 := GetSingleton()
	if i1 == i2 {
		fmt.Println("相等")
	}else {
		fmt.Println("不等")
	}
}
