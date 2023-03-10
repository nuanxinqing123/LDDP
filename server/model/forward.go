package model

import "gorm.io/gorm"

type Forward struct {
	gorm.Model
	// 转发名称
	ForwardName string `json:"forward_name"`
	// 订单转发地址
	ForwardAddress string `json:"forward_address"`
	// 退单转发地址
	ForwardRefundAddress string `json:"forward_refund_address"`
}

type ForwardSimple struct {
	ID          int    `json:"id"`
	ForwardName string `json:"forward_name"`
}

// OrderForwardPage 分页数据
type OrderForwardPage struct {
	Page     int64     `json:"page"`      // 总页数
	PageData []Forward `json:"page_data"` // 分页查询数据
}

// ForwardAdd 添加数据
type ForwardAdd struct {
	ForwardName          string `json:"forward_name" binding:"required"`
	ForwardAddress       string `json:"forward_address" binding:"required"`
	ForwardRefundAddress string `json:"forward_refund_address" binding:"required"`
}

// ForwardUpdate 修改数据
type ForwardUpdate struct {
	ID                   int    `json:"id" binding:"required"`
	ForwardName          string `json:"forward_name" binding:"required"`
	ForwardAddress       string `json:"forward_address" binding:"required"`
	ForwardRefundAddress string `json:"forward_refund_address" binding:"required"`
}

// ForwardDelete 删除数据
type ForwardDelete struct {
	ID int `json:"id" binding:"required"`
}

// ForwardTest API数据
type ForwardTest struct {
	ForwardAddress       string `json:"forward_address" binding:"required"`
	ForwardRefundAddress string `json:"forward_refund_address" binding:"required"`
}
