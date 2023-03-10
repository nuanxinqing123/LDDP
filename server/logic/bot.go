package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
)

// GetBotAPISetting 获取配置
func GetBotAPISetting() (res.ResCode, model.BotConfig) {
	return res.CodeSuccess, dao.GetBotKey()
}

// SaveBotAPISetting 修改配置
func SaveBotAPISetting(p *model.BotConfig) res.ResCode {
	// 删除最后一位&符号
	p.AuthIP = p.AuthIP[:len(p.AuthIP)-1]
	dao.UpdateBotKey(p)
	return res.CodeSuccess
}

// BotUserDR 消费/扣除/退回 用户点券
func BotUserDR(p *model.DeductionRefund) (res.ResCode, string) {
	// 查询订单
	order := dao.GetOneOrderData(p.OrderID)
	if order.OrderID == "" {
		return res.CodeBotError, "查询订单不存在"
	}

	// 查询用户数据
	user := dao.GetUserIDData(order.OrderUID)

	if p.Type == 1 {
		// 扣除
		user.Points -= int64(p.Points)
		dao.UpdateUser(user)

		// 记录积分变动
		uroc := &model.UserRecordsOfConsumption{
			UserID:      user.UserID,
			VoteOrder:   order.OrderID,
			TaskState:   "消费",
			VoteTickets: int64(p.Points),
		}
		go dao.CreateConsumption(uroc)
	} else {
		// 退款
		user.Points += int64(p.Points)
		dao.UpdateUser(user)

		// 记录积分变动
		uroc := &model.UserRecordsOfConsumption{
			UserID:      user.UserID,
			VoteOrder:   order.OrderID,
			TaskState:   "退回",
			VoteTickets: int64(p.Points),
		}
		go dao.CreateConsumption(uroc)
	}

	return res.CodeSuccess, "完成"
}

// BotOrderData 获取项目订单列表
func BotOrderData(t, s string) (res.ResCode, []model.Order) {
	// 查询任务类名下所有等待状态订单
	return res.CodeSuccess, dao.GetOrderType(t, s)
}

// BotOrderUpdate 修改订单信息
func BotOrderUpdate(p *model.BotOrderUpdate) (res.ResCode, string) {
	// 查询订单
	o := dao.GetOneOrderData(p.OrderID)
	if o.OrderID == "" {
		return res.CodeBotError, "订单号不存在"
	}

	if p.OrderState == 2 {
		if p.OrderStateReason == "" {
			return res.CodeBotError, "无订单终止原因"
		}
	}

	if o.OrderState == 2 || o.OrderState == 1 || o.OrderState == 4 {
		return res.CodeBotError, "订单状态已禁止操作"
	}
	o.OrderState = p.OrderState
	o.OrderStateReason = p.OrderStateReason
	o.OrderStatus = p.OrderStatus

	// 保存订单
	dao.SaveOrder(o)
	return res.CodeSuccess, "完成"
}
