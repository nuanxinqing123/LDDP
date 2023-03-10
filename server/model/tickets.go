package model

import "gorm.io/gorm"

type Tickets struct {
	gorm.Model
	TicketsKey        string `json:"tickets_key,omitempty"`         // Tickets值
	TicketsKeyPoints  int    `json:"tickets_key_points,omitempty"`  // Tickets数
	TicketsKeyRemarks string `json:"tickets_key_remarks,omitempty"` // Tickets备注
	TicketsKeyState   bool   `json:"tickets_key_state,omitempty"`   // Tickets状态（true：启用、false：禁用）
}

// TicketsPageData Tickets分页数据
type TicketsPageData struct {
	Page     int64     `json:"page"`
	PageData []Tickets `json:"page_data"`
}

// CreateTickets 生成Tickets
type CreateTickets struct {
	TicketsKeyCount   int    `json:"tickets_key_count" binding:"required"`  // CD-KEY生成数量
	TicketsKeyPoints  int    `json:"tickets_key_points" binding:"required"` // CD-KEY积分
	TicketsKeyPrefix  string `json:"tickets_key_prefix"`                    // CD-KEY前缀
	TicketsKeyRemarks string `json:"tickets_key_remarks"`                   // CD-KEY标识
}

// DelTickets 删除CDK数据
type DelTickets struct {
	ID int `json:"id" binding:"required"` // Tickets ID
}
