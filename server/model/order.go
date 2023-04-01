package model

import "gorm.io/gorm"

// Order 订单记录
type Order struct {
	gorm.Model
	OrderTaskType    string `json:"order_task_type,omitempty"`    // 订单类型
	OrderID          string `json:"order_id,omitempty"`           // 订单号
	OrderUID         string `json:"order_uid,omitempty"`          // 订单归属用户
	OrderTickets     int64  `json:"order_tickets,omitempty"`      // 订单所需点券（非真实扣除积分，实际以真实扣除为准）
	OrderNumber      int    `json:"order_number,omitempty"`       // 订单任务数量
	OrderVariable    string `json:"order_variable,omitempty"`     // 任务变量
	OrderRemarks     string `json:"order_remarks,omitempty"`      // 订单备注
	OrderState       int    `json:"order_state"`                  // 订单状态（-1：等待中、0：进行中、1：已完成、2：终止、3：退款中、4：退款完成）
	OrderStateReason string `json:"order_state_reason,omitempty"` // 订单终止原因
	OrderStatus      string `json:"order_status"`                 // 订单实况
	OrderIP          string `json:"order_ip,omitempty"`           // 操作用户IP
}

// OrderPage Order分页查询
type OrderPage struct {
	Page     int64   `json:"page"`      // 总页数
	PageData []Order `json:"page_data"` // 分页查询数据
}

// OrderData 订单数据
type OrderData struct {
	// 今日订单数
	OrderToday int64 `json:"order_today"`
	// 7日订单数
	OrderSeven int64 `json:"order_seven"`
	// 30日订单数
	OrderThirty int64 `json:"order_thirty"`
	// 总订单数
	OrderTotal int64 `json:"order_total"`
}

// VoteOrder 任务订单
type VoteOrder struct {
	OrderTaskType string `json:"task_type" binding:"required"` // 任务类名
	OrderModel    string `json:"model" binding:"required"`     // 任务模型（单个：single、多个：multiple）
	OrderVariable string `json:"variableB" binding:"required"` // 任务变量
	OrderNumber   int    `json:"number"`                       // 任务数量
	OrderRemarks  string `json:"remarks"`                      // 订单备注
}

type OrderChart struct {
	SevenDaysOrder []int64  `json:"seven_days_order"` // 七日订单数量
	SevenDaysDate  []string `json:"seven_days_date"`
}
