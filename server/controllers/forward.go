package controllers

import (
	"LDDP/server/logic"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"LDDP/utils/snowflake"
	val "LDDP/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func OrderForwardSimple(c *gin.Context) {
	// 处理业务
	resCode, data := logic.OrderForwardSimple()

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// OrderForwardDivisionData 分页查询
func OrderForwardDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.GetDivisionOrderForwardData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// OrderForwardSearch 搜索
func OrderForwardSearch(c *gin.Context) {
	// 获取搜索关键字
	s := c.Query("search")
	resCode, data := logic.GetSearchOrderForwardData(s)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// OrderForwardAdd 添加
func OrderForwardAdd(c *gin.Context) {
	// 获取参数
	p := new(model.ForwardAdd)
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
	resCode, msg := logic.OrderForwardAdd(p)
	switch resCode {
	case res.CodeProjectError:
		res.ResErrorWithMsg(c, res.CodeProjectError, msg)
	case res.CodeSuccess:
		res.ResSuccess(c, msg)
	}
}

// OrderForwardUpdate 更新
func OrderForwardUpdate(c *gin.Context) {
	// 获取参数
	p := new(model.ForwardUpdate)
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
	resCode, msg := logic.ForwardUpdate(p)
	switch resCode {
	case res.CodeProjectError:
		res.ResErrorWithMsg(c, res.CodeProjectError, msg)
	case res.CodeSuccess:
		res.ResSuccess(c, msg)
	}
}

// OrderForwardDelete 删除
func OrderForwardDelete(c *gin.Context) {
	// 获取参数
	p := new(model.ForwardDelete)
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
	resCode, msg := logic.ForwardDelete(p)
	switch resCode {
	case res.CodeProjectError:
		res.ResErrorWithMsg(c, res.CodeProjectError, msg)
	case res.CodeSuccess:
		// 删除成功
		res.ResSuccess(c, "删除成功")
	}
}

// OrderForwardApiTest API测试
func OrderForwardApiTest(c *gin.Context) {
	// 获取参数
	p := new(model.ForwardTest)
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
	UID, _ := c.Get(CtxUserKey)
	resCode, msg := logic.OrderForwardApiTest(UID, p)
	switch resCode {
	case res.CodeProjectError:
		res.ResErrorWithMsg(c, res.CodeProjectError, msg)
	case res.CodeSuccess:
		res.ResSuccess(c, msg)
	}
}

// OrderForwardDemonstrate 订单转发接口演示
func OrderForwardDemonstrate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"order_task_type":    "test",
		"order_id":           strconv.FormatInt(snowflake.GenID(), 10),
		"order_uid":          "3761986723141",
		"order_tickets":      int64(2 * 10),
		"order_number":       10,
		"order_variable":     "测试变量",
		"order_remarks":      "测试订单转发",
		"order_state":        -1,
		"order_state_reason": "",
		"order_status":       "",
		"order_ip":           "127.0.0.1",
	})
}

// OrderRefundDemonstrate 订单退款接口演示
func OrderRefundDemonstrate(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"order_id":  strconv.FormatInt(snowflake.GenID(), 10),
			"order_uid": "3761986723141",
		})
}
