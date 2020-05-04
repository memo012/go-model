package main

import "fmt"

type Subject interface {
	Do() string
}
type RealSubject struct {
}

func (*RealSubject) Do() string {
	return "real subject make proxy ..."
}

type ProxySubject struct {
	real  RealSubject
	money int
}

func (p *ProxySubject) Do() string {
	if p.money > 0 {
		return p.real.Do()
	} else {
		return "请充值..."
	}
}

func main() {
	proxyFactory := &ProxySubject{
		money: 62,
	}
	fmt.Println(proxyFactory.Do())
}
