package apiv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/task-dispatcher/service"
	"gitlab.com/task-dispatcher/types"
)

func TaskCreate(c *gin.Context) (data interface{}, err error) {
	raw, err := c.GetRawData()
	if err != nil {
		err = fmt.Errorf("couldn't get RawData")
		return
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("raw Bytes length is 0")
	}
	taskId, err := getHandler(c).Create(string(raw))
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"task_id": taskId,
	}, nil
}

func TaskStatus(c *gin.Context) (data interface{}, err error) {
	taskId := c.GetInt64("task_id")
	if taskId <= 0 {
		return nil, fmt.Errorf("invalid task_id")
	}
	return getHandler(c).Status(taskId)
}

type ApplyParams struct {
	WorkerName string `json:"worker_name"`
}

func TaskApply(c *gin.Context) (data interface{}, err error) {
	var param = ApplyParams{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	workerName := param.WorkerName
	if workerName == "" {
		return nil, fmt.Errorf("invalid worker_name")
	}
	workerLogId, input, err := getHandler(c).Apply(workerName)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"apply_id": workerLogId,
		"input":    input,
	}, nil
}

type SubmitParams struct {
	ApplyId    int64                 `json:"apply_id"`
	WorkerName string                `json:"worker_name"`
	State      types.TaskWorkerState `json:"state"`
	Output     string                `json:"output"`
	ErrMsg     string                `json:"err_msg"`
}

func TaskSubmit(c *gin.Context) (data interface{}, err error) {
	var param = SubmitParams{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	ok, err := getHandler(c).Submit(param.ApplyId, param.State, param.Output, param.WorkerName, param.ErrMsg)
	if err != nil {
		return
	}
	if !ok {
		return nil, fmt.Errorf("submit faield by expected err")
	}
	return "ok", nil
}

func getHandler(c *gin.Context) types.TaskManager {
	hd, _ := service.GetTaskHandler(c.GetHeader(types.TaskTypeKey))
	return hd
}
