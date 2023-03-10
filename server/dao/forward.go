package dao

import (
	"LDDP/server/model"
)

// GetOrderForwardData 查询转发表
func GetOrderForwardData() []model.Forward {
	var v []model.Forward
	DB.Find(&v)
	return v
}

// GetDivisionOrderForwardData 分页查询
func GetDivisionOrderForwardData(page int) []model.Forward {
	var p []model.Forward
	if page == 1 {
		// 获取前20条数据
		DB.Order("id desc").Limit(20).Offset(0).Find(&p)
	} else {
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&p)
	}
	return p
}

// GetOrderForwardDataPage 查询转发表总数据
func GetOrderForwardDataPage() int64 {
	var c []model.Forward
	return DB.Find(&c).RowsAffected
}

// GetSearchOrderForwardData 搜索
func GetSearchOrderForwardData(s string) []model.Forward {
	var v []model.Forward
	DB.Where("forward_name LIKE ?", "%"+s+"%").Find(&v)
	return v
}

// ForwardSearchOne 搜索转发名是否存在
func ForwardSearchOne(s string) model.Forward {
	var v model.Forward
	DB.Where("forward_name = ?", s).First(&v)
	return v
}

// ForwardSearchByID ID搜索转发
func ForwardSearchByID(id int) model.Forward {
	var v model.Forward
	DB.Where("id = ?", id).First(&v)
	return v
}

// ForwardAdd 创建转发
func ForwardAdd(p *model.Forward) error {
	return DB.Create(&p).Error
}

// ForwardUpdate 修改转发
func ForwardUpdate(p model.Forward) error {
	return DB.Save(&p).Error
}

// ForwardDelete 删除转发
func ForwardDelete(id int) error {
	return DB.Delete(&model.Forward{}, id).Error
}

// GetForwardByID 根据ID查询转发
func GetForwardByID(id int) model.Forward {
	var p model.Forward
	DB.Where("id = ?", id).First(&p)
	return p
}
