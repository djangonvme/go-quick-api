package service

import (
	"fmt"
	"gitlab.com/task-dispatcher/types"
)

func GetTaskHandler(taskType string) (types.TaskManager, error) {
	switch taskType {
	case types.TaskTypeLotusCommit:
		return NewLotusCommitTaskHandler(), nil
	case types.TaskTypeLotusWPost:
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("invalid taskType, couldn't match handler")
	}
}
