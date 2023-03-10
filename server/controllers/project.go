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

// ProjectDivisionData 分页查询
func ProjectDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.GetDivisionProjectData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// ProjectDivisionDataAll 获取全部项目数据
func ProjectDivisionDataAll(c *gin.Context) {
	resCode, data := logic.GetDivisionProjectDataAll()

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// ProjectSearch 筛选查询
func ProjectSearch(c *gin.Context) {
	tp := c.Query("type")
	s := c.Query("s")
	data := dao.ProjectSearch(tp, s)
	res.ResSuccess(c, data)
}

// ProjectAdd 创建项目
func ProjectAdd(c *gin.Context) {
	// 获取参数
	p := new(model.ProjectAdd)
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
	resCode, msg := logic.ProjectAdd(p)
	switch resCode {
	case res.CodeProjectError:
		res.ResErrorWithMsg(c, res.CodeProjectError, msg)
	case res.CodeSuccess:
		// 创建项目成功
		UID, _ := c.Get(CtxUserKey)
		zap.L().Info("操作：" + UID.(string) + " 创建项目[操作成功]")
		res.ResSuccess(c, "创建成功")
	}
}

// ProjectUpdate 更新项目
func ProjectUpdate(c *gin.Context) {
	// 获取参数
	p := new(model.ProjectUpdate)
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
	resCode := logic.ProjectUpdate(p)
	switch resCode {
	case res.CodeSuccess:
		// 修改项目成功
		UID, _ := c.Get(CtxUserKey)
		zap.L().Info("操作：" + UID.(string) + " 修改项目[操作成功]")
		res.ResSuccess(c, "更新成功")
	}
}

// ProjectDelete 删除Project数据
func ProjectDelete(c *gin.Context) {
	// 获取参数
	p := new(model.ProjectDelete)
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
	dao.DelProject(p)
	res.ResSuccess(c, "删除成功")
}

// UserTaskList 可选任务列表
func UserTaskList(c *gin.Context) {
	// 处理业务
	resCode, data := logic.UserTaskList()

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}
