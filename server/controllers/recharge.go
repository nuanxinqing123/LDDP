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

// RechargeLogData 充值数据：以20条数据分割
func RechargeLogData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.RechargeLogData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// RechargeSearch 充值数据：TicketsKey搜索/UserID搜索
func RechargeSearch(c *gin.Context) {
	s := c.Query("s")
	resCode, data := logic.RechargeSearch(s)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserRechargeLogData 充值数据：以20条数据分割
func UserRechargeLogData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserRechargeLogData(UID, page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserRechargeTickets 用户充值：用户点券充值
func UserRechargeTickets(c *gin.Context) {
	// 获取参数
	p := new(model.UserRecharge)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignUpHandle with invalid param", zap.Error(err))

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
	resCode, msg := logic.UserRechargeTickets(UID, p, RemoteIP)
	switch resCode {
	case res.CodeTicketsError:
		// 点券错误
		res.ResErrorWithMsg(c, res.CodeTicketsError, msg)
	case res.CodeSuccess:
		// 充值成功
		res.ResSuccess(c, "充值成功")
	}
}
