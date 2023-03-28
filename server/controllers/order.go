package controllers

import (
	"LDDP/server/logic"
	"LDDP/server/model"
	res "LDDP/utils/response"
	val "LDDP/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetDivisionOrderData 订单分页查询
func GetDivisionOrderData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.GetDivisionOrderData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// GetOrderData 搜索订单
func GetOrderData(c *gin.Context) {
	// 查询搜索数据
	tp := c.Query("type")
	state := c.Query("state")
	s := c.Query("s")
	zap.L().Debug("【订单搜索】方法：" + tp + " 状态：" + state + " 值：" + s)

	res.ResSuccess(c, logic.GetOrderData(tp, state, s))
}

// UserOrderData 订单数据
func UserOrderData(c *gin.Context) {
	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserOrderData(UID)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserGetDivisionOrderData 订单分页查询
func UserGetDivisionOrderData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserGetDivisionOrderData(UID, page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserOrderRefund 订单退款
func UserOrderRefund(c *gin.Context) {
	// 获取订单号
	order := c.Query("order")

	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserOrderRefund(UID, order)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// StartVote 发起订单任务
func StartVote(c *gin.Context) {
	// 获取参数
	p := new(model.VoteOrder)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("Handle with invalid param", zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.ResError(c, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(c, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 处理业务
	var RemoteIP string
	// 获取IP地址
	if "127.0.0.1" == c.RemoteIP() {
		RemoteIP = c.GetHeader("X-Real-IP")
	} else {
		RemoteIP = c.RemoteIP()
	}

	UID, _ := c.Get(CtxUserKey)
	resCode, msg := logic.StartVote(UID, p, RemoteIP)
	switch resCode {
	case res.CodeVoteError:
		res.ResErrorWithMsg(c, res.CodeVoteError, msg)
	case res.CodeSuccess:
		// 任务创建成功
		zap.L().Info("下单用户UID：" + UID.(string) + "、下单IP地址：" + RemoteIP)
		res.ResSuccess(c, "任务创建成功")
	}
}
