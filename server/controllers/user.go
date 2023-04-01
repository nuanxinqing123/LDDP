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

const (
	CtxUserKey = "userKey"
)

// CreateVerificationCode 创建验证码
func CreateVerificationCode(c *gin.Context) {
	id, bs64, err := logic.CaptMake()
	if err != nil {
		zap.L().Error("[创建验证码]失败，原因：" + err.Error())
		res.ResError(c, res.CodeServerBusy)
	} else {
		res.ResSuccess(c, gin.H{
			"id":   id,
			"bs64": bs64,
		})
	}
}

// SignUpHandle 注册请求
func SignUpHandle(c *gin.Context) {
	// 获取参数
	p := new(model.UserSignUp)
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

	// 业务处理
	resCode, msg := logic.SignUp(p)
	switch resCode {
	case res.CodeRegisterError:
		// 注册错误
		res.ResErrorWithMsg(c, res.CodeRegisterError, msg)
	case res.CodeServerBusy:
		// 内部服务错误
		res.ResErrorWithMsg(c, res.CodeServerBusy, msg)
	case res.CodeSuccess:
		// 注册成功
		zap.L().Info("邮箱：" + p.Email + "，注册用户成功")
		res.ResSuccess(c, "注册完成")
	}
}

// SignInHandle 登录请求
func SignInHandle(c *gin.Context) {
	// 获取参数
	p := new(model.UserSignIn)
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

	resCode, msg := logic.SignIn(p, RemoteIP)
	switch resCode {
	case res.CodeLoginError:
		// 邮箱或者密码错误
		res.ResErrorWithMsg(c, res.CodeLoginError, msg)
	case res.CodeAbnormalEnvironment:
		res.ResErrorWithMsg(c, res.CodeAbnormalEnvironment, msg)
	case res.CodeServerBusy:
		// 生成Token出错
		res.ResErrorWithMsg(c, res.CodeServerBusy, msg)
	case res.CodeSuccess:
		// 登录成功,返回Token
		zap.L().Info("用户：" + p.Email + "，IP" + RemoteIP + "。登录成功")
		res.ResSuccess(c, gin.H{
			"token": msg,
		})
	}
}

// UserAbnormalCode 登录异常-发送验证码
func UserAbnormalCode(c *gin.Context) {
	// 获取参数
	p := new(model.UserAbnormalEmail)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

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
	resCode, msg := logic.AbnormalEmail(p)
	switch resCode {
	case res.CodeAbnormalError:
		res.ResErrorWithMsg(c, res.CodeAbnormalError, msg)
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 发送成功
		res.ResSuccess(c, "发送成功，请前往邮箱查看验证码")
	}
}

// UserAbnormalSignin 登录异常-登录
func UserAbnormalSignin(c *gin.Context) {
	// 获取参数
	p := new(model.UserAbnormalSignin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

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

	resCode, msg := logic.AbnormalSignin(p, RemoteIP)
	switch resCode {
	case res.CodeAbnormalError:
		// 修改密码错误
		res.ResErrorWithMsg(c, res.CodeAbnormalError, msg)
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 登录成功
		res.ResSuccess(c, msg)
	}
}

// UserFindPwd 找回密码 - 发送验证码
func UserFindPwd(c *gin.Context) {
	// 获取参数
	p := new(model.UserFindPwd)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

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
	resCode, msg := logic.UserFindPwd(p)
	switch resCode {
	case res.CodeRePwdError:
		// 修改密码错误
		res.ResErrorWithMsg(c, res.CodeRePwdError, msg)
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, "发送成功")
	}
}

// UserRePwd 找回密码 - 修改密码
func UserRePwd(c *gin.Context) {
	// 获取参数
	p := new(model.UserRePwd)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

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
	resCode, msg := logic.UserRePwd(p)
	switch resCode {
	case res.CodeRePwdError:
		// 修改密码错误
		res.ResErrorWithMsg(c, res.CodeRePwdError, msg)
	case res.CodeSuccess:
		// 修改成功
		var RemoteIP string
		UID, _ := c.Get(CtxUserKey)
		// 获取IP地址
		if "127.0.0.1" == c.RemoteIP() {
			RemoteIP = c.GetHeader("X-Real-IP")
		} else {
			RemoteIP = c.RemoteIP()
		}
		go dao.UpdateUserLoginIP(UID, RemoteIP)
		res.ResSuccess(c, "修改成功")
	}
}

// AdminPanelData 网站数据
func AdminPanelData(c *gin.Context) {
	resCode, data := logic.AdminPanelData()
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, data)
	}
}

// GetOrderChart 网站七日图表数据
func GetOrderChart(c *gin.Context) {
	resCode, data := logic.GetOrderChart()
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, data)
	}
}

// UserDivisionData 用户分页查询
func UserDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	resCode, data := logic.UserDivisionData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserSearch 用户筛选搜索
func UserSearch(c *gin.Context) {
	// 获取查询方法
	tp := c.Query("type")
	// 模糊查询
	s := c.Query("s")
	resCode, data := logic.UserSearch(tp, s)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// UserInformationUpdate 用户数据更新
func UserInformationUpdate(c *gin.Context) {
	// 获取参数
	p := new(model.UpdateUserData)
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
	resCode := logic.UserInformationUpdate(p)
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "更新成功")
	}
}

// LoginDivisionData 登录日志分页查询
func LoginDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	resCode, data := logic.LoginDivisionData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// LoginSearch 登录日志筛选搜索
func LoginSearch(c *gin.Context) {
	s := c.Query("s")
	data := dao.LoginSearch(s)
	res.ResSuccess(c, data)
}

// UserLoginDivisionData 用户登录日志分页查询
func UserLoginDivisionData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")

	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.UserLoginDivisionData(UID, page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// GetUserOneData 用户信息：获取
func GetUserOneData(c *gin.Context) {
	// 处理业务
	UID, _ := c.Get(CtxUserKey)
	resCode, data := logic.GetUserOneData(UID)
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, data)
	}
}

// UserLogout 退出登录
func UserLogout(c *gin.Context) {
	// 处理业务
	resCode := logic.UserLogout()
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "退出成功")
	}
}
