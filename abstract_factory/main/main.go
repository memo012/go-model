package main

import "designMode/abstract_factory"

func main() {
	app := &abstract_factory.MysqlFactory{}
	detailDao := app.CreateOrderDetailDao()
	detailDao.SaveOrderDetail()
	mainDao := app.CreateOrderMainDao()
	mainDao.SaveOrderMain()
}