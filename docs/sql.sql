# 此目录可不提交
# 用到的sql只是示例用法
# 本例中用到的表

use gindemo;

create table `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
     name varchar(16) not null default '' comment '姓名',
     mobile varchar(20) not null default 0 comment '手机号',
     password varchar(64) not null default 0 comment '密码',
     status tinyint(2) not null  default 0 comment '状态',
    `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
    `deleted_at` datetime default null COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key idx_mobile (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

create table `user_token` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    user_id bigint(20) not null default 0 comment '用户id',
    token varchar(64) not null default 0 comment 'token',
    `expired_at` int(11) NOT NULL DEFAULT 0 COMMENT '过期时间',
    `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
    `deleted_at` datetime default null COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key idx_uid (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户token表';
