package service

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/types"
	"testing"
	"time"
)

type SectorCids struct {
	Unsealed cid.Cid
	Sealed   cid.Cid
}

func getCid(va string) cid.Cid {
	h, err := mh.Sum([]byte(va), mh.SHA3, 4)
	if err != nil {
		panic(err)
	}
	cidRes := cid.NewCidV1(128, h)
	return cidRes
}

func TestAddTask(t *testing.T) {
	sCids := SectorCids{
		Sealed:   getCid("xxx"),
		Unsealed: getCid("123"),
	}

	fmt.Println(sCids)

	b, err := json.Marshal(sCids)

	if err != nil {
		fmt.Println("Mar err", err.Error())
		return
	}

	s2 := base64.URLEncoding.EncodeToString(b)
	b2, err := base64.URLEncoding.DecodeString(s2)
	if err != nil {
		fmt.Println("dec err", err)
		return
	}
	var s3 SectorCids

	err = json.Unmarshal(b2, &s3)
	if err != nil {
		fmt.Println("um err", err)
		return
	}

	fmt.Println(s3)

	maddr, err := address.NewFromString("f01231")
	if err != nil {
		fmt.Println("new err", err.Error())
		return
	}
	aid, _ := address.IDFromAddress(maddr)

	sss := abi.ActorID(aid).String()
	fmt.Println(sss)

}

func TestRedsiLogc(t *testing.T) {
	//app.Init(app.LogService, app.DbService, app.RedisService)
	res, err := app.Redis().Lock("test.redis.lock", "abc", 5*time.Second)
	fmt.Println("lock res:", res, "err:", err)

	res, err = app.Redis().Unlock("test.redis.lock", "abc")
	fmt.Println("Unlock res:", res, "err:", err)

	res, err = app.Redis().Lock("test.redis.lock", "abc", 5*time.Second)
	fmt.Println("lock res:", res, "err:", err)

	/*	var wg sync.WaitGroup
		num := 10000
		wg.Add(num)
		for i:=0;i<num;i++{
			cur := i
			go func() {
				defer wg.Done()
				res, err := app.Redis.Lock(fmt.Sprintf("test.redis.lock%d", cur), "abc", 5*time.Second)
				if err != nil || ! res {
					fmt.Println("failed lock res", cur, " res: ", res, "err:", err)
				} else {
					fmt.Println("ok", cur)
				}
			}()
		}

		wg.Wait()*/

}

func TestLotusCommitTaskHandler_Create(t *testing.T) {
	// app.Init(app.LogService, app.DbService, app.RedisService)
	handler := NewLotusCommitTaskHandler()
	c1param := types.CommitInputParam{
		Sector: types.SectorRef{
			ID: abi.SectorID{
				Miner:  abi.ActorID(0),
				Number: abi.SectorNumber(2),
			},
			ProofType: abi.RegisteredSealProof_StackedDrg2KiBV1_1,
		},
		Ticket: abi.SealRandomness([]byte("ticket")),
		Seed:   abi.InteractiveSealRandomness([]byte("seed")),
		Pieces: []abi.PieceInfo{
			{
				Size:     2048,
				PieceCID: getCid("123"),
			},
		},
		Cids: types.SectorCids{
			Sealed:   getCid("sealed"),
			Unsealed: getCid("unsealed"),
		},
	}

	c1b, err := json.Marshal(c1param)
	if err != nil {
		fmt.Println("Marshal err: ", err.Error())
		return
	}
	comm1Input := base64.URLEncoding.EncodeToString(c1b)
	taskId, err := handler.Create(comm1Input)
	fmt.Println("taskId: ", taskId, "err: ", err)
}

func TestLotusCommitTaskHandler_Apply_Submit(t *testing.T) {
	//app.Init(app.LogService, app.DbService, app.RedisService)
	handler := NewLotusCommitTaskHandler()
	workerName := "HK-172.1.1.1"
	doAndSubmit := func(workerLogId int64, input string, workerName string) error {
		param := types.CommitInputParam{}
		inputBytes, err := base64.URLEncoding.DecodeString(input)
		if err != nil {
			return fmt.Errorf("DecodeString input faield: %v", err.Error())
		}
		if err = json.Unmarshal(inputBytes, &param); err != nil {
			return fmt.Errorf("unmarshal input faield: %v", err.Error())
		}
		fmt.Println("do commit start")
		time.Sleep(2 * time.Second)
		fmt.Println("do commit end")
		commit2Proof := hex.EncodeToString([]byte("helloworldzwuv"))
		ok, err := handler.Submit(workerLogId, "finished", commit2Proof, workerName, "")
		if err != nil {
			return fmt.Errorf("submit failed: %v", err.Error())
		}
		if ok {
			return nil
		}
		return fmt.Errorf("submit failed by some reasons unknow")
	}
	count := 0
	for {
		count++
		select {
		case <-time.Tick(2 * time.Second):
			app.Log().Infof("try get task apply.... %v", count)
			wName := workerName + fmt.Sprintf("_%v", count)
			workerId, input, err := handler.Apply(wName)
			if err != nil {
				fmt.Println("apply failed:", err.Error())
				continue
			}
			fmt.Println("apply get task success!", "worker_log_id", workerId, "input: ", input)
			if err := doAndSubmit(workerId, input, wName); err != nil {
				fmt.Println("doAndSubmit failed:", err.Error())
				return
			}
			fmt.Println("doAndSubmit success! workerLogId: ", workerId)
			return
		}

	}
}

func TestLotusCommitTaskHandler_Status(t *testing.T) {
	handler := NewLotusCommitTaskHandler()
	data, err := handler.Status(1)
	fmt.Println("err: ", err)
	fmt.Println(data)
}

func TestLotusCommitTaskHandler_TryRevert(t *testing.T) {
	handler := NewLotusCommitTaskHandler()
	handler.Revert()
	//
	/*	getDoingList := func() (list []model.LotusCommitTask, err error) {
			err = app.Db.Model(&model.LotusCommitTask{}).
				Preload("Workers", fmt.Sprintf("(end_time=0 AND %d - start_time > 1)", time.Now().Unix())).
				Where("state =?", types.TaskStateDoing).
				Order("id asc").
				Offset(0).
				Limit(200).
				Find(&list).Error
			if err != nil && gorm.IsRecordNotFoundError(err) {
				err = nil
			}
			return
		}

		list, err := getDoingList()

		fmt.Println("err", err)

		for _, v := range list {
			for _, v2 := range v.Workers {
				fmt.Println("log id: ", v2.ID, "task_id: ", v2.TaskId, "task_id2: ", v2.Task.ID)
			}
		}*/

}

func TestPrelOad(t *testing.T) {
	handler := NewLotusCommitTaskHandler()
	/*	data, err := handler.taskWorkerDetail(1)
		if err != nil {
			fmt.Println("err: ", err.Error())
			return
		} else {
			fmt.Println("task_id", data.Task.ID)
		}*/

	data2, err2 := handler.taskDetail(1)
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return
	} else {
		fmt.Println(len(data2.Workers), data2.Workers)
	}

}
