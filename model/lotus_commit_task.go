package model

import (
	"github.com/filecoin-project/go-state-types/abi"
	"gitlab.com/task-dispatcher/types"
)

type LotusCommitTask struct {
	Model
	// ID                int64  `gorm:"primary_key column:id"`
	SectorId     abi.SectorNumber `gorm:"sector_id"` // 扇区id
	MinerId      abi.ActorID      `gorm:"miner_id"`  // 矿工号id
	Seed         string           `gorm:"seed"`
	State        types.TaskState  `gorm:"state"` // 状态: waiting=等待 doing=执行中 finished=已完成 failed=失败 drop=丢弃
	Memo         string           `gorm:"memo"`
	Commit1Input string           `gorm:"commit1_input"`
	Commit2Proof string           `gorm:"commit2_proof"`
	//Workers      []LotusCommitTaskWorker `gorm:"foreignkey:id;association_foreignkey:task_id"`

	Workers []LotusCommitTaskWorker `gorm:"foreignkey:task_id;association_foreignkey:id" `
	// Workers []LotusCommitTaskWorker `gorm:"foreignkey:id;association_foreignkey:task_id" `

	// CreatedAt         string // 创建时间
	// UpdatedAt         string // 更新时间
	// DeletedAt         string // 删除时间
}
