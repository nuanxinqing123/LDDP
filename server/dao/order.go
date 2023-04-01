package dao

import "LDDP/server/model"

// GetDivisionOrderDataAll 条件查询Order数据
func GetDivisionOrderDataAll(page int) []model.Order {
	var v []model.Order
	if page == 1 {
		DB.Order("id desc").Limit(20).Offset(0).Find(&v)
	} else {
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&v)
	}

	return v
}

// GetOrderDataPage 获取Order表总数据
func GetOrderDataPage() int64 {
	var c []model.Order
	return DB.Find(&c).RowsAffected
}

// GetOrderData 查询搜索的订单
func GetOrderData(s string) []model.Order {
	var v []model.Order
	DB.Where("order_id LIKE ?", "%"+s+"%").Find(&v)
	return v
}

// GetOrderTaskData 任务变量搜索订单
func GetOrderTaskData(s string) []model.Order {
	var v []model.Order
	DB.Where("order_variable LIKE ?", "%"+s+"%").Find(&v)
	return v
}

// GetOrderTypeData 订单类名&状态搜索订单
func GetOrderTypeData(t string, s int) []model.Order {
	var v []model.Order
	DB.Where("order_task_type LIKE ? AND order_state = ?", "%"+t+"%", s).Find(&v)
	return v
}

// UserGetDivisionOrderDataAll 条件查询Order数据
func UserGetDivisionOrderDataAll(uid any, page int) []model.Order {
	var v []model.Order
	if page == 1 {
		DB.Where("order_uid = ?", uid).Order("id desc").Limit(100).Offset(0).Find(&v)
	} else {
		DB.Where("order_uid = ?", uid).Order("id desc").Limit(100).Offset((page - 1) * 100).Find(&v)
	}

	return v
}

// UserGetOrderDataPage 获取Order表总数据
func UserGetOrderDataPage(uid any) int64 {
	var c []model.Order
	return DB.Where("order_uid = ?", uid).Find(&c).RowsAffected
}

// CreateOrder 创建订单
func CreateOrder(p *model.Order) {
	DB.Create(&p)
}

// GetOneOrderData 查询订单
func GetOneOrderData(s string) model.Order {
	var v model.Order
	DB.Where("order_id = ?", s).Find(&v)
	return v
}

// GetOrderType 查询指定类型订单
func GetOrderType(t, s string) []model.Order {
	var o []model.Order
	DB.Where("order_task_type = ? AND order_state = ?", t, s).Find(&o)
	return o
}

// SaveOrder 保存订单
func SaveOrder(p model.Order) {
	DB.Save(&p)
}

// UserOrderRefund 订单退款
func UserOrderRefund(UID any, order string) {
	var o model.Order
	DB.Where("order_uid = ? AND order_id = ?", UID, order).First(&o)
	o.OrderState = 3
	DB.Save(&o)
}

// GetOrderCountData 获取指定日期的订单数量
func GetOrderCountData(s, e string) int64 {
	var r []model.Order
	return DB.Where("created_at between ? and ?", s, e).Find(&r).RowsAffected
}

// GetUserOrderCountData 获取用户指定日期的订单数量
func GetUserOrderCountData(uid any, s, e string) int64 {
	var r []model.Order
	return DB.Where("order_uid = ? AND created_at between ? and ?", uid, s, e).Find(&r).RowsAffected
}
