package model

type LotusCommitTaskWorker struct {
	Model
	TaskId       int64           `gorm:"task_id"` // 扇区id
	Worker       string          `gorm:"worker"`  // 任务执行的机器ip或hostname
	Memo         string          `gorm:"memo"`
	State        string          `gorm:"state"`
	Commit2Proof string          `gorm:"commit2_proof"`
	StartTime    int64           `gorm:"start_name"` // 执行开始时间戳
	EndTime      int64           `gorm:"end_time"`   // 执行结束时间戳
	Task         LotusCommitTask `gorm:"foreignkey:id;association_foreignkey:task_id" `

	// CreatedAt         string // 创建时间
	// UpdatedAt         string // 更新时间
	// DeletedAt         string // 删除时间
}
