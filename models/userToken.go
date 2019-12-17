package models

// UserToken 用户token表
type UserToken struct {
	Model
	UserID    int64  `gorm:"index"` // 用户id
	Token     string // token
	ExpiredAt int64  // 过期时间
}
