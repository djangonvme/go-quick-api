package model

import (
	"time"
)

// TUser [...]
type TUser struct {
	ID        int64     `gorm:"column:id;auto increment" `
	MinerID   string    `gorm:"index:miner_id;column:miner_id;type:varchar(100);not null;default:'';comment:'qubic miner addr'" ` // qubic miner addr
	Username  string    `gorm:"index:username;column:username;type:varchar(100);not null;default:'';comment:'username'" `         // username
	Password  string    `gorm:"column:password;type:varchar(100);not null;default:''" `
	Phone     string    `gorm:"column:phone;type:varchar(100);not null;default:''" `
	Email     string    `gorm:"column:email;type:varchar(100);not null;default:''" `
	Token     string    `gorm:"column:token;type:varchar(512);not null;default:''" `
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" `
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP" `
}

// TableName get sql table name.获取数据库表名
func (m *TUser) TableName() string {
	return "t_user"
}
