package logic

import (
	"LDDP/server/dao"
	"LDDP/server/model"
	res "LDDP/utils/response"
	"go.uber.org/zap"
)

// GetSettings 获取所有配置信息
func GetSettings() (res.ResCode, interface{}) {
	data, err := dao.GetSettings()
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, nil
	}

	return res.CodeSuccess, data
}

// SaveSettings 保存网站信息
func SaveSettings(p *[]model.WebSettings) res.ResCode {
	if err := dao.SaveSettings(p); err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	return res.CodeSuccess
}

// GetSetting 获取一个配置信息
func GetSetting(name string) (res.ResCode, model.WebSettings) {
	data, err := dao.GetSetting(name)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, data
	}

	// 限制前端只能获取公告信息
	return res.CodeSuccess, data
}
