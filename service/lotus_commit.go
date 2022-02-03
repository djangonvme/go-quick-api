package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/jinzhu/gorm"
	"gitlab.com/task-dispatcher/erron"
	"gitlab.com/task-dispatcher/model"
	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/pkg/util"
	"gitlab.com/task-dispatcher/types"
	"strings"
	"time"
)

type LotusCommitTaskHandler struct{}

func NewLotusCommitTaskHandler() *LotusCommitTaskHandler {
	return &LotusCommitTaskHandler{}
}

func (t *LotusCommitTaskHandler) Create(input string) (taskId int64, err error) {
	param := types.Commit2InputParam{}
	if err = json.Unmarshal([]byte(input), &param); err != nil {
		err = fmt.Errorf("json.Unmarshal inputBytes failed: %v", err.Error())
		return
	}
	if err = t.checkCommitInputParam(param); err != nil {
		return
	}
	commit1OutBytes, err := base64.URLEncoding.DecodeString(param.Commit1Out)
	if err != nil {
		return 0, fmt.Errorf("DecodeString Commit1Out: %v", err.Error())
	}
	commitOutMd5 := util.MD5(commit1OutBytes)
	existTask, err := t.getTask(param.MinerId, param.SectorId, commitOutMd5)
	if err != nil {
		return
	}
	if existTask != nil && existTask.ID > 0 {
		return existTask.ID, nil
	}
	insert := model.LotusCommit2Task{
		MinerId:       param.MinerId,
		SectorId:      param.SectorId,
		State:         types.TaskStateWaiting,
		Commit1Out:    param.Commit1Out,
		Commit1OutMd5: commitOutMd5,
	}
	err = app.Db().Model(model.LotusCommit2Task{}).Create(&insert).Error
	if err != nil {
		app.Log().Errorf("insert task failed: %v  data: %+v", err.Error(), insert)
		return
	}
	return insert.ID, nil
}

func (t *LotusCommitTaskHandler) Result(taskId int64) (result map[string]interface{}, err error) {
	task := &model.LotusCommit2Task{}
	err = app.Db().Model(&model.LotusCommit2Task{}).Where("id=?", taskId).First(task).Error
	if err != nil {
		return
	}
	return map[string]interface{}{
		"sector_id":     task.SectorId,
		"miner_id":      task.MinerId,
		"commit2_proof": task.Commit2Proof,
		"state":         task.State,
	}, nil
}

func (t *LotusCommitTaskHandler) Apply(workerName string) (workerLogId int64, checksum string, c2param string, err error) {
	var doingNum int
	err = app.Db().Model(&model.LotusCommit2TaskWorker{}).Where("worker=?", workerName).Where("state=?", types.TaskWorkerStateDoing).Count(&doingNum).Error
	if err != nil {
		return
	}
	if doingNum > 0 {
		err = fmt.Errorf("worker %v has doing tasks, please apply after that comeplete", workerName)
		return
	}

	getWaitingTask := func() ([]int64, error) {
		max := 500
		var taskIds []int64
		err = app.Db().Model(&model.LotusCommit2Task{}).Where("state =?", types.TaskStateWaiting).
			Order("id desc").
			Offset(0).Limit(max).Pluck("id", &taskIds).Error
		if err != nil && gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		if err != nil {
			return nil, err
		}
		return taskIds, nil
	}
	apply := func(workerName string, taskId int64) (wid int64, c2param *types.Commit2InputParam, err error) {
		task, err := t.taskDetail(taskId)
		if err != nil {
			return
		}
		if task.Commit1Out == "" {
			err = fmt.Errorf("apply Commit1Input is empty")
			return
		}
		err = app.Db().Transaction(func(tx *gorm.DB) error {
			insertWorker := &model.LotusCommit2TaskWorker{
				TaskId:    taskId,
				StartTime: time.Now().Unix(),
				Worker:    workerName,
				State:     types.TaskWorkerStateDoing,
			}
			err = tx.Model(&model.LotusCommit2TaskWorker{}).Create(insertWorker).Error
			if err != nil {
				return err
			}
			wid = insertWorker.ID
			if wid == 0 {
				return fmt.Errorf("apply failed: unexpect worker insert id is 0, task: %d worker: %v", taskId, workerName)
			}
			update := tx.Model(&model.LotusCommit2Task{}).Where("id=? and state =? ", taskId, types.TaskStateWaiting).Updates(map[string]interface{}{
				"state": types.TaskStateDoing,
			})
			if update.Error != nil {
				return update.Error
			}
			if update.RowsAffected == 0 {
				return fmt.Errorf("apply failed: no rows changed, task: %d worker: %v", taskId, workerName)
			}
			return nil
		})

		c2param = &types.Commit2InputParam{
			SectorId:   task.SectorId,
			MinerId:    task.MinerId,
			Commit1Out: task.Commit1Out,
		}
		return wid, c2param, err
	}
	// apply end
	taskIds, err := getWaitingTask()
	if err != nil {
		return
	}
	var applyTaskId int64
	var c2paramInfo *types.Commit2InputParam
	num := len(taskIds)
	retry := 0
	for i := num; i > 0; i-- {
		if retry > 20 {
			app.Log().Errorf("apply error retry too more times")
			break
		}
		applyTaskId = taskIds[util.RandNum(int64(num))]
		ok, err := app.Locker().Lock(t.lockKey(applyTaskId), "", time.Second*30) // 1 min for updating state=doing
		if err != nil {
			app.Log().Errorf("apply lock task %d faield: %v", applyTaskId, err.Error())
		} else if !ok {
			retry++
			continue
		}
		workerLogId, c2paramInfo, err = apply(workerName, applyTaskId)
		if err != nil {
			app.Log().Errorf("apply task %d for %s faield: %s try: %d", applyTaskId, workerName, err.Error(), i)
			retry++
			continue
		} else {
			break
		}
	}
	// get no task
	if workerLogId == 0 {
		err = erron.ErrNoTaskAvailable
		return
	}
	if c2paramInfo == nil {
		err = fmt.Errorf("unexpect err: apply gen wokerLogId(%d) success, but find commit1Input is empty", workerLogId)
		return
	}
	app.Log().Infof("LotusCommitTaskHandler apply success! workerName: %v workerLogId: %v taskId: %v c2param: %v", workerName, workerLogId, applyTaskId, c2param)
	c2paramBytes, err := json.Marshal(c2paramInfo)
	checksum = genChecksum(c2paramInfo.MinerId, c2paramInfo.SectorId, c2paramInfo.Commit1Out)
	return workerLogId, checksum, string(c2paramBytes), err

}

func (t *LotusCommitTaskHandler) Submit(workerLogId int64, state types.TaskWorkerState, commit2Proof string, checksum string, errMsg string) (bool, error) {
	if state != types.TaskWorkerStateFailed && state != types.TaskWorkerStateFinished {
		return false, fmt.Errorf("submitted invalid worker state")
	}
	updateWorker := func(db *gorm.DB) error {
		update := map[string]interface{}{
			"state":    state,
			"end_time": time.Now().Unix(),
		}
		if state == types.TaskWorkerStateFinished {
			update["commit2_proof"] = commit2Proof
		}
		if errMsg != "" {
			update["memo"] = errMsg
		}
		return db.Model(&model.LotusCommit2TaskWorker{}).Where("id=?", workerLogId).Updates(update).Error
	}

	workerLog, err := t.taskWorkerDetail(workerLogId)
	if err != nil {
		return false, fmt.Errorf("fetch woker log by wokerLogId(%d) failed: %v", workerLogId, err)
	}
	if workerLog.ID == 0 {
		return false, fmt.Errorf("fetch woker log by wokerLogId(%d) failed: %v", workerLogId, "record not found!")
	}
	if workerLog.State == types.TaskWorkerStateReverted {
		return false, fmt.Errorf("task has been reverted, submit reject! wid: %v", workerLogId)
	}

	taskDetail := workerLog.Task
	if taskDetail.ID == 0 {
		return false, fmt.Errorf("submit tast failed: task not exists by wokerLogId: %d", workerLogId)
	}
	if (taskDetail.State == types.TaskStateFinished && taskDetail.Commit2Proof != "") || taskDetail.State == types.TaskStateDropped {
		if err = updateWorker(app.Db().DB); err != nil {
			app.Log().Errorf("LotusCommitTaskHandler Submit updateWorker failed: %v", err.Error())
		}
		return true, fmt.Errorf("submit task failed: task areadly finished! taskId: %d workerLogId: %d", taskDetail.ID, workerLogId)
	}

	genSum := genChecksum(taskDetail.MinerId, taskDetail.SectorId, taskDetail.Commit1Out)

	if genSum != checksum {
		return false, fmt.Errorf("submit task failed: worker/task info not matched by checksum! workerLogId: %d gensum: %v checksum: %v", workerLogId, genSum, checksum)
	}
	if state == types.TaskWorkerStateFinished && commit2Proof == "" {
		return false, fmt.Errorf("submit task with finished state must pass finished ouput result(c2outProof)! wokerLogId: %d", workerLogId)
	}
	updateTask := func(tx *gorm.DB) error {
		update := map[string]interface{}{}
		if state == types.TaskWorkerStateFinished {
			update["commit2_proof"] = commit2Proof
			update["state"] = types.TaskStateFinished
		} else {
			update["state"] = types.TaskStateWaiting
		}
		return tx.Model(&model.LotusCommit2Task{}).Where("id=?", taskDetail.ID).Updates(update).Error
	}
	err = app.Db().Transaction(func(tx *gorm.DB) error {
		if err = updateWorker(tx); err != nil {
			return err
		}
		if err = updateTask(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		app.Log().Errorf("LotusCommitTaskHandler submit failed: %v taskId: %d workerLogId: %d", err.Error(), taskDetail.ID, workerLogId)
		return false, err
	}
	app.Log().Infof("LotusCommitTaskHandler submit success! taskId: %d workerLogId: %d state: %v  commit2Proof: %v", taskDetail.ID, workerLogId, state, commit2Proof)
	return true, nil
}

func (t *LotusCommitTaskHandler) Revert() {
	taskTimeoutSeconds := 3600 + 1200
	// timeTicker := time.Tick(5 * time.Minute)
	timeTicker := time.Tick(60 * time.Second)
	doRevert := func() error {
		var timeoutWorkers []model.LotusCommit2TaskWorker
		err := app.Db().Model(&model.LotusCommit2TaskWorker{}).Preload("Task").
			Where("state =?", types.TaskWorkerStateDoing).
			Where("end_time=0 and ? - start_time > ?", time.Now().Unix(), taskTimeoutSeconds).
			Find(&timeoutWorkers).Error

		if err != nil {
			return err
		}
		for _, item := range timeoutWorkers {
			if item.Task.State == types.TaskStateDoing {
				err = app.Db().Transaction(func(tx *gorm.DB) error {
					err = tx.Model(&model.LotusCommit2Task{}).Where("id=?", item.Task.ID).Updates(map[string]interface{}{
						"state": types.TaskStateWaiting,
						"memo":  "reverted state to " + types.TaskStateWaiting,
					}).Error
					if err != nil {
						return err
					}
					err = tx.Model(&model.LotusCommit2TaskWorker{}).Where("task_id=?", item.Task.ID).Updates(map[string]interface{}{
						"state":      types.TaskWorkerStateReverted,
						"memo":       "reverted state to " + types.TaskWorkerStateReverted,
						"end_time":   time.Now().Unix(),
						"deleted_at": time.Now(),
					}).Error
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					app.Log().Errorf("revert task state failed: %v task_id: %v", err, item.Task.ID)
				}
			}
		}
		return nil
	}

	for {
		select {
		case <-timeTicker:
			err := doRevert()
			if err != nil {
				app.Log().Errorf("revert failed: %v", err.Error())
			}
		}
	}
}

func (t *LotusCommitTaskHandler) getTask(minerId string, sectorId abi.SectorNumber, c1outMd5 string) (task *model.LotusCommit2Task, err error) {
	task = &model.LotusCommit2Task{}
	err = app.Db().Model(model.LotusCommit2Task{}).
		Where("miner_id=? and sector_id=? and commit1_out_md5=?", minerId, sectorId, c1outMd5).
		First(task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return
	}
	return
}
func (t *LotusCommitTaskHandler) lockKey(taskId int64) string {
	return fmt.Sprintf("taskDispatcher.LotusCommitTaskApply.lockTask_%d", taskId)
}

func (t *LotusCommitTaskHandler) taskDetail(id int64) (data *model.LotusCommit2Task, err error) {
	data = &model.LotusCommit2Task{}
	err = app.Db().Model(&model.LotusCommit2Task{}).Preload("Workers").Where("id=?", id).First(data).Error
	return
}
func (t *LotusCommitTaskHandler) taskWorkerDetail(id int64) (data *model.LotusCommit2TaskWorker, err error) {
	data = &model.LotusCommit2TaskWorker{}
	err = app.Db().Model(&model.LotusCommit2TaskWorker{}).Preload("Task").Where("id=?", id).First(data).Error
	return
}

func (t *LotusCommitTaskHandler) checkCommitInputParam(param types.Commit2InputParam) error {
	var errMsgs []string
	if param.SectorId < 0 {
		errMsgs = append(errMsgs, "sectorId is null")
	}
	if param.MinerId == "" {
		errMsgs = append(errMsgs, "minerId is null")
	}
	if param.Commit1Out == "" || len(param.Commit1Out) == 0 {
		errMsgs = append(errMsgs, "Commit1Out is null")
	}
	if len(errMsgs) > 0 {
		return fmt.Errorf("check CommitInputParam failed: %v", strings.Join(errMsgs, ";"))
	}
	return nil
}

func genChecksum(minerId string, sectorId abi.SectorNumber, commit1out string) string {
	return util.MD5String(fmt.Sprintf("%v%v%v", minerId, sectorId, commit1out))
}
