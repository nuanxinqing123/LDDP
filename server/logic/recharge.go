package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"go.uber.org/zap"
	"strconv"
	"unicode"
)

// RechargeLogData 充值数据：以20条数据分割
func RechargeLogData(page string) (res.ResCode, model.RechargePage) {
	var data []model.Recharge
	var conPage model.RechargePage

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.UserRechargeInfoAll(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.UserRechargeInfoAll(1)
		} else {
			// 查询指定页数的数据
			data = dao.UserRechargeInfoAll(intPage)
		}
	}

	// 查询总页数
	count := dao.GetRechargeDataPageAll()
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

// RechargeSearch 充值数据：CDK搜索/UserID搜索
func RechargeSearch(s string) (res.ResCode, []model.Recharge) {
	// 判断是CDK搜索还是UserID搜索
	IsCDK := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			IsCDK = true
			break
		}
	}

	return res.CodeSuccess, dao.RechargeSearch(s, IsCDK)
}

// UserRechargeLogData 充值数据：以20条数据分割
func UserRechargeLogData(uid any, page string) (res.ResCode, model.RechargePage) {
	var data []model.Recharge
	var conPage model.RechargePage

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.UserUserRechargeInfoAll(uid, 1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.UserUserRechargeInfoAll(uid, 1)
		} else {
			// 查询指定页数的数据
			data = dao.UserUserRechargeInfoAll(uid, intPage)
		}
	}

	// 查询总页数
	count := dao.UserGetRechargeDataPageAll(uid)
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

// UserRechargeTickets 用户充值：用户点券充值
func UserRechargeTickets(uid any, p *model.UserRecharge, RemoteIP string) (res.ResCode, string) {
	// 校验CDK是否存在
	cdkData := dao.GetTicketsData(p.RechargeTickets)
	if cdkData.TicketsKey == "" {
		return res.CodeTicketsError, "请检查您的卡密是否有效"
	} else if cdkData.TicketsKeyState == false {
		return res.CodeTicketsError, "请检查您的卡密是否有效"
	}

	// 获取用户数据
	userData := dao.GetUserIDData(uid)
	// 充值用户额度
	userData.Points += int64(cdkData.TicketsKeyPoints)
	dao.UpdateUser(userData)

	// 已使用,禁用卡密
	cdkData.TicketsKeyState = false
	go dao.UpdateFalseTickets(cdkData)
	// 记录充值记录
	go dao.InsertUserRechargeLog(uid, int64(cdkData.TicketsKeyPoints), cdkData.TicketsKey, RemoteIP)
	return res.CodeSuccess, "充值成功"
}
