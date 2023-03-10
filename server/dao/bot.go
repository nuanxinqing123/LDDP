package dao

import "LDDP/server/model"

// GetBotKey 获取Bot配置信息
func GetBotKey() model.BotConfig {
	var bot model.BotConfig
	DB.First(&bot)
	return bot
}

// UpdateBotKey 更新Bot配置信息
func UpdateBotKey(p *model.BotConfig) {
	var bot model.BotConfig
	result := DB.First(&bot)
	if result.RowsAffected == 0 {
		DB.Create(&p)
		return
	} else {
		bot.PassWord = p.PassWord
		bot.AuthIP = p.AuthIP
		DB.Save(&bot)
	}
}
