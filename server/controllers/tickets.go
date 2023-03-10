package controllers

import (
	"LDDP/server/logic"
	"LDDP/server/model"
	res "LDDP/utils/response"
	val "LDDP/utils/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// TicketsDivisionData Tickets分页查询
func TicketsDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.TicketsDivisionData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// TicketsSearch Tickets数据查询
func TicketsSearch(c *gin.Context) {
	tp := c.Query("type")
	state := c.Query("state")
	s := c.Query("s")
	zap.L().Debug("【Tickets搜索】值：" + s)
	resCode, data := logic.TicketsSearch(tp, state, s)

	switch resCode {
	case res.CodeTicketsError:
		res.ResErrorWithMsg(c, res.CodeTicketsError, "查询数据不存在")
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// TicketsAdd 批量生成Tickets
func TicketsAdd(c *gin.Context) {
	// 获取参数
	p := new(model.CreateTickets)
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
	resCode := logic.TicketsAdd(p)
	switch resCode {
	case res.CodeServerBusy:
		res.ResErrorWithMsg(c, res.CodeServerBusy, "生成Tickets失败，请检查日志获取报错信息")
	case res.CodeTicketsError:
		res.ResErrorWithMsg(c, res.CodeTicketsError, "创建Tickets已写入数据库，但生成下载文件失败")
	case res.CodeSuccess:
		// 生成成功
		res.ResSuccess(c, "生成成功")
	}
}

// TicketsDataDownload 下载Tickets文件
func TicketsDataDownload(c *gin.Context) {
	Filename := "Tickets.txt"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", Filename))
	c.File("./" + Filename)
	go logic.TicketsDataDelete()
}

// TicketsDelete 删除Tickets数据
func TicketsDelete(c *gin.Context) {
	// 获取参数
	p := new(model.DelTickets)
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
	resCode := logic.TicketsDelete(p)
	switch resCode {
	case res.CodeTicketsError:
		res.ResErrorWithMsg(c, res.CodeTicketsError, "删除Tickets失败")
	case res.CodeSuccess:
		// 删除成功
		res.ResSuccess(c, "删除成功")
	}
}
