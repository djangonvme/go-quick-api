package utils

import "strconv"

func Int64ToInt(id64 int64) int {
	strInt64 := strconv.FormatInt(id64, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

func StrToInt64(s string) int64 {
	// string到int64
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
