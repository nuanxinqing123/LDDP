package dao

import (
	"LDDP/server/model"
	"LDDP/utils/ip"
)

// CreateLoginInfo 创建用户IP登录记录
func CreateLoginInfo(IP string, uid, f string) {
	// 获取IP信息
	info := ip.GetIPInfo(IP)

	// 记录登录信息
	var i model.LoginIpData
	i.UserID = uid
	i.Ip = IP
	i.IpAddr = info.Province + info.City + info.Area
	i.IpIsp = info.Isp
	i.IpNet = info.Net
	i.LoginFunction = f
	DB.Create(&i)
}
