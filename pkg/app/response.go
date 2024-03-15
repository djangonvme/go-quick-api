package app

import (
	"encoding/json"
	"time"

	"gitlab.com/qubic-pool/erron"
)

type Response struct {
	Code      erron.ErrCode `json:"code"`
	Msg       string        `json:"msg"`
	Timestamp int64         `json:"timestamp"`
	Data      interface{}   `json:"data"`
}

func (r *Response) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ResponseWithPager struct {
	Pager Pager       `json:"pager"`
	List  interface{} `json:"list"`
}

func ResponseFail(err erron.ErrorIF) *Response {
	return ResponseApi(err, struct{}{})
}

func ResponseFailByCode(code erron.ErrCode) *Response {
	return ResponseApi(erron.New(code, ""), struct{}{})
}

func ResponseApi(err erron.ErrorIF, data interface{}) *Response {
	if err == nil {
		err = erron.New(erron.Success, "")
	}
	errMsg := err.Error()
	if errMsg == "" {
		errMsg = err.Code().Text()
	}
	return &Response{
		err.Code(),
		errMsg,
		time.Now().Unix(),
		data,
	}
}

func PagerResponse(pager Pager, list interface{}) *ResponseWithPager {
	return &ResponseWithPager{
		Pager: pager,
		List:  list,
	}
}
