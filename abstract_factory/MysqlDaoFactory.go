package abstract_factory

import "fmt"

type MysqlMainDao struct {
}

func (*MysqlMainDao) SaveOrderMain()  {
	fmt.Println("mysqlMainDao run ...")
}

type MysqlDetailDao struct {
}

func (*MysqlDetailDao) SaveOrderDetail()  {
	fmt.Println("mysqlMainDao run ...")
}


