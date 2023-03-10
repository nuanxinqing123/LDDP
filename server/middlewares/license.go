package middlewares

import (
	_const "LDDP/server/const"
	res "LDDP/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LicenseCheck 检查运行许可
func LicenseCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查授权
		if _const.LicenseState {
			c.Next()
		} else {
			// 许可到期
			zap.L().Info("许可失效，停止服务")
			zap.L().Warn("许可失效，停止服务")
			zap.L().Error("许可失效，停止服务")
			res.ResError(c, res.CodeServerBusy)
			c.Abort()
			return
		}
	}
}
