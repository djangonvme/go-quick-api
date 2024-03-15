package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func AutoMigrate(db *gorm.DB) {
	db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("created_at", time.Now())
		scope.SetColumn("updated_at", time.Now())
	})
	db.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("updated_at", time.Now())
	})
	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&TUser{})

}
