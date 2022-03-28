package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/task-dispatcher/erron"
	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/service"
	"gitlab.com/task-dispatcher/types"
	"strconv"
)

func TaskCreate(c *gin.Context) (data interface{}, err error) {
	raw, err := app.GetSaveRawData(c)
	if err != nil {
		return
	}
	if len(raw) == 0 {
		return nil, errors.Errorf("raw Bytes length is 0")
	}
	taskId, err := getHandler(c).Create(string(raw))
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"task_id": taskId,
	}, nil
}

func TaskResult(c *gin.Context) (data interface{}, err error) {
	value := c.Query("task_id")
	value2, _ := strconv.Atoi(value)
	taskId := uint64(value2)
	if taskId <= 0 {
		return nil, errors.Errorf("invalid task_id")
	}
	return getHandler(c).Result(taskId)
}

type ApplyParams struct {
	WorkerName string `json:"worker_name"`
}

type ApplyRes struct {
	ApplyId     uint64 `json:"apply_id"`
	TaskInput   string `json:"task_input"`
	Checksum    string `json:"checksum"`
	Description string `json:"description"`
}

func TaskApply(c *gin.Context) (data interface{}, err error) {

	var param = ApplyParams{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	workerName := param.WorkerName
	if workerName == "" {
		return nil, errors.Errorf("invalid worker_name")
	}

	var output = ApplyRes{}
	workerLogId, checksum, input, err := getHandler(c).Apply(workerName)
	if err != nil {
		if errors.Is(err, erron.ErrNoTaskAvailable) {
			output.Description = err.Error()
			err = nil
			return output, nil
		}
		return nil, err
	}
	output.TaskInput = input
	output.ApplyId = workerLogId
	output.Checksum = checksum

	return output, nil
}

type SubmitParams struct {
	ApplyId  uint64                `json:"apply_id"`
	Checksum string                `json:"checksum"`
	State    types.TaskWorkerState `json:"state"`
	Output   string                `json:"output"`
	ErrMsg   string                `json:"err_msg"`
}

func TaskSubmit(c *gin.Context) (data interface{}, err error) {
	var param = SubmitParams{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	ok, err := getHandler(c).Submit(param.ApplyId, param.State, param.Output, param.Checksum, param.ErrMsg)
	if err != nil {
		return
	}
	if !ok {
		return nil, errors.Errorf("submit faield by expected err")
	}
	return map[string]string{"submitted": "success"}, nil
}

func getHandler(c *gin.Context) types.TaskManager {
	hd, _ := service.GetTaskHandler(c.GetHeader(types.TaskTypeKey))
	return hd
}
