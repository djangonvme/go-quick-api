package service

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/jinzhu/gorm"
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
	inputBytes, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		err = fmt.Errorf("DecodeString(input) failed: %v", err.Error())
		return
	}
	param := types.CommitInputParam{}
	if err = json.Unmarshal(inputBytes, &param); err != nil {
		err = fmt.Errorf("json.Unmarshal inputBytes failed: %v", err.Error())
		return
	}
	if err = t.checkCommitInputParam(param); err != nil {
		return
	}
	seedHex := hex.EncodeToString(param.Seed)
	existTask, err := t.getTask(param.Sector.ID.Miner.String(), param.Sector.ID.Number, seedHex)
	if err != nil {
		return
	}
	if existTask != nil && existTask.ID > 0 {
		return existTask.ID, nil
	}
	insert := model.LotusCommitTask{
		MinerId:      param.Sector.ID.Miner,
		SectorId:     param.Sector.ID.Number,
		Seed:         seedHex,
		State:        types.TaskStateWaiting,
		Commit1Input: input,
	}
	err = app.Db().Model(model.LotusCommitTask{}).Create(&insert).Error
	if err != nil {
		app.Log().Errorf("insert task failed: %v  data: %+v", err.Error(), insert)
		return
	}
	return insert.ID, nil
}

func (t *LotusCommitTaskHandler) Status(taskId int64) (result map[string]interface{}, err error) {
	task := &model.LotusCommitTask{}
	err = app.Db().Model(&model.LotusCommitTask{}).Where("id=?", taskId).First(task).Error
	if err != nil {
		return
	}
	return map[string]interface{}{
		"id":            task.ID,
		"sector_id":     task.SectorId,
		"miner_id":      task.MinerId,
		"seed":          task.Seed,
		"commit2_proof": task.Commit2Proof,
		"state":         task.State,
	}, nil
}

func (t *LotusCommitTaskHandler) Apply(workerName string) (workerLogId int64, commit1Input string, err error) {
	getWaitingTask := func() ([]int64, error) {
		max := 500
		var taskIds []int64
		err = app.Db().Model(&model.LotusCommitTask{}).Where("state =?", types.TaskStateWaiting).
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
	apply := func(workerName string, taskId int64) (wid int64, input string, err error) {
		task, err := t.taskDetail(taskId)
		if err != nil {
			return
		}
		if task.Commit1Input == "" {
			err = fmt.Errorf("apply Commit1Input is empty")
			return
		}
		err = app.Db().Transaction(func(tx *gorm.DB) error {
			insertWorker := &model.LotusCommitTaskWorker{
				TaskId:    taskId,
				StartTime: time.Now().Unix(),
				Worker:    workerName,
			}
			err = tx.Model(&model.LotusCommitTaskWorker{}).Create(insertWorker).Error
			if err != nil {
				return err
			}
			wid = insertWorker.ID
			if wid == 0 {
				return fmt.Errorf("apply failed: unexpect worker insert id is 0, task: %d worker: %v", taskId, workerName)
			}
			update := tx.Model(&model.LotusCommitTask{}).Where("id=? and state =? ", taskId, types.TaskStateWaiting).Updates(map[string]interface{}{
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
		return wid, task.Commit1Input, err
	}

	taskIds, err := getWaitingTask()
	if err != nil {
		return
	}
	var applyTaskId int64
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
		workerLogId, commit1Input, err = apply(workerName, applyTaskId)
		if err != nil {
			app.Log().Errorf("apply task %d for %s faield: %s try: %d", applyTaskId, workerName, err.Error(), i)
			retry++
			continue
		} else {
			break
		}
	}
	if workerLogId == 0 {
		return 0, "", fmt.Errorf("no task availiable, please retry later! (%v)", err.Error())
	}
	if commit1Input == "" {
		return 0, "", fmt.Errorf("unexpect err: apply gen wokerLogId(%d) success, but find commit1Input is empty", workerLogId)
	}

	app.Log().Infof("LotusCommitTaskHandler apply success! workerName: %v workerLogId: %v taskId: %v commit1Input: %v", workerName, workerLogId, applyTaskId, commit1Input)
	return workerLogId, commit1Input, nil
}

func (t *LotusCommitTaskHandler) Submit(workerLogId int64, state types.TaskWorkerState, commit2Proof string, workerName string, errMsg string) (bool, error) {
	if state != types.TaskWorkerStateFailed && state != types.TaskWorkerStateFinished {
		return false, fmt.Errorf("submitted invalid worker state")
	}
	updateWorker := func(tx *gorm.DB) error {
		return tx.Model(&model.LotusCommitTaskWorker{}).Where("id=?", workerLogId).Updates(map[string]interface{}{
			"state":         state,
			"end_time":      time.Now().Unix(),
			"commit2_proof": commit2Proof,
			"memo":          errMsg,
		}).Error
	}

	workerLog, err := t.taskWorkerDetail(workerLogId)
	if err != nil {
		return false, fmt.Errorf("find woker log by wokerLogId(%d) failed: %v", workerLogId, err)
	}
	taskDetail := workerLog.Task
	if taskDetail.ID == 0 {
		return false, fmt.Errorf("submit tast failed: task not exists by wokerLogId: %d", workerLogId)
	}
	if taskDetail.State == types.TaskStateFinished && taskDetail.Commit2Proof != "" {
		if err = updateWorker(app.Db().DB); err != nil {
			app.Log().Errorf("LotusCommitTaskHandler Submit updateWorker failed: %v", err.Error())
		}
		return true, fmt.Errorf("submit task failed: task areadly finished! taskId: %d workerLogId: %d", taskDetail.ID, workerLogId)
	}
	if workerLog.Worker != workerName {
		return false, fmt.Errorf("submit task failed: worker not matched! workerLogId: %d", workerLogId)
	}
	if state == types.TaskWorkerStateFinished && commit2Proof == "" {
		return false, fmt.Errorf("submit task with finished state must pass finished ouput result(c2outProof)! wokerLogId: %d", workerLogId)
	}
	updateTask := func(tx *gorm.DB) error {
		u := map[string]interface{}{"state": state}
		if state == types.TaskWorkerStateFinished {
			u["commit2_proof"] = commit2Proof
			u["state"] = types.TaskStateFinished
		} else {
			u["state"] = types.TaskStateWaiting
		}
		return tx.Model(&model.LotusCommitTask{}).Where("id=?", taskDetail.ID).Updates(u).Error
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
	app.Log().Infof("LotusCommitTaskHandler submit success! taskId: %d workerLogId: %d workerName: %v  state: %v  commit2Proof: %v", taskDetail.ID, workerLogId, workerName, state, commit2Proof)
	return true, nil
}

func (t *LotusCommitTaskHandler) Revert() {
	taskTimeoutSeconds := 3600 + 1200
	// taskTimeoutSeconds := 3
	timeTicker := time.Tick(5 * time.Minute)
	//timeTicker := time.Tick(3 * time.Second)

	getDoingList := func() (list []model.LotusCommitTask, err error) {
		err = app.Db().Model(&model.LotusCommitTask{}).
			Preload("Workers", fmt.Sprintf("(end_time=0 AND %d - start_time > %d)", time.Now().Unix(), taskTimeoutSeconds)).
			Where("state =?", types.TaskStateDoing).
			Order("id asc").
			Offset(0).
			Limit(500).
			Find(&list).Error
		if err != nil && gorm.IsRecordNotFoundError(err) {
			err = nil
		}
		return
	}
	revertTask := func(taskId int64) error {
		return app.Db().Model(&model.LotusCommitTask{}).
			Where("id=?", taskId).
			Updates(map[string]interface{}{
				"state": types.TaskStateWaiting,
			}).Error
	}
	revertWorker := func(workerLogId int64) error {
		return app.Db().Model(&model.LotusCommitTaskWorker{}).
			Where("id=?", workerLogId).
			Updates(map[string]interface{}{
				"state":    types.TaskWorkerStateReverted,
				"end_time": time.Now().Unix(),
				"memo":     "system reverted",
			}).Error
	}
	revert := func(list []model.LotusCommitTask) {
		app.Log().Infof("revert start, total: %d", len(list))
		for _, task := range list {
			if err := revertTask(task.ID); err != nil {
				app.Log().Errorf("revetTask failed: %v", err.Error())
			} else {
				app.Log().Infof("revert task done! task_id: %d", task.ID)
			}
			for _, wk := range task.Workers {
				if err := revertWorker(wk.ID); err != nil {
					app.Log().Errorf("revertWorker failed: %v ", err.Error())
				} else {
					app.Log().Infof("revert task worker log done! workerLogId: %d", wk.ID)
				}
			}
			continue
		}
	}
	for {
		select {
		case <-timeTicker:
			app.Log().Infof("try revert start")
			list, err := getDoingList()
			if err != nil {
				app.Log().Errorf("getDoingList failed: %v", err.Error())
				continue
			}
			if len(list) == 0 {
				app.Log().Infof("checked no revert list")
				continue
			}
			revert(list)
		}
	}
}

func (t *LotusCommitTaskHandler) getTask(minerId string, sectorId abi.SectorNumber, seed string) (task *model.LotusCommitTask, err error) {
	task = &model.LotusCommitTask{}
	err = app.Db().Model(model.LotusCommitTask{}).
		Where("miner_id=? and sector_id=? and seed=?", minerId, sectorId, seed).
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

func (t *LotusCommitTaskHandler) taskDetail(id int64) (data *model.LotusCommitTask, err error) {
	data = &model.LotusCommitTask{}
	err = app.Db().Model(&model.LotusCommitTask{}).Preload("Workers").Where("id=?", id).First(data).Error
	return
}
func (t *LotusCommitTaskHandler) taskWorkerDetail(id int64) (data *model.LotusCommitTaskWorker, err error) {
	data = &model.LotusCommitTaskWorker{}
	err = app.Db().Model(&model.LotusCommitTaskWorker{}).Preload("Task").Where("id=?", id).First(data).Error
	return
}

func (t *LotusCommitTaskHandler) checkCommitInputParam(param types.CommitInputParam) error {
	var errMsgs []string
	if param.Sector.ID.Miner <= 0 {
		errMsgs = append(errMsgs, "miner is null")
	}
	if param.Sector.ID.Number < 0 {
		errMsgs = append(errMsgs, "sector number is null")
	}
	if param.Sector.ProofType < 0 {
		errMsgs = append(errMsgs, "sector ProofType is null")
	}
	if len(param.Ticket) == 0 {
		errMsgs = append(errMsgs, "ticket is null")
	}
	if len(param.Seed) == 0 {
		errMsgs = append(errMsgs, "seed is null")
	}
	if len(param.Pieces) == 0 || len(param.Pieces[0].PieceCID.Bytes()) == 0 || param.Pieces[0].Size <= 0 {
		errMsgs = append(errMsgs, "pieces is null")
	}
	if len(param.Cids.Sealed.Bytes()) == 0 {
		errMsgs = append(errMsgs, "cids.Sealed is null")
	}
	if len(param.Cids.Unsealed.Bytes()) == 0 {
		errMsgs = append(errMsgs, "cids.Unsealed is null")
	}
	if len(errMsgs) > 0 {
		return fmt.Errorf("check CommitInputParam failed: %v", strings.Join(errMsgs, ";"))
	}
	return nil
}
