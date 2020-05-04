package abstract_factory

type MysqlFactory struct {
}

func (*MysqlFactory) CreateOrderDetailDao() OrderDetailDao {
	return &MysqlDetailDao{}
}

func (*MysqlFactory) CreateOrderMainDao() OrderMainDao {
	return &MysqlMainDao{}
}