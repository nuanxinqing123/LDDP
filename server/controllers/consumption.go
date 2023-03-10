package controllers

import (
	"LDDP/server/logic"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"github.com/gin-gonic/gin"
)

// ConsumptionLogData 消费记录：收费/退款：以20条数据分割
func ConsumptionLogData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.ConsumptionLogData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// ConsumptionSearch 消费记录：收费/退款：搜索
func ConsumptionSearch(c *gin.Context) {
	var data []model.UserRecordsOfConsumption
	var resCode res.ResCode

	s := c.Query("s")
	resCode, data = logic.ConsumptionSearch(s)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserConsumptionDivisionData 用户消费分页查询
func UserConsumptionDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserConsumptionDivisionData(UID, page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}
