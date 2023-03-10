package model

// WebSettings 网站配置模型
type WebSettings struct {
	Key   string `json:"key" gorm:"primaryKey"`
	Value string `json:"value"`
}
