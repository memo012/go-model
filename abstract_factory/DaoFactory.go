package abstract_factory

type OrderMainDao interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

// 抽象工厂接口
type DaoFactory interface {
	CreateOrderDetailDao() OrderDetailDao
	CreateOrderMainDao() OrderMainDao
}
