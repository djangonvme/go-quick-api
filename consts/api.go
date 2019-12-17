package consts

const ApiCodeSuccess = 200
const ApiCodeError = 400
const ApiCodeTokenValid = 401

var ApiCodeMaps = map[int]string{
	ApiCodeError:      "请求失败",
	ApiCodeSuccess:    "请求成功",
	ApiCodeTokenValid: "TOKEN 失效",
}
