package helper

import (
	"time"
)

const DBTIME_FORMAT  = "2006-01-02 03:04:05"

func GetExpiredAt(afterSeconds int) time.Time {
	timestamp := time.Now().Unix() + IntToInt64(afterSeconds)
	tm2 := time.Unix(timestamp, 0)
	date := tm2.Format(DBTIME_FORMAT)
	time,_:= time.Parse(DBTIME_FORMAT, date)
	return time
}

func IsExpired(t time.Time) bool {
	t1 := t.Unix()
	t2 := time.Now().Unix()
	if t1 < t2 {
		return true
	}
	return false
}
