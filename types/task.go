package types

type TaskManager interface {
	Create(input string) (taskId int64, err error)
	Result(taskId int64) (result map[string]interface{}, err error)

	Apply(workerName string) (applyId int64, checksum string, inputParam string, err error)
	Submit(applyId int64, state TaskWorkerState, outputResult string, checksum string, errMsg string) (bool, error)

	Revert()
}
