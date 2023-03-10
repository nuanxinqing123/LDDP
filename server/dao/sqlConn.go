package dao

import (
	"LDDP/server/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"os"
	"time"
)

var DB *gorm.DB
var err error
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Init() {
	// 配置日志
	var newLogger logger.Interface
	if viper.GetString("app.mode") == "debug" {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				LogLevel:                  logger.Warn, // 日志级别
				IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		)
	} else {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				LogLevel:                  logger.Error, // 日志级别
				IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,         // 禁用彩色打印
			},
		)
	}

	// 连接MySQL
	DB, err = gorm.Open(sqlite.Open("config/app.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		zap.L().Error("SQLite 发生错误, 原因：" + err.Error())
		panic(err.Error())
	}

	// 自动迁移
	err = DB.AutoMigrate(
		&model.User{},
		&model.JWTAdmin{},
		&model.Tickets{},
		&model.Recharge{},
		&model.Project{},
		&model.Forward{},
		&model.Order{},
		&model.UserRecordsOfConsumption{},
		&model.LoginIpData{},
		&model.WebSettings{},
		&model.BotConfig{})

	if err != nil {
		zap.L().Error("SQLite 自动迁移失败, 原因：" + err.Error())
		panic(err.Error())
	}

	return
}

// InitWebSettings 初始化数据表
func InitWebSettings() {
	// 检查JWT密钥表是否存在
	jwtKey := GetJWTKey()
	if jwtKey == "" || len(jwtKey) < 10 {
		// 生成密码并写入数据库
		b := make([]rune, 18)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := range b {
			b[i] = letters[r.Intn(62)]
		}
		zap.L().Debug("生成密钥：" + string(b))
		CreateJWTKey(string(b))
	}

	// 判断Settings是否是第一次创建
	settings, err := GetSettings()
	if err != nil {
		zap.L().Error("InitWebSettings 发生错误")
		panic(err.Error())
	}
	if len(settings) == 0 {
		zap.L().Debug("Init WebSettings")
		p := &[]model.WebSettings{
			{Key: "web_name", Value: "LDDP"},       // 网站名称
			{Key: "register", Value: "1"},          // 全局注册开关(1、打开 2、关闭)
			{Key: "pao", Value: "1"},               // 全局下单开关(1、打开 2、关闭)
			{Key: "check_ip", Value: "2"},          // 是否校验IP下单(1、打开 2、关闭)
			{Key: "notice_switch", Value: "1"},     // 弹窗公告开关(1、打开 2、关闭)
			{Key: "notice", Value: "1"},            // 弹窗公告
			{Key: "task_notice", Value: "<p></p>"}, // 任务公告
		}

		err = SaveSettings(p)
		if err != nil {
			zap.L().Error("InitWebSettings 发生错误")
			panic(err.Error())
		}
	}
}
