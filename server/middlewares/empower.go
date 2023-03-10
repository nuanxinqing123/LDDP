package middlewares

import (
	_const "LDDP/server/const"
	res "LDDP/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// VersionEmpower 版本授权
func VersionEmpower() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查授权版本
		if _const.EmpowerVersion == "v2" {
			c.Next()
		} else {
			// 此版本暂不支持该功能
			zap.L().Info("此版本暂不支持该功能")
			res.ResErrorWithMsg(c, res.CodeVersionError, "此版本暂不支持该功能")
			c.Abort()
			return
		}
	}
}
