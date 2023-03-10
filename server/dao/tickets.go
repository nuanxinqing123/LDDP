package dao

import (
	"LDDP/server/model"
)

// TicketsDivisionTicketsData Tickets分页查询
func TicketsDivisionTicketsData(page int) []model.Tickets {
	var tickets []model.Tickets
	if page == 1 {
		// 获取第一页数据（25条）
		DB.Order("id desc").Limit(20).Offset(0).Find(&tickets)
	} else {
		// 获取第N页数据
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&tickets)
	}
	return tickets
}

// GetTicketsDataPage 获取Tickets表总数
func GetTicketsDataPage() int64 {
	var c []model.Tickets
	result := DB.Find(&c)
	return result.RowsAffected
}

// TicketsRemarksSearch Tickets标识模糊搜索
func TicketsRemarksSearch(state, s string) []model.Tickets {
	var c []model.Tickets
	if state == "全部" {
		DB.Where("tickets_key_remarks LIKE ?", "%"+s+"%").Find(&c)
	} else if state == "未使用" {
		DB.Where("tickets_key_remarks LIKE ? and tickets_key_state = ?", "%"+s+"%", true).Find(&c)
	} else {
		DB.Where("tickets_key_remarks LIKE ? and tickets_key_state = ?", "%"+s+"%", false).Find(&c)
	}

	return c
}

// TicketsValueSearch Tickets值模糊搜索
func TicketsValueSearch(state, s string) []model.Tickets {
	var t []model.Tickets

	if state == "全部" {
		DB.Where("tickets_key LIKE ?", "%"+s+"%").Find(&t)
	} else if state == "未使用" {
		DB.Where("tickets_key LIKE ? and tickets_key_state = ?", "%"+s+"%", true).Find(&t)
	} else {
		DB.Where("tickets_key LIKE ? and tickets_key_state = ?", "%"+s+"%", false).Find(&t)
	}
	return t
}

// TicketsAdd 批量生成Tickets
func TicketsAdd(c *model.Tickets) {
	DB.Create(&c)
}

// TicketsDelete 删除model.Tickets数据
func TicketsDelete(p *model.DelTickets) error {
	return DB.Where("id = ?", p.ID).Delete(&model.Tickets{}).Error
}

// GetTicketsData 查询Tickets信息
func GetTicketsData(p string) model.Tickets {
	var c model.Tickets
	DB.Where("tickets_key = ?", p).First(&c)
	return c
}

// UpdateFalseTickets 禁用Tickets
func UpdateFalseTickets(p model.Tickets) {
	DB.Save(&p)
}

// GetOrderSevenData 获取七日卡密激活数量
func GetOrderSevenData(s, e string) int64 {
	var r []model.Tickets
	return DB.Where("updated_at between ? and ? and tickets_key_state = ?", s, e, false).Find(&r).RowsAffected
}
