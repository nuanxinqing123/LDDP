package dao

import "LDDP/server/model"

// GetRechargeDataPageAll 用户充值：全部数据量
func GetRechargeDataPageAll() int64 {
	var c []model.Recharge
	result := DB.Find(&c)
	return result.RowsAffected
}

// UserRechargeInfoAll 用户充值：全部记录查询
func UserRechargeInfoAll(page int) []model.Recharge {
	var recharge []model.Recharge

	if page == 1 {
		DB.Order("id desc").Limit(20).Offset(0).Find(&recharge)
	} else {
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&recharge)
	}

	return recharge
}

// RechargeSearch 用户充值：记录查询搜索
func RechargeSearch(s string, is bool) []model.Recharge {
	var r []model.Recharge
	if is {
		// CDK搜索
		DB.Where("recharge_tickets LIKE ?", "%"+s+"%").Find(&r)
	} else {
		// UserID搜索
		DB.Where("recharge_uid LIKE ?", "%"+s+"%").Find(&r)
	}

	return r
}

// UserUserRechargeInfoAll 用户充值：全部记录查询
func UserUserRechargeInfoAll(uid any, page int) []model.Recharge {
	var recharge []model.Recharge

	if page == 1 {
		DB.Where("recharge_uid = ?", uid).Order("id desc").Limit(20).Offset(0).Find(&recharge)
	} else {
		DB.Where("recharge_uid = ?", uid).Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&recharge)
	}

	return recharge
}

// UserGetRechargeDataPageAll 用户充值：全部数据量
func UserGetRechargeDataPageAll(uid any) int64 {
	var c []model.Recharge
	result := DB.Where("recharge_uid = ?", uid).Find(&c)
	return result.RowsAffected
}

// InsertUserRechargeLog 记录用户充值记录
func InsertUserRechargeLog(uid any, points int64, tickets, RemoteIP string) {
	var rechargeLog model.Recharge
	rechargeLog.RechargeUID = uid.(string)
	rechargeLog.RechargePoints = points
	rechargeLog.RechargeTickets = tickets
	rechargeLog.RechargeIP = RemoteIP
	DB.Create(&rechargeLog)
}

// GetRechargeToDayUpload 查询今日充值总额
func GetRechargeToDayUpload(s string) int64 {
	var c []model.Recharge
	DB.Where("created_at >= ?", s).Find(&c)
	total := int64(0)
	for _, v := range c {
		total += v.RechargePoints
	}
	return total
}
