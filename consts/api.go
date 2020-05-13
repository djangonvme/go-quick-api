package consts

import (
	"errors"
	"fmt"
)

const ApiCodeSuccess = 200 // success code
const ApiCodeError = 400   // general fail code
const ApiCodeTokenValid = 401
const ApiCodeTimeout = 408

var ApiCodeMaps = map[int]string{
	ApiCodeError:      "请求失败",
	ApiCodeSuccess:    "请求成功",
	ApiCodeTokenValid: "请求失败,TOKEN失效请先登陆",
	ApiCodeTimeout:    "请求超时, 程序执行中断",
}

func GetApiMsgByCode(code int) (msg string, err error) {
	if v, ok := ApiCodeMaps[code]; ok {
		msg = v
	} else {
		msg = fmt.Sprintf("Unknown api code [%d] in ApiCodeMaps ", code)
		err = errors.New(msg)
	}
	return
}
