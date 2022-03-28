package types

type TaskManager interface {
	Create(input string) (taskId uint64, err error)
	Result(taskId uint64) (result map[string]interface{}, err error)

	Apply(workerName string) (applyId uint64, checksum string, inputParam string, err error)
	Submit(applyId uint64, state TaskWorkerState, outputResult string, checksum string, errMsg string) (bool, error)

	Revert()
}
