package dao

import "LDDP/server/model"

// GetUserRecordsOfConsumptionData 获取UserRecordsOfConsumption消费数据总数
func GetUserRecordsOfConsumptionData() int64 {
	var c []model.UserRecordsOfConsumption
	result := DB.Find(&c)
	return result.RowsAffected
}

// GetUserRecordsOfConsumptionPageData 条件查询con数据
func GetUserRecordsOfConsumptionPageData(page int) []model.UserRecordsOfConsumption {
	var con []model.UserRecordsOfConsumption
	if page == 1 {
		DB.Order("id desc").Limit(20).Offset(0).Find(&con)
	} else {
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&con)
	}
	return con
}

// ConsumptionSearch 查询指定订单的消费记录
func ConsumptionSearch(order string) []model.UserRecordsOfConsumption {
	var c []model.UserRecordsOfConsumption
	DB.Where("vote_order LIKE ?", "%"+order+"%").Find(&c)
	return c
}

// CreateConsumption 创建订单消费记录
func CreateConsumption(p *model.UserRecordsOfConsumption) {
	DB.Create(&p)
}

// UserConsumptionDivisionDataNum 获取UserRecordsOfConsumption消费数据总数
func UserConsumptionDivisionDataNum(uid any) int64 {
	var c []model.UserRecordsOfConsumption
	result := DB.Where("user_id = ?", uid).Find(&c)
	return result.RowsAffected
}

// UserConsumptionDivisionData 条件查询con数据
func UserConsumptionDivisionData(uid any, page int) []model.UserRecordsOfConsumption {
	var con []model.UserRecordsOfConsumption
	if page == 1 {
		DB.Where("user_id = ?", uid).Order("id desc").Limit(20).Offset(0).Find(&con)
	} else {
		DB.Where("user_id = ?", uid).Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&con)
	}
	return con
}

// GetConsumptionToDayUpload 获取今日消费积分
func GetConsumptionToDayUpload(s, e string) int64 {
	var r []model.UserRecordsOfConsumption
	DB.Where("created_at between ? and ?", s, e).Find(&r)
	i := 0
	for _, re := range r {
		if re.TaskState == "消费" {
			i += int(re.VoteTickets)
		} else {
			i -= int(re.VoteTickets)
		}
	}
	return int64(i)
}
