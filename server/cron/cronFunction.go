package cron

import (
	_const "LDDP/server/const"
	"LDDP/server/model"
	"LDDP/utils/requests"
	"github.com/goccy/go-json"
)

func CheckVersion() {
	// 版本号
	var v model.RemoteVersion

	// 获取仓库版本信息
	url := "https://version.6b7.xyz/lddp_version.json"
	r, errFunc := requests.Version("GET", url, "", "")
	if errFunc != nil {
		_const.LicenseState = false
		return
	}
	// 序列化内容
	errFunc = json.Unmarshal(r, &v)
	if errFunc != nil {
		_const.LicenseState = false
		return
	}

	if v.Version != _const.LocVersion {
		_const.LicenseState = false
	} else {
		_const.LicenseState = true
	}
}
