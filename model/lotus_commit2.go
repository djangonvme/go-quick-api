package model

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/go-quick-api/types"
)

type LotusCommit2Task struct {
	ID            uint64                   `gorm:"column:id;AUTO_INCREMENT;not null"`
	MinerID       string                   `gorm:"unique_index:idx_sidmidc1;column:miner_id;type:varchar(128);not null;default:''"` // 矿工号id, f01234
	SectorID      abi.SectorNumber         `gorm:"unique_index:idx_sidmidc1;column:sector_id;type:int(11);not null;default:0"`      // 扇区id
	State         types.TaskState          `gorm:"column:state;type:varchar(32);not null;default:''"`                               // 状态: waiting=等待 doing=执行中 finished=已完成 dropped=丢弃
	Commit1OutMd5 string                   `gorm:"unique_index:idx_sidmidc1;column:commit1_out_md5;type:char(32);not null"`         // c1 md5
	Commit1Out    string                   `gorm:"column:commit1_out;type:longtext;default:null"`                                   // c1结果: base64UrlEncodeBytesToString
	Commit2Proof  string                   `gorm:"column:commit2_proof;type:text;default:null"`                                     // c2证明结果
	Memo          string                   `gorm:"column:memo;type:varchar(255);default:null"`
	Workers       []LotusCommit2TaskWorker `gorm:"foreignkey:task_id;association_foreignkey:id" `
	Model
}

type LotusCommit2TaskWorker struct {
	ID           uint64                `gorm:"column:id;AUTO_INCREMENT;not null"`
	TaskID       uint64                `gorm:"index:idx_taskid;column:task_id;type:bigint(20) unsigned;not null"`
	Worker       string                `gorm:"column:worker;type:varchar(128);not null;default:''"`               // 机器标识hostname,ip
	State        types.TaskWorkerState `gorm:"column:state;type:varchar(32);not null;default:''"`                 // finished,failed,reverted
	StartTime    int64                 `gorm:"index:idx_stime;column:start_time;type:int(11);not null;default:0"` // 执行开始时间戳
	EndTime      int64                 `gorm:"column:end_time;type:int(11);not null;default:0"`                   // 执行结束时间戳
	Commit2Proof string                `gorm:"column:commit2_proof;type:text;default:null"`                       // c2结果输出
	Memo         string                `gorm:"column:memo;type:varchar(255);default:null"`                        // memo
	Task         LotusCommit2Task      `gorm:"foreignkey:id;association_foreignkey:task_id"`
	Model
}
