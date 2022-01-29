package types

type TaskManager interface {
	Create(input string) (taskId int64, err error)
	Status(taskId int64) (result map[string]interface{}, err error)

	Apply(workerName string) (workerLogId int64, inputParam string, err error)
	Submit(workerLogId int64, state TaskWorkerState, outputResult string, workerName string, errMsg string) (bool, error)

	Revert()
}
