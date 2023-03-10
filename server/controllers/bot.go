package controllers

import (
	"LDDP/server/dao"
	"LDDP/server/logic"
	"LDDP/server/model"
	res "LDDP/utils/response"
	val "LDDP/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetBotAPISetting 获取配置
func GetBotAPISetting(c *gin.Context) {
	// 处理业务
	resCode, data := logic.GetBotAPISetting()
	switch resCode {
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, data)
	}
}

// SaveBotAPISetting 修改配置
func SaveBotAPISetting(c *gin.Context) {
	// 获取参数
	p := new(model.BotConfig)
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
	resCode := logic.SaveBotAPISetting(p)
	switch resCode {
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, "修改成功")
	}
}

// BotUserDR 消费/扣除/退回 用户点券
func BotUserDR(c *gin.Context) {
	// 获取参数
	p := new(model.DeductionRefund)
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
	resCode, msg := logic.BotUserDR(p)
	switch resCode {
	case res.CodeBotError:
		res.ResErrorWithMsg(c, res.CodeBotError, msg)
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "更新成功")
	}
}

// BotOrderData 获取项目订单列表
func BotOrderData(c *gin.Context) {
	tp := c.Query("type")     // 任务类名
	state := c.Query("state") // 任务状态

	// 处理业务
	resCode, data := logic.BotOrderData(tp, state)
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, data)
	}
}

// BotOrderUpdate 修改订单信息
func BotOrderUpdate(c *gin.Context) {
	// 获取参数
	p := new(model.BotOrderUpdate)
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
	resCode, msg := logic.BotOrderUpdate(p)
	switch resCode {
	case res.CodeBotError:
		res.ResErrorWithMsg(c, res.CodeBotError, msg)
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "更新成功")
	}
}

// BotOrderSearch 获取具体订单号订单信息
func BotOrderSearch(c *gin.Context) {
	order := c.Query("order") // 订单号

	// 处理业务
	//resCode, data := dao.GetOneOrderData(order)
	data := dao.GetOneOrderData(order)
	//switch resCode {
	//case res.CodeSuccess:
	//	// 更新成功
	//	res.ResSuccess(c, data)
	//}
	res.ResSuccess(c, data)
}
