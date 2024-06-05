package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"LDDP/server"
	"LDDP/server/cron"
	"LDDP/server/dao"
	"LDDP/server/gcache"
	"LDDP/server/settings"
	"LDDP/utils/logger"
	"LDDP/utils/snowflake"
	"LDDP/utils/validator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("viper init failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Logger init failed, err:%v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())
	zap.L().Debug("Logger success init ...")

	// 初始化数据库
	dao.Init()
	dao.InitWebSettings()
	zap.L().Debug("SQLite success init ...")

	// 初始化翻译器
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("validator init failed, err:%v\n", err)
		return
	}
	zap.L().Debug("Validator success init ...")

	// 注册雪花ID算法
	if err := snowflake.Init(); err != nil {
		fmt.Printf("snowflake init failed, err:%v\n", err)
		return
	}
	zap.L().Debug("Snowflake success init ...")

	// 启动定时服务
	if err := cron.Task(); err != nil {
		fmt.Printf("Cron init failed, err:%v\n", err)
		return
	}
	zap.L().Debug("Cron success init ...")

	// 注册缓存
	gcache.InitCache()
	zap.L().Debug("Gcache success init ...")

	// 配置运行模式
	if viper.GetString("app.mode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册路由
	r := server.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	fmt.Println(" ")
	if viper.GetString("app.mode") == "debug" {
		fmt.Println("运行模式：Debug模式")
	} else {
		fmt.Println("运行模式：Release模式")
	}
	fmt.Println("监听端口：" + strconv.Itoa(viper.GetInt("app.port")))
	fmt.Println(" ")
	zap.L().Info("监听端口：" + strconv.Itoa(viper.GetInt("app.port")))

	// 启动
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置5秒超时
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	zap.L().Info("Service ready to shut down")

	// 创建五秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 十秒内优雅关闭服务（将未处理完成的请求处理完再关闭服务），超过十秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Service timed out has been shut down：", zap.Error(err))
	}

	zap.L().Info("Service has been shut down")
}
