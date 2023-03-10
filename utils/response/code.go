// -*- coding: utf-8 -*-

package response

type ResCode int64

const (
	CodeSuccess ResCode = 2000

	CodeInvalidParam = 4999 + iota
	CodeServerBusy
	CodeInvalidRouterRequested

	CodeInvalidToken
	CodeNeedLogin
	CodeEmailError
	CodeLoginError
	CodeRegisterError
	CodeRePwdError
	CodeTicketsError
	CodeOrderError
	CodeVoteError
	CodeBotError
	CodeProjectError
	CodeAbnormalEnvironment
	CodeAbnormalError
	CodeVersionError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "Success",

	CodeInvalidParam:           "请求参数错误",
	CodeServerBusy:             "服务繁忙",
	CodeInvalidRouterRequested: "请求无效路由",

	CodeInvalidToken:        "无效的Token",
	CodeNeedLogin:           "未登录",
	CodeEmailError:          "邮箱错误",
	CodeLoginError:          "登录错误",
	CodeRegisterError:       "注册错误",
	CodeRePwdError:          "修改密码错误",
	CodeTicketsError:        "卡密错误",
	CodeOrderError:          "订单错误",
	CodeVoteError:           "任务错误",
	CodeBotError:            "操作错误",
	CodeProjectError:        "项目错误",
	CodeAbnormalEnvironment: "异常环境",
	CodeAbnormalError:       "异常错误",
	CodeVersionError:        "版本错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
