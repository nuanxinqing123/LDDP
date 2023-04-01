package dao

import (
	"LDDP/server/model"
	"time"
)

// GetUserRecord 获取User表记录数量
func GetUserRecord() int64 {
	var user []model.User
	result := DB.Find(&user)
	return result.RowsAffected
}

// EmailGetUserData 邮箱获取用户信息
func EmailGetUserData(email string) (bool, model.User) {
	var user model.User
	// 通过邮箱和用户名查询用户信息
	DB.Where("email = ?", email).First(&user)

	// 判断是否已注册
	if user.Email != "" {
		// 存在
		return true, user
	} else {
		// 不存在
		return false, user
	}
}

// GetUserIDData 用户UID获取用户信息
func GetUserIDData(uid any) model.User {
	var user model.User
	// 通过用户ID查询用户信息
	DB.Where("user_id = ?", uid).First(&user)
	return user
}

// InsertUser 创建新用户
func InsertUser(user *model.User) (err error) {
	err = DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDivisionUserData 条件查询用户数据
func GetDivisionUserData(page int) []model.User {
	var user []model.User
	if page == 1 {
		// 获取第一页数据（25条）
		DB.Order("id desc").Limit(20).Offset(0).Find(&user)
	} else {
		// 获取第N页数据
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&user)
	}
	return user
}

// UserIDSearch 用户ID模糊查询
func UserIDSearch(s string) []model.User {
	var user []model.User
	DB.Where("user_id LIKE ?", "%"+s+"%").Find(&user)
	return user
}

// UserEmailSearch 用户邮箱模糊查询
func UserEmailSearch(s string) []model.User {
	var user []model.User
	DB.Where("email LIKE ?", "%"+s+"%").Find(&user)
	return user
}

// UserInformationUpdate 用户数据更新
func UserInformationUpdate(p *model.UpdateUserData) {
	var user model.User
	DB.Where("id = ?", p.ID).First(&user)
	user.Email = p.Email
	user.Points = p.Points
	user.Role = p.Role
	user.Remarks = p.Remarks
	user.IsState = p.IsState
	DB.Save(&user)
}

// UpdateUserLoginIP 更新用户登录IP地址
func UpdateUserLoginIP(uid any, ip string) {
	var u model.User
	DB.Where("user_id = ?", uid).First(&u)
	u.LoginIP = ip
	DB.Save(&u)
}

// GetUserLoginRecord 获取用户登录日志记录数量
func GetUserLoginRecord() int64 {
	var user []model.LoginIpData
	result := DB.Find(&user)
	return result.RowsAffected
}

// GetDivisionUserLoginData 登录日志分页查询
func GetDivisionUserLoginData(page int) []model.LoginIpData {
	var user []model.LoginIpData
	if page == 1 {
		// 获取第一页数据（25条）
		DB.Order("id desc").Limit(20).Offset(0).Find(&user)
	} else {
		// 获取第N页数据
		DB.Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&user)
	}
	return user
}

// LoginSearch	登录日志模糊查询
func LoginSearch(s string) []model.LoginIpData {
	var user []model.LoginIpData
	e := time.Now().Format("2006-01-02")
	e7 := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	DB.Where("user_id = ? and created_at between ? and ?", s, e7, e).Find(&user)
	return user
}

// UserGetDivisionUserLoginData 用户登录日志分页查询
func UserGetDivisionUserLoginData(uid any, page int) []model.LoginIpData {
	var user []model.LoginIpData
	if page == 1 {
		// 获取第一页数据（25条）
		DB.Where("user_id = ?", uid).Order("id desc").Limit(20).Offset(0).Find(&user)
	} else {
		// 获取第N页数据
		DB.Where("user_id = ?", uid).Order("id desc").Limit(20).Offset((page - 1) * 20).Find(&user)
	}
	return user
}

// UserGetUserLoginRecord 获取用户登录日志记录数量
func UserGetUserLoginRecord(uid any) int64 {
	var user []model.LoginIpData
	result := DB.Where("user_id = ?", uid).Find(&user)
	return result.RowsAffected
}

// UpdateUser 保存用户信息
func UpdateUser(p model.User) {
	DB.Save(&p)
}

// GetToDayRegisterUser 查询今日注册用户数量
func GetToDayRegisterUser(s string) int64 {
	var user []model.User
	result := DB.Where("created_at >= ?", s).Find(&user)
	return result.RowsAffected
}

// GetActiveUser 查询七日内活跃用户数量
func GetActiveUser(s string) int64 {
	var user []model.User
	return DB.Where("updated_at >= ?", s).Find(&user).RowsAffected
}

// GetEffectiveUser 查询用户余额不为0的用户数量
func GetEffectiveUser() int64 {
	var user []model.User
	return DB.Where("points > ?", 0).Find(&user).RowsAffected
}

// UpdateUserPwd 找回密码-修改密码
func UpdateUserPwd(uid string, pwd string) error {
	return DB.Model(&model.User{}).Where("user_id = ?", uid).Update("user_key", pwd).Error
}
