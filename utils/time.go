package utils

import (
	"time"
)

const YMDHIS = "2006-01-02 15:04:05" //常规类型

func DateNow() string {
	return time.Now().Format(YMDHIS)
}

func IsExpired(t time.Time) bool {
	t1 := t.Unix()
	t2 := time.Now().Unix()
	if t1 < t2 {
		return true
	}
	return false
}
