create database task_dispatcher;
use task_dispatcher;


drop table lotus_commit_task;
create table if not exists `lotus_commit_task`  (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `miner_id` int not null default 0 comment '矿工号id,int',
    `sector_id` int not null default 0 comment '扇区id',
    `seed` char(255) not null comment 'seed value',
    `state` varchar(16) not null  default '' comment '状态: waiting=等待 doing=执行中 finished=已完成 dropped=丢弃',
    `commit1_input` text comment 'c1入参数: base64UrlEncodeBytesToString',
    `commit2_proof` text comment 'c2结果输出: Hex string',
    `memo` varchar(128) comment 'memo',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique key idx_mid_sid_sd (`miner_id`, `sector_id`, `seed`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='lotus_commit_task表';

drop table lotus_commit_task_worker;
create table if not exists `lotus_commit_task_worker`  (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `task_id` bigint(20) unsigned NOT NULL comment '',
    `worker` varchar(128) not null default '' comment '机器标识hostname|ip',
    `state` varchar(16) not null default '' comment 'finished,failed,reverted',
    `start_time` int not null default 0 COMMENT '执行开始时间戳',
    `end_time` int not null default 0 COMMENT '执行结束时间戳',
    `commit2_proof` text comment 'c2结果输出: Hex string',
    `memo` varchar(128) comment 'memo',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key idx_tskid (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='lotus_commit_task_worker表';
