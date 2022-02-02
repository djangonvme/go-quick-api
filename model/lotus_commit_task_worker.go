package model

import "gitlab.com/task-dispatcher/types"

type LotusCommit2TaskWorker struct {
	Model
	TaskId       int64                 `gorm:"task_id"` // 扇区id
	Worker       string                `gorm:"worker"`  // 任务执行的机器ip或hostname
	Memo         string                `gorm:"memo"`
	State        types.TaskWorkerState `gorm:"state"`
	Commit2Proof string                `gorm:"commit2_proof"`
	StartTime    int64                 `gorm:"start_name"` // 执行开始时间戳
	EndTime      int64                 `gorm:"end_time"`   // 执行结束时间戳
	Task         LotusCommit2Task      `gorm:"foreignkey:id;association_foreignkey:task_id" `

	// CreatedAt         string // 创建时间
	// UpdatedAt         string // 更新时间
	// DeletedAt         string // 删除时间
}
