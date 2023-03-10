package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"go.uber.org/zap"
	"strconv"
)

// ConsumptionLogData 消费记录：收费/退款：以20条数据分割
func ConsumptionLogData(page string) (res.ResCode, model.ConsumptionPageData) {
	var data []model.UserRecordsOfConsumption
	var conPage model.ConsumptionPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetUserRecordsOfConsumptionPageData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetUserRecordsOfConsumptionPageData(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetUserRecordsOfConsumptionPageData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetUserRecordsOfConsumptionData()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		conPage.Page = z + 1
	} else {
		conPage.Page = z
	}
	conPage.PageData = data

	return res.CodeSuccess, conPage
}

// ConsumptionSearch 消费记录：收费/退款：搜索
func ConsumptionSearch(c string) (res.ResCode, []model.UserRecordsOfConsumption) {
	return res.CodeSuccess, dao.ConsumptionSearch(c)
}

// UserConsumptionDivisionData 用户消费分页查询
func UserConsumptionDivisionData(uid any, page string) (res.ResCode, model.ConsumptionPageData) {
	var data []model.UserRecordsOfConsumption
	var conPage model.ConsumptionPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.UserConsumptionDivisionData(uid, 1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.UserConsumptionDivisionData(uid, 1)
		} else {
			// 查询指定页数的数据
			data = dao.UserConsumptionDivisionData(uid, intPage)
		}
	}

	// 查询总页数
	count := dao.UserConsumptionDivisionDataNum(uid)
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		conPage.Page = z + 1
	} else {
		conPage.Page = z
	}
	conPage.PageData = data

	return res.CodeSuccess, conPage
}
