package model

import "time"

type Model struct {
	ID        int64      `gorm:"primary_key column:id" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
