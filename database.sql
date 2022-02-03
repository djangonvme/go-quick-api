create database task_dispatcher;
use task_dispatcher;


drop table lotus_commit2_task;
create table if not exists `lotus_commit2_task`  (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `miner_id` varchar(128) not null default "" comment '矿工号id, f01234',
    `sector_id` int not null default 0 comment '扇区id',
    `state` varchar(32) not null  default '' comment '状态: waiting=等待 doing=执行中 finished=已完成 dropped=丢弃',
    `commit1_out_md5` char(32) not null comment 'c1 md5',
    `commit1_out` LONGTEXT comment 'c1结果base64',
    `commit2_proof` text comment 'c2证明结果',
    `memo` varchar(255) comment 'memo',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique key idx_sidmidc1 (`sector_id`, `miner_id`, `commit1_out_md5`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='lotus_commit2_task表';

drop table lotus_commit2_task_worker;
create table if not exists `lotus_commit2_task_worker`  (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `task_id` bigint(20) unsigned NOT NULL comment '',
    `worker` varchar(128) not null default '' comment '机器标识hostname,ip',
    `state` varchar(32) not null default '' comment 'finished,failed,reverted',
    `start_time` int not null default 0 COMMENT '执行开始时间戳',
    `end_time` int not null default 0 COMMENT '执行结束时间戳',
    `commit2_proof` text comment 'c2结果输出',
    `memo` varchar(255) comment 'memo',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key idx_taskid (`task_id`),
    key idx_stime (`start_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='lotus_commit2_task_worker表';
