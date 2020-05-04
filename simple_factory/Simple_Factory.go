package main

import "fmt"

type API interface {
	Say(name string) string
}

type Chinese struct {
}
func (c *Chinese) Say(name string) string {
	return "你好" + name
}

type English struct {
}
func (c *English) Say(name string) string {
	return "hello " + name
}

func NewApi(name string) API {
	if name == "zn" {
		return &Chinese{}
	}else if name == "en" {
		return &English{}
	}else {
		return nil
	}
}

func main() {
	api := NewApi("zn")
	name := api.Say("强子")
	fmt.Println(name)

	apiEn := NewApi("en")
	nameEn := apiEn.Say("memo012")
	fmt.Println(nameEn)
}
