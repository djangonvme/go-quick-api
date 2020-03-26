#创建数据库
create database if not exists ginapicommon default character set utf8;

-- ------------------------------- user and grant -------------------------------------------
# 创建用户和授权MYSQL8.0官方语法有改动
#创建账户
create user 'appadmin'@'%' identified by  '123456';
#赋予权限，with grant option这个选项表示该用户可以将自己拥有的权限授权给别人
grant all privileges on ginapicommon.* to 'appadmin'@'%' with grant option;

#改密码&授权超用户，flush privileges 命令本质上的作用是将当前user和privilige表中的用户信息/权限设置从mysql库(MySQL数据库的内置库)中提取到内存里
create user 'app'@'%' identified by  '123456';
grant SELECT,INSERT on *.* to 'app'@'%' with grant option;
flush privileges;

-- ------------------------------- user and grant -------------------------------------------



# 此目录可不提交
# 用到的sql只是示例用法
# 本例中用到的表
use ginapicommon;
create table if not exists `user`  (
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
# pwd is 123456
INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`name`,`mobile`,`password`,`status`) VALUES (1577091432,1577091432,NULL,'hello','18000000000','c2cb76e4f39bd2d5c78395c7df7f94b1fa84a78097e7ec3fa905bcdfff699029',0);
# user token
create table if not exists `user_token` (
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

