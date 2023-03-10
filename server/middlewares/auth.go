package middlewares

import (
	"LDDP/server/dao"
	"LDDP/utils/jwt"
	res "LDDP/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

const (
	CtxUserKey = "userKey"
)

// UserAuth 用户基于JWT的认证中间件（非缓存）
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			c.Abort()
			return
		}
		// 检查是否为真实用户
		cUser := dao.GetUserIDData(mc.UserID)

		if cUser.Email == "" {
			c.Abort()
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			return
		} else {
			if cUser.IsState == false {
				// 账号已被封禁
				res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户已被封禁")
				c.Abort()
				return
			} else {
				//将当前请求的userID信息保存到请求的上下文c上
				c.Set(CtxUserKey, cUser.UserID)
				c.Next()
			}
		}
	}
}

// AdminAuth 管理员JWT的认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			c.Abort()
			return
		}

		// 检查是否属于管理员
		cAdmin := dao.GetUserIDData(mc.UserID)

		if cAdmin.Role != "admin" {
			c.Abort()
			res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户状态已失效")
			return
		} else {
			if cAdmin.IsState == false {
				// 账号已被封禁
				res.ResErrorWithMsg(c, res.CodeInvalidToken, "用户已被封禁")
				c.Abort()
				return
			} else {
				//将当前请求的userID信息保存到请求的上下文c上
				c.Set(CtxUserKey, cAdmin.UserID)
				c.Next()
			}
		}
	}
}

// BotAuth Bot密码IP的认证中间件
func BotAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求IP地址
		var RemoteIP string
		// 获取IP地址
		if "127.0.0.1" == c.RemoteIP() {
			RemoteIP = c.GetHeader("X-Real-IP")
		} else {
			RemoteIP = c.RemoteIP()
		}

		// 校验是否属于授权IP
		botData := dao.GetBotKey()
		ipL := strings.Split(botData.AuthIP, "&")
		ok := 0
		for _, i := range ipL {
			if i == RemoteIP {
				ok = 1
				break
			}
		}

		if ok == 0 {
			zap.L().Warn("疑似异常流量请求「授权IP错误」，IP地址：" + RemoteIP)
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}

		// 判断请求密码是否正确
		pd := c.Query("password")
		if botData.PassWord != pd {
			zap.L().Warn("疑似异常流量请求「Bot密码错误」，IP地址：" + RemoteIP)
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
