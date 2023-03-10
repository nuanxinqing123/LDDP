package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	"LDDP/utils/requests"
	res "LDDP/utils/response"
	"LDDP/utils/snowflake"
	"go.uber.org/zap"
	"strconv"
)

// OrderForwardSimple 简约查询
func OrderForwardSimple() (res.ResCode, []model.ForwardSimple) {
	var fs []model.ForwardSimple
	data := dao.GetOrderForwardData()
	for _, datum := range data {
		var fse model.ForwardSimple
		fse.ID = int(datum.ID)
		fse.ForwardName = datum.ForwardName
		fs = append(fs, fse)
	}
	return res.CodeSuccess, fs
}

// GetDivisionOrderForwardData 分页查询
func GetDivisionOrderForwardData(page string) (res.ResCode, model.OrderForwardPage) {
	var data []model.Forward
	var pPage model.OrderForwardPage
	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetDivisionOrderForwardData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetDivisionOrderForwardData(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetDivisionOrderForwardData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetOrderForwardDataPage()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pPage.Page = z + 1
	} else {
		pPage.Page = z
	}
	pPage.PageData = data

	return res.CodeSuccess, pPage

}

// GetSearchOrderForwardData 搜索
func GetSearchOrderForwardData(s string) (res.ResCode, []model.Forward) {
	return res.CodeSuccess, dao.GetSearchOrderForwardData(s)
}

// OrderForwardAdd 创建转发
func OrderForwardAdd(p *model.ForwardAdd) (res.ResCode, string) {
	// 检查转发名是否重复
	if dao.ForwardSearchOne(p.ForwardName).ForwardName != "" {
		return res.CodeProjectError, "转发名存在重复"
	}

	pr := &model.Forward{
		ForwardName:          p.ForwardName,
		ForwardAddress:       p.ForwardAddress,
		ForwardRefundAddress: p.ForwardRefundAddress,
	}

	if err := dao.ForwardAdd(pr); err != nil {
		return res.CodeProjectError, "创建转发失败"
	}

	return res.CodeSuccess, "创建转发成功"
}

// ForwardUpdate 修改转发
func ForwardUpdate(p *model.ForwardUpdate) (res.ResCode, string) {
	pr := dao.ForwardSearchByID(p.ID)
	pr.ForwardName = p.ForwardName
	pr.ForwardAddress = p.ForwardAddress
	pr.ForwardRefundAddress = p.ForwardRefundAddress

	if err := dao.ForwardUpdate(pr); err != nil {
		return res.CodeProjectError, "修改转发失败"
	}

	return res.CodeSuccess, "修改转发成功"
}

// ForwardDelete 删除转发
func ForwardDelete(p *model.ForwardDelete) (res.ResCode, string) {
	// 检查转发名是否重复
	if dao.ForwardSearchByID(p.ID).ID == 0 {
		return res.CodeProjectError, "转发不存在"
	}
	// 检查是否有项目绑定此转发
	if dao.ProjectSearchByForwardID(p.ID) {
		return res.CodeProjectError, "转发已被绑定，无法删除"
	}

	if err := dao.ForwardDelete(p.ID); err != nil {
		return res.CodeProjectError, "删除转发失败"
	}

	return res.CodeSuccess, "删除转发成功"
}

// OrderForwardApiTest 转发测试
func OrderForwardApiTest(uid any, p *model.ForwardTest) (res.ResCode, string) {
	vote := &model.Order{
		OrderTaskType:    "test",
		OrderID:          strconv.FormatInt(snowflake.GenID(), 10),
		OrderUID:         uid.(string),
		OrderTickets:     int64(2 * 10),
		OrderNumber:      10,
		OrderVariable:    "测试变量",
		OrderRemarks:     "测试订单转发",
		OrderState:       -1,
		OrderStateReason: "",
		OrderIP:          "127.0.0.1",
	}
	code, msg := SendOrderRequestTest(vote, p.ForwardAddress)
	if code == res.CodeProjectError {
		return code, msg
	}
	code, msg = SendOrderRefundRequestTest(SOrderRF{
		OrderID:  strconv.FormatInt(snowflake.GenID(), 10),
		OrderUID: uid.(string),
	}, p.ForwardRefundAddress)
	if code == res.CodeProjectError {
		return code, msg
	}
	return res.CodeSuccess, "转发测试成功"
}

// SendOrderRequestTest 发送订单请求
func SendOrderRequestTest(so *model.Order, url string) (res.ResCode, string) {
	sso, err := json.Marshal(so)
	if err != nil {
		zap.L().Error("订单转发请求失败：" + err.Error())
		return res.CodeProjectError, "订单转发请求失败，原因：" + err.Error()
	}
	// 发送请求
	result, err := requests.Requests("POST", url, string(sso), "")
	if err != nil {
		zap.L().Error("退单转发请求失败：" + err.Error())
		return res.CodeProjectError, "订单转发请求失败，原因：" + err.Error()
	} else {
		// 解析返回数据
		var sor Result
		err = json.Unmarshal(result, &sor)
		if err != nil {
			zap.L().Error("退单转发请求失败：" + err.Error())
			return res.CodeProjectError, "订单转发请求失败，原因：" + err.Error()
		} else {
			if sor.Code != 2000 {
				return res.CodeProjectError, "订单转发请求失败，原因：返回非2000代码"
			}
		}
	}
	return res.CodeSuccess, "订单转发请求成功"
}

// SendOrderRefundRequestTest 发送退单请求
func SendOrderRefundRequestTest(rf SOrderRF, url string) (res.ResCode, string) {
	zap.L().Debug("SendOrderRefundRequest 发送退单请求, 订单ID：" + rf.OrderID)
	sorf, err := json.Marshal(rf)
	if err != nil {
		zap.L().Error("退单转发请求失败：" + err.Error())
		return res.CodeProjectError, "退单转发请求失败，原因：" + err.Error()
	}
	// 发送请求
	result, err := requests.Requests("POST", url, string(sorf), "")
	if err != nil {
		zap.L().Error("退单转发请求失败：" + err.Error())
		return res.CodeProjectError, "退单转发请求失败，原因：" + err.Error()
	} else {
		// 解析返回数据
		var sor Result
		err = json.Unmarshal(result, &sor)
		if err != nil {
			zap.L().Error("退单转发请求失败：" + err.Error())
			return res.CodeProjectError, "退单转发请求失败，原因：" + err.Error()
		} else {
			if sor.Code != 2000 {
				return res.CodeProjectError, "退单转发请求失败，原因：返回非2000代码"
			}
		}
	}
	return res.CodeSuccess, "退单转发请求成功"
}
