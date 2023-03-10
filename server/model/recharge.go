package model

import "gorm.io/gorm"

// Recharge 用户充值记录
type Recharge struct {
	gorm.Model
	RechargeUID     string `json:"recharge_uid,omitempty"`     // 充值用户ID
	RechargeTickets string `json:"recharge_tickets,omitempty"` // 充值Tickets
	RechargePoints  int64  `json:"recharge_points,omitempty"`  // 充值点券
	RechargeIP      string `json:"recharge_ip,omitempty"`      // 操作用户IP
}

// UserRecharge 用户充值
type UserRecharge struct {
	RechargeTickets string `json:"recharge_tickets" binding:"required"`
}

// RechargePage 用户充值记录分页数据
type RechargePage struct {
	Page     int64      `json:"page"`      // 总页数
	PageData []Recharge `json:"page_data"` // 分页查询数据
}
