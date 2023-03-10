package model

import "gorm.io/gorm"

// JWTAdmin JWT表
type JWTAdmin struct {
	gorm.Model
	SecretKey string // JWT签名密钥
}

// CheckToken 检查Token是否有效
type CheckToken struct {
	JWToken string `json:"token" binding:"required"`
}
