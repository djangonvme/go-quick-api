package service

import (
	"github.com/pkg/errors"
	"gitlab.com/task-dispatcher/types"
)

func GetTaskHandler(taskType string) (types.TaskManager, error) {
	switch taskType {
	case types.TaskTypeLotusCommit:
		return NewLotusCommitTaskHandler(), nil
	case types.TaskTypeLotusWPost:
		return nil, errors.Errorf("not impl")
	default:
		return nil, errors.Errorf("invalid taskType, couldn't match handler")
	}
}
