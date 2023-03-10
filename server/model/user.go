package model

import (
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	/*
		gorm.Model：基础结构（ID、CreatedAt、UpdatedAt、DeletedAt）
	*/
	gorm.Model
	UserID  string `json:"user_id,omitempty"`  // 用户ID
	Email   string `json:"email,omitempty"`    // 用户邮箱
	UserKey string `json:"user_key,omitempty"` // 用户Key
	Points  int64  `json:"points"`             // 用户点券
	Role    string `json:"role"`               // 用户权限（user：用户、admin：管理员、proxy：代理）
	IsState bool   `json:"is_state"`           // 用户状态（true：启用、false：封禁）
	LoginIP string `json:"login_ip"`           // 上次登录IP
}

// UserData 用户数据
type UserData struct {
	UserID         string   `json:"user_id"`          // 用户ID
	Email          string   `json:"email"`            // 用户邮箱
	Points         int64    `json:"points"`           // 用户积分
	ToDayOrder     int64    `json:"to_day_order"`     // 今日订单数量
	YesDayOrder    int64    `json:"yes_day_order"`    // 昨日订单数量
	SevenDaysOrder []int64  `json:"seven_days_order"` // 七日订单数量
	SevenDaysDate  []string `json:"seven_days_date"`  // 七日日期数组
	Role           string   `json:"role"`             // 用户角色
}

// UserSignUp 用户注册
type UserSignUp struct {
	Email   string `json:"email" binding:"required"`
	UserKey string `json:"user_key" binding:"required"`
	Capt    string `json:"capt" binding:"required"`
	Id      string `json:"id" binding:"required"`
}

// UserSignIn 用户登录
type UserSignIn struct {
	Email   string `json:"email" binding:"required"`
	UserKey string `json:"user_key"  binding:"required"`
	IsSave  bool   `json:"is_save"`
	Capt    string `json:"capt" binding:"required"`
	Id      string `json:"id" binding:"required"`
}

// UserAbnormalEmail 登录异常 - 发送验证码
type UserAbnormalEmail struct {
	Email string `json:"email" binding:"required"`
}

// UserAbnormalSignin 登录异常 - 登录
type UserAbnormalSignin struct {
	Email   string `json:"email" binding:"required"`
	UserKey string `json:"user_key" binding:"required"`
	VfCode  string `json:"vf_code" binding:"required"`
	IsSave  bool   `json:"is_save"`
}

// UserFindPwd 用户找回密码 - 发送验证码
type UserFindPwd struct {
	Email string `json:"email" binding:"required"`
	Capt  string `json:"capt" binding:"required"`
	Id    string `json:"id" binding:"required"`
}

// UserRePwd 用户找回密码 - 修改密码
type UserRePwd struct {
	Email   string `json:"email" binding:"required"`
	UserKey string `json:"user_key" binding:"required"`
	Code    string `json:"code" binding:"required"`
}

// SpecialUsersCache 特殊用户缓存
type SpecialUsersCache struct {
	UserId string
	Code   string
}

// AdminWebData 管理员首页数据
type AdminWebData struct {
	// 今日用户消费点券总数
	ToDayUserConsume int64 `json:"to_day_user_consume"`
	// 今日订单总数
	ToDayOrder int64 `json:"to_day_order"`
	// 站点总订单数
	AllOrder int64 `json:"all_order"`
	// 站点总用户数
	AllUser int64 `json:"all_user"`
	// 今日充值总数
	ToDayRecharge int64 `json:"to_day_recharge"`
	// 今日新增用户数
	ToDayRegisterUser int64 `json:"to_day_register_user"`
	// 活跃用户数
	ActiveUser int64 `json:"active_user"`
	// 有效用户数
	EffectiveUser int64 `json:"effective_user"`
}

// UserPageData 用户分页数据
type UserPageData struct {
	Page     int64  `json:"page"`
	PageData []User `json:"page_data"`
}

// UpdateUserData 用户数据修改
type UpdateUserData struct {
	ID      int64  `json:"id" binding:"required"`    // 用户ID
	Email   string `json:"email" binding:"required"` // 用户邮箱
	Points  int64  `json:"points"`                   // 用户点券
	Role    string `json:"role"`                     // 角色身份
	IsState bool   `json:"is_state"`                 // 用户状态（true：启用、false：封禁）
}

// DeleteUserData 用户数据删除
type DeleteUserData struct {
	ID int `json:"id"  binding:"required"`
}
