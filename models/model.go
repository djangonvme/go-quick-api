package models

import "time"

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *time.Time
}
