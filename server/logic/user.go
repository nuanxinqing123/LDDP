package logic

import (
	"LDDP/server/dao"
	"LDDP/server/gcache"
	"LDDP/server/model"
	"LDDP/utils/email"
	"LDDP/utils/encryption"
	"LDDP/utils/jwt"
	"LDDP/utils/requests"
	res "LDDP/utils/response"
	"LDDP/utils/snowflake"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// 代替官方JSON库
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// IP 地址序列化
type location struct {
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

// CaptMake 生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverMath

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverMath{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}

// SignUp 注册业务
func SignUp(p *model.UserSignUp) (res.ResCode, string) {
	// 检查验证码是否正确
	if !CaptVerify(p.Id, p.Capt) {
		return res.CodeRegisterError, "验证码错误"
	}

	// 查询是否允许注册
	data, _ := dao.GetSetting("register")
	if data.Value != "1" {
		return res.CodeRegisterError, "暂时不开放注册..."
	}

	// 验证邮箱格式
	if !email.VerifyEmailFormat(p.Email) {
		return res.CodeRegisterError, "邮件格式错误"
	}

	// 判断是否已存在账户
	result, _ := dao.EmailGetUserData(p.Email)
	if result == true {
		return res.CodeRegisterError, "当前邮箱已被注册"
	}

	// 构造User实例
	user := &model.User{
		UserID:  strconv.FormatInt(snowflake.GenID(), 10),
		Email:   p.Email,
		UserKey: encryption.Sha1(p.UserKey),
		Points:  0,
		Role:    "user",
		IsState: true,
		LoginIP: "",
	}

	// 判断是否为第一个账号
	if dao.GetUserRecord() == 0 {
		// 第一个注册账号为管理员账号
		user.Role = "admin"
	}

	// 保存进数据库
	if err := dao.InsertUser(user); err != nil {
		zap.L().Error("Error inserting database, err:", zap.Error(err))
		return res.CodeServerBusy, "服务繁忙"
	}
	return res.CodeSuccess, "注册成功"
}

// SignIn 登录业务
func SignIn(p *model.UserSignIn, RemoteIP string) (res.ResCode, string) {
	// 检查验证码是否正确
	if !CaptVerify(p.Id, p.Capt) {
		return res.CodeLoginError, "验证码错误"
	}

	// 检查用户名是否存在
	result, user := dao.EmailGetUserData(p.Email)
	if result == false {
		// 不存在
		zap.L().Warn("【登录业务：用户名不存在】登录IP：" + RemoteIP + "、登录结果：失败")
		return res.CodeLoginError, "用户不存在"
	} else {
		// 检查用户是否被封禁
		if user.IsState == false {
			zap.L().Warn("【登录业务：账户已被封禁】登录IP：" + RemoteIP + "、登录邮箱：" + p.Email + "、登录结果：失败")
			return res.CodeLoginError, "账户已被封禁"
		}

		// 判断密码是否正确
		if user.UserKey != encryption.Sha1(p.UserKey) {
			zap.L().Warn("【登录业务：密码错误】登录IP：" + RemoteIP + "、登录邮箱：" + p.Email + "、登录结果：失败")
			return res.CodeLoginError, "密码错误"
		} else {
			// 密码正确, 校验是否异地登录
			if s, _ := dao.GetSetting("check_ip"); s.Value == "1" {
				if user.LoginIP != "" {
					b1, b2 := CheckIf(user.LoginIP, RemoteIP)
					if b2 {
						return res.CodeServerBusy, "服务繁忙"
					} else {
						if !b1 {
							// 异地登录
							return res.CodeAbnormalEnvironment, "用户登录环境异常"
						}
					}
				}
			}

			// 密码正确, 返回生成的Token（userSecret：密码前六位）
			var token string
			var err error
			if p.IsSave {
				token, err = jwt.GenToken(user.UserID, user.UserKey[:6], 7)
			} else {
				token, err = jwt.GenToken(user.UserID, user.UserKey[:6], 1)
			}
			if err != nil {
				zap.L().Error("An error occurred in token generation, err:", zap.Error(err))
				return res.CodeServerBusy, "服务繁忙"
			}

			// 记录登录IP
			go dao.UpdateUserLoginIP(user.UserID, RemoteIP)
			return res.CodeSuccess, token
		}
	}
}

// AbnormalEmail 登录异常 - 发送验证码
func AbnormalEmail(p *model.UserAbnormalEmail) (res.ResCode, string) {
	// 判断是否已存在账户
	result, user := dao.EmailGetUserData(p.Email)
	if result == false {
		return res.CodeAbnormalError, "用户不存在"
	}

	// 缓存查询验证码
	_, err := gcache.GetCache(user.Email + "_login")
	if err != nil {
		var f model.SpecialUsersCache

		// 生成验证码, 发送邮件
		rand.Seed(time.Now().UnixNano())
		bytes := make([]byte, 5)
		for i := 0; i < 5; i++ {
			b := rand.Intn(26) + 65
			bytes[i] = byte(b)
		}
		str := "您的登录验证码为：" + string(bytes) + "， (5分钟内有效，本邮件由系统自动发出，请勿直接回复)"
		zap.L().Debug("str地址：" + str)

		// 发送邮件
		if err := email.SendEmailCode("LDDP - 登录验证码", user.Email, str); err != nil {
			zap.L().Error("邮件发送失败，err:", zap.Error(err))
			return res.CodeAbnormalError, "邮件发送失败，请稍等片刻再尝试"
		}

		// 发送成功，数据存入缓存
		f.UserId = user.UserID
		f.Code = string(bytes)
		v, err := json.Marshal(f)
		if err != nil {
			return res.CodeAbnormalError, "邮件发送失败，请稍等片刻再尝试"
		}
		go gcache.TimingCache(user.Email+"_login", string(v), time.Minute*5)

		// 返回
		return res.CodeSuccess, "验证码发送成功"
	}
	return res.CodeAbnormalError, "存在未过期验证码，请勿重复发送"
}

// AbnormalSignin 登录异常 - 登录
func AbnormalSignin(p *model.UserAbnormalSignin, RemoteIP string) (res.ResCode, string) {
	// 缓存查询Token
	uTk, err := gcache.GetCache(p.Email + "_login")
	if err != nil {
		zap.L().Error("[登录异常-登录]失败，原因：" + err.Error())
		return res.CodeAbnormalError, "暂无登录验证码"
	}

	// 序列化字符串
	var f model.SpecialUsersCache
	err = json.Unmarshal([]byte(uTk.(string)), &f)
	if err != nil {
		zap.L().Error("[登录异常-登录]失败，原因：" + err.Error())
		return res.CodeServerBusy, "业务繁忙"
	}

	if p.VfCode != f.Code {
		return res.CodeAbnormalError, "验证码错误"
	}

	// 登录验证, 检查用户名是否存在
	_, user := dao.EmailGetUserData(p.Email)
	// 检查用户是否被封禁
	if user.IsState == false {
		return res.CodeAbnormalError, "账户已被封禁"
	}

	// 判断密码是否正确
	if user.UserKey != encryption.Sha1(p.UserKey) {
		return res.CodeAbnormalError, "密码错误"
	} else {
		// 密码正确, 返回生成的Token
		var token string
		if p.IsSave {
			token, err = jwt.GenToken(user.UserID, user.UserKey[:6], 7)
		} else {
			token, err = jwt.GenToken(user.UserID, user.UserKey[:6], 1)
		}
		if err != nil {
			zap.L().Error("An error occurred in token generation, err:", zap.Error(err))
			return res.CodeServerBusy, "服务繁忙"
		}

		// 记录登录IP
		go dao.UpdateUserLoginIP(user.UserID, RemoteIP)
		// 删除缓存中的验证码
		go gcache.DeleteCache(p.Email + "_login")

		return res.CodeSuccess, token
	}
}

// UserFindPwd 找回密码 - 发送Token
func UserFindPwd(p *model.UserFindPwd) (res.ResCode, string) {
	// 检查验证码是否正确
	if !CaptVerify(p.Id, p.Capt) {
		return res.CodeRePwdError, "验证码错误"
	}

	result, user := dao.EmailGetUserData(p.Email)
	if result == false {
		return res.CodeRePwdError, "用户不存在"
	}

	// 缓存查询验证码
	_, err := gcache.GetCache(p.Email + "_fKey")
	if err != nil {
		// 生成验证码
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(999999)
		value, _ := dao.GetSetting("web_title")
		data := "【" + value.Value + "】提示您。您的账户正在找回密码，请输入验证码：" + strconv.Itoa(num) + "。完成验证【验证码有效期：60分钟】"
		// 发送邮件
		if err := email.SendEmailCode("LDDP - 找回密码验证码", user.Email, data); err != nil {
			zap.L().Error("邮件发送失败，err:", zap.Error(err))
			return res.CodeAbnormalError, "邮件发送失败，请稍等片刻再尝试"
		}

		// 发送成功，数据写入缓存
		var f model.SpecialUsersCache
		f.UserId = user.UserID
		f.Code = strconv.Itoa(num)
		v, err := json.Marshal(f)
		if err != nil {
			return res.CodeRePwdError, "验证码发送失败，请稍后尝试"
		}
		go gcache.TimingCache(p.Email+"_fKey", string(v), time.Hour)
		return res.CodeSuccess, "验证码发送成功"
	}
	return res.CodeAbnormalError, "存在未过期验证码，请勿重复发送"
}

// UserRePwd 找回密码 - 修改密码
func UserRePwd(p *model.UserRePwd) (res.ResCode, string) {
	// 缓存查询Token
	code, err := gcache.GetCache(p.Email + "_fKey")
	if err != nil {
		return res.CodeRePwdError, "验证码已失效或不存在"
	}

	// 序列化内容
	var f model.SpecialUsersCache
	err = json.Unmarshal([]byte(code.(string)), &f)
	if err != nil {
		zap.L().Error("[找回密码 - 修改密码]错误，原因：" + err.Error())
		return res.CodeServerBusy, "服务繁忙"
	}
	// 判断验证码
	if p.Code != f.Code {
		return res.CodeRePwdError, "验证码错误"
	}

	// 修改密码
	err = dao.UpdateUserPwd(f.UserId, encryption.Sha1(p.UserKey))
	if err != nil {
		zap.L().Error("[修改密码]错误, 原因：" + err.Error())
		return res.CodeRePwdError, "修改密码失败，请稍后重试"
	}

	// 删除Redis中的Token
	go gcache.DeleteCache(p.Email + "_fKey")

	return res.CodeSuccess, "修改密码成功"
}

// CheckIf 检查是否属于异地登录
func CheckIf(ip1, ip2 string) (bool, bool) {
	/*
		ip1：原始IP
		ip2：登录IP
		bool1：是否异地登录
		bool2：是否解析出错
	*/
	// IP地址相同，跳过验证
	if ip1 == ip2 {
		return true, false
	}

	// 查询IP地址
	url1 := "https://ip.useragentinfo.com/json?ip=" + ip1
	addr1, err := requests.Requests("GET", url1, "", "")
	if err != nil {
		return false, true
	}

	var l1 location
	var l2 location
	// 数据绑定
	err = json.Unmarshal(addr1, &l1)
	if err != nil {
		zap.L().Error(err.Error())
		return false, true
	}

	url2 := "https://ip.useragentinfo.com/json?ip=" + ip2
	addr2, err := requests.Requests("GET", url2, "", "")
	if err != nil {
		return false, true
	}
	// 数据绑定
	err = json.Unmarshal(addr2, &l2)
	if err != nil {
		zap.L().Error(err.Error())
		return false, true
	}

	if l1.City == l2.City {
		return true, false
	} else {
		return false, false
	}
}

// AdminPanelData 网站数据
func AdminPanelData() (res.ResCode, model.AdminWebData) {
	var data model.AdminWebData

	// 获取今天的时间
	e := time.Now().Format("2006-01-02")
	e7 := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	zap.L().Debug("[今日时间]开始时间" + e + " 00:00:00 & 结束时间：" + e + " 23:59:59")
	zap.L().Debug("[七日时间]开始时间" + e7 + " 00:00:00 & 结束时间：" + e + " 23:59:59")

	// 今日用户消费点券总数
	data.ToDayUserConsume = dao.GetConsumptionToDayUpload(e+" 00:00:00", e+" 23:59:59")
	// 今日订单总数
	data.ToDayOrder = dao.GetOrderCountData(e+" 00:00:00", e+" 23:59:59")
	// 站点总订单数
	data.AllOrder = dao.GetOrderDataPage()
	// 站点总用户数
	data.AllUser = dao.GetUserRecord()
	// 今日充值总数
	data.ToDayRecharge = dao.GetRechargeToDayUpload(e)
	// 今日新增用户数
	data.ToDayRegisterUser = dao.GetToDayRegisterUser(e)
	// 活跃用户数
	data.ActiveUser = dao.GetActiveUser(e7)
	// 有效用户数
	data.EffectiveUser = dao.GetEffectiveUser()
	return res.CodeSuccess, data
}

// GetOrderChart 获取7日图标订单数据
func GetOrderChart() (res.ResCode, *model.OrderChart) {
	var sdo []int64
	var sdd []string

	// 获取图表数据
	for i := 7; i > 0; i-- {
		sdo = append(sdo, dao.GetOrderCountData(time.Now().AddDate(0, 0, -i).Format("2006-01-02")+" 00:00:00", time.Now().AddDate(0, 0, -i).Format("2006-01-02")+" 23:59:59"))
		sdd = append(sdd, time.Now().AddDate(0, 0, -i).Format("2006-01-02"))
	}

	SiteData := &model.OrderChart{
		SevenDaysOrder: sdo,
		SevenDaysDate:  sdd,
	}

	return res.CodeSuccess, SiteData
}

// UserDivisionData 用户分页查询
func UserDivisionData(page string) (res.ResCode, model.UserPageData) {
	var data []model.User
	var pageData model.UserPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetDivisionUserData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetDivisionUserData(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetDivisionUserData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetUserRecord()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pageData.Page = z + 1
	} else {
		pageData.Page = z
	}

	// 删除密码
	for i := 0; i < len(data); i++ {
		data[i].UserKey = "safety protection"
	}

	pageData.PageData = data

	return res.CodeSuccess, pageData
}

// UserSearch 用户ID模糊查询
func UserSearch(tp, s string) (res.ResCode, []model.User) {
	var data []model.User
	if tp == "用户UID" {
		data = dao.UserIDSearch(s)
	} else {
		data = dao.UserEmailSearch(s)
	}

	// 删除密码
	for i := 0; i < len(data); i++ {
		data[i].UserKey = "safety protection"
	}
	return res.CodeSuccess, data
}

// UserInformationUpdate 用户数据更新
func UserInformationUpdate(p *model.UpdateUserData) res.ResCode {
	dao.UserInformationUpdate(p)
	return res.CodeSuccess
}

// LoginDivisionData 用户登录日志分页查询
func LoginDivisionData(page string) (res.ResCode, model.LoginPageData) {
	var data []model.LoginIpData
	var pageData model.LoginPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.GetDivisionUserLoginData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.GetDivisionUserLoginData(1)
		} else {
			// 查询指定页数的数据
			data = dao.GetDivisionUserLoginData(intPage)
		}
	}

	// 查询总页数
	count := dao.GetUserLoginRecord()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pageData.Page = z + 1
	} else {
		pageData.Page = z
	}

	pageData.PageData = data

	return res.CodeSuccess, pageData
}

// UserLoginDivisionData 用户登录日志分页查询
func UserLoginDivisionData(uid any, page string) (res.ResCode, model.LoginPageData) {
	var data []model.LoginIpData
	var pageData model.LoginPageData

	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = dao.UserGetDivisionUserLoginData(uid, 1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = dao.UserGetDivisionUserLoginData(uid, 1)
		} else {
			// 查询指定页数的数据
			data = dao.UserGetDivisionUserLoginData(uid, intPage)
		}
	}

	// 查询总页数
	count := dao.UserGetUserLoginRecord(uid)
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		pageData.Page = z + 1
	} else {
		pageData.Page = z
	}

	pageData.PageData = data

	return res.CodeSuccess, pageData
}

// GetUserOneData 用户信息：获取
func GetUserOneData(uid any) (res.ResCode, *model.UserData) {
	var sdo []int64
	var sdd []string

	data := dao.GetUserIDData(uid)
	// 获取今日订单数量
	t := time.Now().Format("2006-01-02")
	// 获取昨日订单数量
	y := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	// 获取图表数据
	for i := 7; i > 0; i-- {
		sdo = append(sdo, dao.GetUserOrderCountData(uid, time.Now().AddDate(0, 0, -i).Format("2006-01-02")+" 00:00:00", time.Now().AddDate(0, 0, -i).Format("2006-01-02")+" 23:59:59"))
		sdd = append(sdd, time.Now().AddDate(0, 0, -i).Format("2006-01-02"))
	}
	// 用户角色
	role := ""
	if data.Role == "proxy" {
		role = "user"
	} else if data.Role == "user" {
		role = "user"
	} else {
		role = "admin"
	}
	ud := &model.UserData{
		UserID:         data.UserID,
		Email:          data.Email,
		Points:         data.Points,
		ToDayOrder:     dao.GetUserOrderCountData(uid, t+" 00:00:00", t+" 23:59:59"),
		YesDayOrder:    dao.GetUserOrderCountData(uid, y+" 00:00:00", y+" 23:59:59"),
		SevenDaysOrder: sdo,
		SevenDaysDate:  sdd,
		Role:           role,
	}

	return res.CodeSuccess, ud
}

// UserLogout 退出登录
func UserLogout() res.ResCode {
	return res.CodeSuccess
}
