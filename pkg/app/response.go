package app

import (
	"time"

	"github.com/go-quick-api/pkg/util"

	"github.com/go-quick-api/erron"
)

// api 接口响应结果
type response struct {
	Code      erron.Code  `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type Pager struct {
	util.Pager
}

// 分页结构数据
type responseWithPager struct {
	Pager Pager       `json:"pager"`
	List  interface{} `json:"list"`
}

func ResponseFail(err erron.E) *response {
	return Response(err, struct{}{})
}

func ResponseFailByCode(code erron.Code) *response {
	return Response(erron.New(code), struct{}{})
}

func Response(err erron.E, data interface{}) *response {
	if err == nil {
		err = erron.New(erron.Success)
	}
	errMsg := err.Msg()
	if errMsg == "" {
		errMsg = err.Code().Msg()
	}
	return &response{
		err.Code(),
		errMsg,
		time.Now().Unix(),
		data,
	}
}

func PagerResponse(pager Pager, list interface{}) *responseWithPager {
	return &responseWithPager{
		Pager: pager,
		List:  list,
	}
}
