package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectType       string `json:"project_type"`              // 项目类型
	ProjectPrice      int    `json:"project_price"`             // 项目价格
	ProjectAgentPrice int    `json:"project_agent_price"`       // 项目代理价格
	ProjectAPI        int    `json:"project_api"`               // 项目API
	ProjectRemarks    string `json:"project_remarks,omitempty"` // 项目备注
	ProjectTips       string `json:"project_tips"`              // 项目提示弹窗
	ProjectAgentState bool   `json:"project_agent_state"`       // 项目代理状态（True：允许下单、False：暂停下单）
	ProjectState      bool   `json:"project_state"`             // 项目状态（True：允许下单、False：暂停下单）
}

// ProjectList 可用上限项目列表
type ProjectList struct {
	ProjectType    string `json:"project_type,omitempty"`  // 项目类型
	ProjectPrice   int    `json:"project_price,omitempty"` // 项目价格
	ProjectRemarks string `json:"project_remarks"`         // 项目备注
	ProjectTips    string `json:"project_tips"`            // 项目提示信息
}

// ProjectPage Project分页查询
type ProjectPage struct {
	Page     int64     `json:"page"`      // 总页数
	PageData []Project `json:"page_data"` // 分页查询数据
}

// ProjectAdd 创建项目
type ProjectAdd struct {
	ProjectType       string `json:"project_type"  binding:"required" `  // 项目类型
	ProjectPrice      int    `json:"project_price"  binding:"required" ` // 项目价格
	ProjectAgentPrice int    `json:"project_agent_price"`                // 项目代理价格
	ProjectRemarks    string `json:"project_remarks"`                    // 项目备注
}

// ProjectUpdate 修改项目
type ProjectUpdate struct {
	ID                int    `json:"id" binding:"required"`
	ProjectType       string `json:"project_type" binding:"required"`  // 项目类型
	ProjectPrice      int    `json:"project_price" binding:"required"` // 项目价格
	ProjectAgentPrice int    `json:"project_agent_price"`              // 项目代理价格
	ProjectAPI        int    `json:"project_api"`                      // 项目API
	ProjectRemarks    string `json:"project_remarks"`                  // 项目备注
	ProjectTips       string `json:"project_tips"`                     // 项目提示弹窗
	ProjectState      bool   `json:"project_state"`                    // 项目状态
	ProjectAgentState bool   `json:"project_agent_state"`              // 项目代理状态（True：允许下单、False：暂停下单）
}

// ProjectDelete 删除项目
type ProjectDelete struct {
	ID int `json:"id" binding:"required"`
}
