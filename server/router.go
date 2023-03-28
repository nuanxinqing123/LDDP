package server

import (
	"LDDP/server/controllers"
	"LDDP/server/middlewares"
	"LDDP/static/bindata"
	"LDDP/utils/logger"
	"html/template"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() *gin.Engine {
	// 创建服务
	r := gin.Default()

	// 配置中间件
	{
		// 配置日志
		if viper.GetString("app.mode") == "" {
			r.Use(logger.GinLogger(), logger.GinRecovery(true))
		}

		// 配置跨域
		if viper.GetString("app.mode") == "debug" {
			r.Use(middlewares.Cors())
		}

		// 检查授权
		r.Use(middlewares.LicenseCheck())
	}

	// 前端静态文件
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		r.SetHTMLTemplate(t)

		// 加载静态文件
		fs := assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		r.StaticFS("/static", &fs)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// 路由组
	{
		// 开放权限组
		open := r.Group("v1/api")
		{
			// 生成验证码
			open.GET("verification/code", controllers.CreateVerificationCode)
			// 用户注册
			open.POST("user/signup", controllers.SignUpHandle)
			// 用户登录
			open.POST("user/signin", controllers.SignInHandle)
			// 登录异常-发送验证码
			open.POST("user/abnormal/code", middlewares.RateLimitMiddleware(time.Second, 2, 1), controllers.UserAbnormalCode)
			// 登录异常-登录
			open.POST("user/abnormal/signin", controllers.UserAbnormalSignin)
			// 找回密码 - 发送验证码
			open.POST("findpwd/message", middlewares.RateLimitMiddleware(time.Second, 2, 1), controllers.UserFindPwd)
			// 找回密码 - 修改密码
			open.POST("findpwd/repwd", controllers.UserRePwd)
			// 订单转发接口演示
			open.GET("order/forward/demonstrate", controllers.OrderForwardDemonstrate)
			// 订单退款接口演示
			open.GET("order/refund/demonstrate", controllers.OrderRefundDemonstrate)
		}

		// 用户权限组
		user := r.Group("v2/api")
		user.Use(middlewares.UserAuth())
		{
			// 用户
			{
				// 信息获取
				user.GET("user/data", controllers.GetUserOneData)
				// 充值点券
				user.POST("user/recharge/tickets", middlewares.RateLimitMiddleware(time.Second, 100, 100), middlewares.VersionEmpower(), controllers.UserRechargeTickets)
				// 获取配置文件
				user.GET("get/setting", controllers.GetSetting)
				// 退出登录
				user.POST("user/logout", controllers.UserLogout)
			}

			// 任务
			{
				// 可选任务列表
				user.GET("user/task/list", controllers.UserTaskList)
				// 任务发起
				user.POST("user/vote/start", middlewares.RateLimitMiddleware(time.Second, 100, 100), controllers.StartVote)
			}

			// 订单
			{
				// 订单数据
				user.GET("user/order/data", controllers.UserOrderData)
				// 分页查询
				user.GET("user/order/division/data", controllers.UserGetDivisionOrderData)
				// 退单
				user.GET("user/order/refund", middlewares.RateLimitMiddleware(time.Second, 100, 100), controllers.UserOrderRefund)
			}

			// 消费
			{
				// 分页查询
				user.GET("user/consumption/division/data", controllers.UserConsumptionDivisionData)
			}

			// 充值
			{
				// 分页查询
				user.GET("user/recharge/division/data", controllers.UserRechargeLogData)
			}

			// 登录数据
			{
				// 分页查询
				user.GET("user/login/division/data", controllers.UserLoginDivisionData)
			}
		}

		// 管理员权限组
		admin := r.Group("v3/api")
		admin.Use(middlewares.AdminAuth())
		{
			// 首页
			{
				// 网站数据
				admin.GET("admin/panel/data", controllers.AdminPanelData)
				// 网站七日图表数据
				admin.GET("admin/panel/data/chart", controllers.GetOrderChart)
			}

			// 用户
			{
				// 分页查询
				admin.GET("user/division/data", controllers.UserDivisionData)
				// 搜索
				admin.GET("user/search", controllers.UserSearch)
				// 修改
				admin.PUT("user/information/update", controllers.UserInformationUpdate)
				// 删除
				//admin.DELETE("user/information/delete", controllers.UserInformationDelete)
			}

			// 点券
			{
				// 分页查询
				admin.GET("tickets/division/data", controllers.TicketsDivisionData)
				// 搜索
				admin.GET("tickets/search", middlewares.VersionEmpower(), controllers.TicketsSearch)
				// 新增
				admin.POST("tickets/add", middlewares.VersionEmpower(), controllers.TicketsAdd)
				// 下载卡密文件
				admin.GET("tickets/data/download", middlewares.VersionEmpower(), controllers.TicketsDataDownload)
				// 删除
				admin.DELETE("tickets/delete", middlewares.VersionEmpower(), controllers.TicketsDelete)
			}

			// 项目
			{
				// 分页查询
				admin.GET("project/division/data", controllers.ProjectDivisionData)
				// 筛选查询
				admin.GET("project/search", controllers.ProjectSearch)
				// 创建
				admin.POST("project/add", controllers.ProjectAdd)
				// 修改
				admin.PUT("project/update", controllers.ProjectUpdate)
				// 删除
				admin.DELETE("project/delete", controllers.ProjectDelete)
			}

			// 订单转发
			{
				// 简约查询
				admin.GET("order/forward/simple", controllers.OrderForwardSimple)
				// 分页查询
				admin.GET("order/forward/division/data", controllers.OrderForwardDivisionData)
				// 搜索
				admin.GET("order/forward/search", middlewares.VersionEmpower(), controllers.OrderForwardSearch)
				// 添加
				admin.POST("order/forward/add", middlewares.VersionEmpower(), controllers.OrderForwardAdd)
				// 修改
				admin.PUT("order/forward/update", middlewares.VersionEmpower(), controllers.OrderForwardUpdate)
				// 删除
				admin.DELETE("order/forward/delete", middlewares.VersionEmpower(), controllers.OrderForwardDelete)
				// API测试
				admin.POST("order/forward/api/test", middlewares.VersionEmpower(), controllers.OrderForwardApiTest)
			}

			// 订单
			{
				// 分页查询
				admin.GET("order/division/data", controllers.GetDivisionOrderData)
				// 搜索
				admin.GET("order/search", controllers.GetOrderData)
			}

			// 消费
			{
				// 分页查询
				admin.GET("consumption/division/data", controllers.ConsumptionLogData)
				// 搜索
				admin.GET("consumption/search", controllers.ConsumptionSearch)
			}

			// 充值
			{
				// 分页查询
				admin.GET("recharge/division/data", controllers.RechargeLogData)
				// 搜索
				admin.GET("recharge/search", controllers.RechargeSearch)
			}

			// 日志
			{
				// 分页查询：登录记录
				admin.GET("login/division/data", controllers.LoginDivisionData)
				// 搜索
				admin.GET("login/search", controllers.LoginSearch)
			}

			// BotAPI
			{
				// 获取配置
				admin.GET("bot-api/setting", controllers.GetBotAPISetting)
				// 修改配置
				admin.PUT("bot-api/setting", controllers.SaveBotAPISetting)
			}

			// 设置
			{
				// 获取全部配置
				admin.GET("set/settings", controllers.GetSettings)
				// 修改全部配置
				admin.PUT("set/settings", controllers.SaveSettings)
			}
		}

		// 机器人权限组
		bot := r.Group("bot/api")
		bot.Use(middlewares.BotAuth())
		{
			// 订单
			{
				// 获取具体订单号订单信息
				bot.GET("order/search", controllers.BotOrderSearch)
				// 获取项目订单列表
				bot.GET("order/data", controllers.BotOrderData)
				// 修改订单信息
				bot.PUT("order/update", controllers.BotOrderUpdate)
			}

			// 项目
			{
				// 获取全部项目数据
				bot.GET("project/division/data", controllers.ProjectDivisionDataAll)
			}

			// 用户
			{
				// 消费/扣除/退回 用户点券
				bot.PUT("user/dr", controllers.BotUserDR)
			}
		}
	}

	return r
}

// loadTemplate 加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name = strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
