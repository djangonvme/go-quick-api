package models

import "time"

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *time.Time
}
/*
create table `user` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,

	`created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
	`updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
	`deleted_at` datetime default null COMMENT '删除时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='xx表';

*/