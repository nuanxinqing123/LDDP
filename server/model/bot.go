package model

import "gorm.io/gorm"

type BotConfig struct {
	// 机器人配置
	gorm.Model
	PassWord string `json:"pass_word" binding:"required"` // 机器人密码
	AuthIP   string `json:"auth_ip" binding:"required"`   // 机器人IP
}

// DeductionRefund 扣除退款
type DeductionRefund struct {
	OrderID string `json:"order_id" binding:"required"` // 订单号
	Type    int    `json:"type" binding:"required"`     // 操作类型（1：扣除、2：退款）
	Points  int    `json:"points" binding:"required"`   // 点券数量
}

// BotOrderUpdate 修改订单状态
type BotOrderUpdate struct {
	OrderID          string `json:"order_id" binding:"required"` // 订单号
	OrderState       int    `json:"order_state"`                 // 订单状态
	OrderStateReason string `json:"order_state_reason"`          // 订单终止原因
	OrderStatus      string `json:"order_status"`                // 订单实况
}
