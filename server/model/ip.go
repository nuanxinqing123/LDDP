package model

import "gorm.io/gorm"

type LoginIpData struct {
	gorm.Model
	UserID        string `json:"user_id,omitempty"`        // 登录用户ID
	Ip            string `json:"ip,omitempty"`             // IP地址
	IpAddr        string `json:"ip_addr"`                  // IP物理地址
	IpIsp         string `json:"ip_isp,omitempty"`         // IP运营商
	IpNet         string `json:"ip_net,omitempty"`         // IP类型
	LoginFunction string `json:"login_function,omitempty"` //  登录方式
}

// LoginPageData 用户分页数据
type LoginPageData struct {
	Page     int64         `json:"page"`
	PageData []LoginIpData `json:"page_data"`
}

// Location IP 地址序列化
type Location struct {
	Country   string `json:"country"`
	ShortName string `json:"short_name"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Area      string `json:"area"`
	Isp       string `json:"isp"`
	Net       string `json:"net"`
	Ip        string `json:"ip"`
	Code      int    `json:"code"`
	Desc      string `json:"desc"`
}
