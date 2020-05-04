package main

import (
	"designMode/decorator"
	"fmt"
)

func main() {
	noddle := &decorator.Ramen{
		Name: "noddle",
		Price: 10,
	}
	egg := &decorator.Egg{
		Noddle: noddle,
		Name: "egg",
		Price: 10,
	}
	suage := &decorator.Sausage{
		Noddle: egg,
		Name: "suage",
		Price: 10,
	}
	fmt.Println(suage.Description())
	fmt.Println(suage.GetPrice())
}
