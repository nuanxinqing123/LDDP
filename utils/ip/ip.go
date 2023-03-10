package ip

import (
	"LDDP/server/model"
	"LDDP/utils/requests"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

// CheckIf 检查是否属于异地登录
func CheckIf(ip1, ip2 string) bool {
	/*
		ip1：原始IP
		ip2：登录IP
		bool1：是否异地登录
	*/
	if ip1 == ip2 {
		return true
	}

	// 查询IP地址
	url1 := "https://ip.useragentinfo.com/json?ip=" + ip1
	addr1, err := requests.Requests("GET", url1, "", "")
	if err != nil {
		return false
	}

	var l1 model.Location
	var l2 model.Location
	// 数据绑定
	err = json.Unmarshal(addr1, &l1)
	if err != nil {
		zap.L().Error("[数据解析]失败，原因：" + err.Error())
		return false
	}

	url2 := "https://ip.useragentinfo.com/json?ip=" + ip2
	addr2, err := requests.Requests("GET", url2, "", "")
	if err != nil {
		return false
	}
	// 数据绑定
	err = json.Unmarshal(addr2, &l2)
	if err != nil {
		zap.L().Error("[数据解析]失败，原因：" + err.Error())
		return false
	}

	if l1.City == l2.City {
		return true
	}

	return false
}

// GetIPInfo 获取IP详细信息
func GetIPInfo(ip1 string) model.Location {
	var l1 model.Location

	// 查询IP地址
	url1 := "https://ip.useragentinfo.com/json?ip=" + ip1
	addr1, err := requests.Requests("GET", url1, "", "")
	if err != nil {
		return l1
	}

	// 数据绑定
	err = json.Unmarshal(addr1, &l1)
	if err != nil {
		zap.L().Error(err.Error())
		return l1
	}

	return l1
}
