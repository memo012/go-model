package template

import "fmt"

type DocSuper struct {
	GetContent func() string
}

func (d *DocSuper) DoOperator()  {
	fmt.Println("对这个文档做了一些处理，文档是:", d.GetContent())
}
