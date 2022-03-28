package service

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
	"github.com/go-quick-api/pkg/util"
	"github.com/go-quick-api/types"
	"io/ioutil"
	"testing"
	"time"
)

var apiUrl = "http://localhost:8180/api/v1"

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

func TestLotusCommitTaskHandler_Create(t *testing.T) {
	c1param := types.Commit2InputParam{
		MinerId:    "f0123",
		SectorId:   1,
		Commit1Out: "d",
	}
	c1b, err := json.Marshal(c1param)
	if err != nil {
		fmt.Println("Marshal err: ", err.Error())
		return
	}
	// comm1Input := base64.URLEncoding.EncodeToString(c1b)

	url := fmt.Sprintf("%s/%v", apiUrl, "task/create")

	var header = map[string]string{"Task-type": "lotus-commit2"}

	res, err := util.HttpPost(url, string(c1b), header)

	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println(string(res.Body))
}

type ApplyResp struct {
	ApplyId int64  `json:"apply_id"`
	Input   string `json:"input"`
}

type StatusParams struct {
	TaskId int64 `json:"task_id"`
}

type ResultResp struct {
	SectorId     int64  `json:"sector_id"`
	MinerId      int64  `json:"miner_id"`
	Commit2Proof string `json:"commit2_proof"`
	State        string `json:"state"`
}

type SubmitParams struct {
	ApplyId    int64                 `json:"apply_id"`
	WorkerName string                `json:"worker_name"`
	State      types.TaskWorkerState `json:"state"`
	Output     string                `json:"output"`
	ErrMsg     string                `json:"err_msg"`
}

type ApiResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func DecodeApiResp(respBytes []byte, data interface{}) error {
	var resp = ApiResp{}
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return err
	}
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(dataBytes, data); err != nil {
		return err
	}
	return nil
}

func TestApply(t *testing.T) {
	var header = map[string]string{"Task-type": "lotus-commit2"}

	apply := func(wName string) (wid int64, input string, err error) {
		url := apiUrl + "/task/apply"
		data := map[string]interface{}{"worker_name": wName}
		res, err := util.HttpPost(url, data, header)
		if err != nil {
			fmt.Println("apply err: ", err.Error())
			return
		}
		fmt.Println("apply res: ", string(res.Body))
		applyResp := ApplyResp{}
		if err = DecodeApiResp(res.Body, &applyResp); err != nil {
			return
		}
		return applyResp.ApplyId, applyResp.Input, nil
	}

	wid, input, err := apply("abc")
	fmt.Println("-------------------------------------")

	fmt.Println("err:  ", err)
	fmt.Println(input)
	fmt.Println(wid)

}

func TestLotusCommitTaskHandler_Apply_Submit(t *testing.T) {
	var header = map[string]string{"Task-type": "lotus-commit2"}
	apply := func(wName string) (wid int64, input string, err error) {
		url := apiUrl + "/task/apply"
		data := map[string]interface{}{"worker_name": wName}
		res, err := util.HttpPost(url, data, header)
		if err != nil {
			fmt.Println("apply err: ", err.Error())
			return
		}
		fmt.Println("apply res: ", string(res.Body))
		applyResp := ApplyResp{}
		if err = DecodeApiResp(res.Body, &applyResp); err != nil {
			return
		}
		return applyResp.ApplyId, applyResp.Input, nil
	}
	submit := func(workerLogId int64, state types.TaskWorkerState, commit2Proof string, workerName string, errMsg string) (bool, error) {
		url := apiUrl + "/task/submit"
		data := SubmitParams{
			ApplyId:    workerLogId,
			State:      state,
			Output:     commit2Proof,
			WorkerName: workerName,
			ErrMsg:     errMsg,
		}
		res, err := util.HttpPost(url, data, header)
		if err != nil {
			fmt.Println("submit err: ", err.Error())
			return false, err
		}
		fmt.Println("submit res: ", string(res.Body))
		if res.StatusOk() {
			return true, nil
		} else {
			return false, errors.Errorf("stauts err ok")
		}
	}

	workerName := "HK-172.1.1.1"
	doAndSubmit := func(workerLogId int64, c2input string, workerName string) error {
		fmt.Println("do commit start")
		time.Sleep(2 * time.Second)
		fmt.Println("do commit end with c2 input: ", c2input)
		commit2Proof := hex.EncodeToString([]byte("helloworldzwuv"))
		ok, err := submit(workerLogId, "finished", commit2Proof, workerName, "")
		if err != nil {
			return errors.Errorf("submit failed: %v", err.Error())
		}
		if ok {
			return nil
		}
		return errors.Errorf("submit failed by some reasons unknow")
	}
	count := 0
	for {
		count++
		select {
		case <-time.Tick(3 * time.Second):
			fmt.Printf("try get task apply.... %v \n", count)
			wName := workerName + fmt.Sprintf("_%v", count)
			workerId, input, err := apply(wName)
			if err != nil {
				fmt.Println("apply failed:", err.Error())
				continue
			}
			if workerId == 0 {
				fmt.Println("apply failed: worker id is 0")
				continue
			}

			fmt.Println("apply get task success!", "worker_log_id", workerId, "input: ", input)
			if err = doAndSubmit(workerId, input, wName); err != nil {
				fmt.Println("doAndSubmit failed:", err.Error())
				continue
			}
			fmt.Println("doAndSubmit success! workerLogId: ", workerId)
			return
		}

	}
}

func TestLotusCommitTaskHandler_Result(t *testing.T) {
	requrl := apiUrl + "/task/result"
	param := map[string]string{
		"task_id": "4",
	}
	headers := map[string]string{"Task-type": "lotus-commit2"}

	resp, err := util.HttpGet(requrl, param, headers)
	if err != nil {
		fmt.Println("HttpGet err: ", err)
		return
	}
	fmt.Println(string(resp.Body))

	if resp.StatusCode != 200 {
		fmt.Println("failed  statusCode !=200")
		return
	}

	var statusRes = ResultResp{}

	if err := DecodeApiResp(resp.Body, &statusRes); err != nil {
		fmt.Println("err decode: ", err.Error())
		return
	}

	fmt.Printf("Success: %+v\n", statusRes)
}

func TestLotusCommitTaskHandler_TryRevert(t *testing.T) {
	handler := NewLotusCommitTaskHandler()
	handler.Revert()
}

func TestLotusCommitTaskHandler_Apply2(t *testing.T) {

	handler := NewLotusCommitTaskHandler()
	wid, _, o, err := handler.Apply("testan")
	fmt.Println("err is: ", err)
	fmt.Println(wid)
	fmt.Println(o)

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

func TestDd(t *testing.T) {

	byes, _ := ioutil.ReadFile("/Users/Django/work/proof.txt")
	fmt.Println("origin total bytes", len(byes))
	b2 := base64.URLEncoding.EncodeToString(byes)
	// fmt.Println(b2)
	fmt.Println("bsae64 len:", len(b2))
}

func TestBY(t *testing.T) {

	msg := fmt.Sprintf("#11 %+v", util.ErrTest)
	fmt.Println(msg)

}
