package model

import (
	"time"
)

type Model struct {

	/*	CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`*/

	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;default:null"`                           // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime;default:null;default:CURRENT_TIMESTAMP"` // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime;default:null"`
}
