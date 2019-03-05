package helper

import (
	"strconv"
)

func Int64ToInt(id64 int64) int {
	strInt64 := strconv.FormatInt(id64, 10)
	id16,_ := strconv.Atoi(strInt64)
	return id16
}

func IntToInt64(n int) int64 {
	string:=strconv.Itoa(n)
	int64, _ := strconv.ParseInt(string, 10, 64)
	return int64
}

