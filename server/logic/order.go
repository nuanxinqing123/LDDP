package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	"LDDP/utils/ip"
	"LDDP/utils/requests"
	res "LDDP/utils/response"
	"LDDP/utils/snowflake"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type SOrderRF struct {
	OrderID  string `json:"order_id,omitempty"`  // 订单号
	OrderUID string `json:"order_uid,omitempty"` // 订单归属用户
}

type Result struct {
	Code int `json:"code"`
}

// GetDivisionOrderData 管理员订单分页查询
func GetDivisionOrderData(page string) (res.ResCode, model.OrderPage) {
	var data []model.Order
	var votePage model.OrderPage
	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetDivisionOrderDataAll(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetDivisionOrderDataAll(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetDivisionOrderDataAll(intPage)
		}
	}

	// 查询总页数
	count := dao.GetOrderDataPage()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		votePage.Page = z + 1
	} else {
		votePage.Page = z
	}

	votePage.PageData = data

	return res.CodeSuccess, votePage
}

// UserOrderData 订单数据
func UserOrderData(uid any) (res.ResCode, model.OrderData) {
	return res.CodeSuccess, model.OrderData{
		OrderToday:  dao.GetUserOrderCountData(uid, time.Now().Format("2006-01-02")+" 00:00:00", time.Now().Format("2006-01-02")+" 23:59:59"),
		OrderSeven:  dao.GetUserOrderCountData(uid, time.Now().Format("2006-01-02")+" 00:00:00", time.Now().AddDate(0, 0, -7).Format("2006-01-02")+" 23:59:59"),
		OrderThirty: dao.GetUserOrderCountData(uid, time.Now().Format("2006-01-02")+" 00:00:00", time.Now().AddDate(0, 0, -30).Format("2006-01-02")+" 23:59:59"),
		OrderTotal:  dao.UserGetOrderDataPage(uid),
	}
}

// UserGetDivisionOrderData 订单分页查询
func UserGetDivisionOrderData(uid any, page string) (res.ResCode, model.OrderPage) {
	var data []model.Order
	var votePage model.OrderPage
	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.UserGetDivisionOrderDataAll(uid, 1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.UserGetDivisionOrderDataAll(uid, 1)
		} else {
			// 查询指定页数的数据
			data = dao.UserGetDivisionOrderDataAll(uid, intPage)
		}
	}

	// 查询总页数
	count := dao.UserGetOrderDataPage(uid)
	// 计算页数
	z := count / 100
	var y int64
	y = count % 100

	if y != 0 {
		votePage.Page = z + 1
	} else {
		votePage.Page = z
	}

	votePage.PageData = data

	return res.CodeSuccess, votePage
}

// GetOrderData 订单搜索
func GetOrderData(tp, state, s string) []model.Order {
	if tp == "订单类名" {
		switch state {
		case "等待中":
			return dao.GetOrderTypeData(s, -1)
		case "进行中":
			return dao.GetOrderTypeData(s, 0)
		case "已完成":
			return dao.GetOrderTypeData(s, 1)
		case "已终止":
			return dao.GetOrderTypeData(s, 2)
		case "退款中":
			return dao.GetOrderTypeData(s, 3)
		case "已退款":
			return dao.GetOrderTypeData(s, 4)
		}
	} else if tp == "订单号" {
		return dao.GetOrderData(s)
	} else {
		return dao.GetOrderTaskData(s)
	}
	return nil
}

// UserOrderRefund 订单退款
func UserOrderRefund(UID any, order string) (res.ResCode, string) {
	dao.UserOrderRefund(UID, order)
	go func() {
		// 查询订单信息
		orderData := dao.GetOneOrderData(order)
		// 查询项目信息
		projectData := dao.ProjectSearchOneTrue(orderData.OrderTaskType)

		fd := dao.GetForwardByID(projectData.ProjectAPI)
		if fd.ForwardAddress != "" {
			var sorf SOrderRF
			sorf.OrderID = orderData.OrderID
			sorf.OrderUID = orderData.OrderUID
			SendOrderRefundRequest(sorf, fd.ForwardRefundAddress, 0)
		}
	}()
	return res.CodeSuccess, "已申请退单"
}

// StartVote 发起任务
func StartVote(uid any, p *model.VoteOrder, IP string) (res.ResCode, string) {
	// 查询是否允许下单
	data, _ := dao.GetSetting("pao")
	if data.Value != "1" {
		return res.CodeVoteError, "系统正在维护，暂时无法下单..."
	}

	// 查询项目是否存在
	prData := dao.ProjectSearchOneTrue(p.OrderTaskType)
	if prData.ProjectType == "" {
		return res.CodeVoteError, "项目不存在"
	}

	// 判断任务范围
	zap.L().Debug("任务票数：" + strconv.Itoa(p.OrderNumber))
	if p.OrderNumber < 1 || p.OrderNumber > 50000 {
		return res.CodeVoteError, "任务票数不正确"
	}

	// 获取用户信息
	userData := dao.GetUserIDData(uid)

	// 判断用户点券是否足够此任务
	if userData.Points < int64(prData.ProjectPrice*p.OrderNumber) {
		return res.CodeVoteError, "余额不足"
	}

	// 效验IP下单地址
	data, _ = dao.GetSetting("check_ip")
	if data.Value == "1" {
		if !ip.CheckIf(userData.LoginIP, IP) {
			return res.CodeVoteError, "安全警告：账号环境异常，禁止下单"
		}
	}

	vote := &model.Order{
		OrderTaskType:    p.OrderTaskType,
		OrderID:          strconv.FormatInt(snowflake.GenID(), 10),
		OrderUID:         userData.UserID,
		OrderTickets:     int64(prData.ProjectPrice * p.OrderNumber),
		OrderNumber:      p.OrderNumber,
		OrderVariable:    p.OrderVariable,
		OrderRemarks:     p.OrderRemarks,
		OrderState:       -1,
		OrderStateReason: "",
		OrderIP:          IP,
	}

	// 订单储存进数据库
	dao.CreateOrder(vote)
	zap.L().Debug("StartVote 订单储存进数据库")

	// 用户预扣费
	zap.L().Debug("用户预扣费前：" + strconv.FormatInt(userData.Points, 10))
	userData.Points -= int64(prData.ProjectPrice * p.OrderNumber)
	// 保存数据
	dao.UpdateUser(userData)
	zap.L().Debug("用户预扣费后：" + strconv.FormatInt(userData.Points, 10))

	// 记录积分变动
	uroc := &model.UserRecordsOfConsumption{
		UserID:      userData.UserID,
		VoteOrder:   vote.OrderID,
		TaskState:   "消费",
		VoteTickets: int64(prData.ProjectPrice * p.OrderNumber),
	}
	go dao.CreateConsumption(uroc)

	// 发送订单请求
	go func() {
		// 查询订单转发地址
		fd := dao.GetForwardByID(prData.ProjectAPI)
		if fd.ForwardAddress != "" {
			SendOrderRequest(vote, fd.ForwardAddress, 0)
		}
	}()

	return res.CodeSuccess, "订单创建成功"
}

// SendOrderRequest 发送订单请求
func SendOrderRequest(so *model.Order, url string, num int) {
	zap.L().Debug("SendOrderRequest 发送订单请求, 订单ID：" + so.OrderID)
	sso, err := json.Marshal(so)
	if err != nil {
		zap.L().Error("发送订单请求失败：" + err.Error())
		OrderDelayPush(so, url, num+1)
	}
	// 发送请求
	result, err := requests.Requests("POST", url, string(sso), "")
	if err != nil {
		OrderDelayPush(so, url, num+1)
	} else {
		// 解析返回数据
		var sor Result
		err = json.Unmarshal(result, &sor)
		if err != nil {
			OrderDelayPush(so, url, num+1)
		} else {
			if sor.Code != 2000 {
				OrderDelayPush(so, url, num+1)
			}
		}
	}
}

// SendOrderRefundRequest 发送退单请求
func SendOrderRefundRequest(rf SOrderRF, url string, num int) {
	zap.L().Debug("SendOrderRefundRequest 发送退单请求, 订单ID：" + rf.OrderID)
	sorf, err := json.Marshal(rf)
	if err != nil {
		zap.L().Error("发送订单请求失败：" + err.Error())
		OrderRefundDelayPush(rf, url, num+1)
	}
	// 发送请求
	result, err := requests.Requests("POST", url, string(sorf), "")
	if err != nil {
		OrderRefundDelayPush(rf, url, num+1)
	} else {
		// 解析返回数据
		var sor Result
		err = json.Unmarshal(result, &sor)
		if err != nil {
			OrderRefundDelayPush(rf, url, num+1)
		} else {
			if sor.Code != 2000 {
				OrderRefundDelayPush(rf, url, num+1)
			}
		}
	}
}

// OrderDelayPush 订单延迟推送
func OrderDelayPush(so *model.Order, url string, num int) {
	if num <= 5 {
		// 重试次数小于5次
		time.Sleep(time.Minute * 3)
		SendOrderRequest(so, url, num+1)
	} else {
		// 重试次数大于5次
		zap.L().Error("[任务]订单推送失败：" + so.OrderID)
		// 查询订单当前订单信息
		od := dao.GetOneOrderData(so.OrderID)
		// 查询用户信息
		ud := dao.GetUserIDData(od.OrderUID)
		ud.Points += od.OrderTickets
		go dao.UpdateUser(ud)

		// 记录退款数据
		uroc := &model.UserRecordsOfConsumption{
			UserID:      ud.UserID,
			VoteOrder:   od.OrderID,
			TaskState:   "退回",
			VoteTickets: od.OrderTickets,
		}
		go dao.CreateConsumption(uroc)

		// 修改订单状态
		if od.OrderState == -1 || od.OrderState == 0 || od.OrderState == 3 {
			od.OrderState = 4
			od.OrderStateReason = "订单暂时无法处理，已退款"
			go dao.SaveOrder(od)
		}
	}
}

// OrderRefundDelayPush 退款订单延迟推送
func OrderRefundDelayPush(rf SOrderRF, url string, num int) {
	if num <= 5 {
		// 重试次数小于5次
		time.Sleep(time.Minute * 3)
		SendOrderRefundRequest(rf, url, num+1)
	} else {
		// 重试次数大于5次
		zap.L().Error("[退款]订单推送失败：" + rf.OrderID)
	}
}
