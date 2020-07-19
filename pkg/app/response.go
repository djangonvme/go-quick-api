package app

import (
	"time"

	"github.com/jangozw/go-quick-api/pkg/util"

	"github.com/jangozw/go-quick-api/erron"
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

// 返回错误， 带错误信息
func ResponseFail(err erron.E) *response {
	return Response(err, struct{}{})
}

func ResponseFailByCode(code erron.Code) *response {
	return Response(erron.New(code), struct{}{})
}

// err 为nil 即为成功
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

// 带分页的输出结果 response.data 的结构
func PagerResponse(pager Pager, list interface{}) *responseWithPager {
	return &responseWithPager{
		Pager: pager,
		List:  list,
	}
}
