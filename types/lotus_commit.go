package types

/**

create table if not exists `lotus_c2_task`  (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `sector_id` int not null default 0 comment '扇区id',
    `miner_id` varchar(128) not null default '' comment '矿工号id',
    `c2in` MEDIUMTEXT not null comment 'c2输入参数',
    `c2out_proof` MEDIUMTEXT comment 'c2结果输出',
    `state` varchar(16) not null  default '' comment '状态: waiting=等待 doing=执行中 finished=已完成 failed=失败 dropped=丢弃',
    `run_count` int not null default 0  comment '已执行次数',
    `worker_ip` varchar(64) not null default '' comment '任务执行的机器ip或hostname',
    `start_time` int not null default 0 COMMENT '执行开始时间戳',
    `end_time` int not null default 0 COMMENT '执行结束时间戳',
    `sector_created_time` int not null default  0 COMMENT '扇区创建时间戳',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key idx_sc_time (`sector_created_time`),
    key idx_mid_sid (`miner_id`, `sector_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='lotus-c2 task表';

*/

type PieceInfo struct {
	Size     uint64 // Size in nodes. For BLS12-381 (capacity 254 bits), must be >= 16. (16 * 8 = 128)
	PieceCID string
}

/*type C2AddTaskParams struct {
	SectorID          uint64      `json:"sector_id" binding:"required"`
	MinerID           string      `json:"miner_id" binding:"required"`
	Ticket            string      `json:"ticket" binding:"required"`
	Seed              string      `json:"seed" binding:"required"`
	SectorCreatedTime int64       `json:"sector_created_time" binding:"required"`
	ProofType         int64       `json:"proof_type" binding:"required"`
	Pieces            []PieceInfo `json:"pieces" binding:"required"`
}
*/

type CreateTaskParam struct {
}
