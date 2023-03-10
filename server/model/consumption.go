package model

import "gorm.io/gorm"

// UserRecordsOfConsumption 用户消费记录表
type UserRecordsOfConsumption struct {
	gorm.Model
	UserID      string `json:"user_id,omitempty"`      // 用户ID
	VoteOrder   string `json:"vote_order,omitempty"`   // 任务订单号
	TaskState   string `json:"task_state,omitempty"`   // 订单状态（消费、退回）
	VoteTickets int64  `json:"vote_tickets,omitempty"` // 操作点券
}

// ConsumptionPageData 消费数据
type ConsumptionPageData struct {
	Page     int64                      `json:"page"`
	PageData []UserRecordsOfConsumption `json:"page_data"`
}
