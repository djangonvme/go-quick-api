package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	YMDHIS = "2006-01-02 15:04:05" // 常规类型
	YMD    = "2006-01-02"          // 常规类型
)

func DateNow() string {
	return time.Now().Format(YMDHIS)
}

func IsExpired(t time.Time) bool {
	t1 := t.Unix()
	t2 := time.Now().Unix()
	return t1 < t2
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间和最后一秒时间
func GetDayStartAndEnd(d time.Time) (st time.Time, et time.Time) {
	st = GetZeroTime(d)
	et = GetLastSecondTime(d)
	return
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一天的最后一秒时间
func GetLastSecondTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

func UnixToStr(un int64, layout string) string {
	return time.Unix(un, 0).Format(layout)
}

func UnixToYMDHIS(un int64) string {
	return time.Unix(un, 0).Format(YMDHIS)
}

func Now() string {
	return time.Unix(time.Now().Unix(), 0).Format(YMDHIS)
}

func UnixToYMD(un int64) string {
	return time.Unix(un, 0).Format("2006-01-02")
}

func UnixToYMDInt(un int64) int {
	a := time.Unix(un, 0).Format("20060102")
	return StrToInt(a)
}

func UnixToHi(un int64) string {
	return time.Unix(un, 0).Format("15:04")
}

// 相差多少个自然日
func TimeDiffDays(start int64, end int64) int {
	return int(GetDayBegin(end)+86400-GetDayBegin(start)) / 86400
}

// 注意 timeStr 时分秒都得带上，否则报错
func DateToUnix(timeStr string) int64 {
	reg := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})\s?(\d{2})?:?(\d{2})?:?(\d{2})?`)
	if reg.MatchString(timeStr) {
		ma := reg.FindStringSubmatch(timeStr)
		var year, month, day, hour, min, sec string
		year = ma[1]
		month = ma[2]
		day = ma[3]
		if ma[4] != "" {
			hour = ma[4]
		} else {
			hour = "00"
		}
		if ma[5] != "" {
			min = ma[5]
		} else {
			min = "00"
		}
		if ma[6] != "" {
			sec = ma[6]
		} else {
			sec = "00"
		}
		timeStr = fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, day, hour, min, sec)
	} else {
		return 0
	}
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return t.Unix()
}

// 今天开始时间
func GetTodayBegin() int64 {
	return GetDayBegin(time.Now().Unix())
}

// 当天开始时间  判断用 >= dayBegin
func GetDayBegin(unixTime int64) int64 {
	timeStr := time.Unix(unixTime, 0).Format("2006-01-02")
	// 使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}

// 当天最后结束时间 判断 要用 < dayEnd
func GetDayEnd(unixTime int64) int64 {
	return GetDayBegin(unixTime) + 86400
}

func GetDateYmdByString(str string, split string) string {
	var pattern string
	if split == "" {
		pattern = `(\d{4})(\d{2})(\d{2})`
	} else if split == "/" {
		pattern = `(\d{4})/(\d{2})/(\d{2})`
	} else {
		pattern = `(\d{4})-(\d{2})-(\d{2})`
	}
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(str) {
		matches := reg.FindStringSubmatch(str)
		return fmt.Sprintf("%s-%s-%s", matches[1], matches[2], matches[3])
	}
	return ""
}

type ParseWorkTime struct {
	DayStartHour   int
	DayStartMinute int
	DayEndHour     int
	DayEndMinute   int
	ToadySignIn    int64
	ToadySignOut   int64
}

// 解析24小时内的上下班时间字符串 : from=08:00 ,to=17:30
func ParseWorkTimeByString(from string, to string) (workTime ParseWorkTime, err error) {
	err = errors.New("from must less than to , check your func args")
	if from == "" || to == "" {
		return
	}
	workTime = ParseWorkTime{
		DayStartHour:   0,
		DayStartMinute: 0,
		DayEndHour:     0,
		DayEndMinute:   0,
		ToadySignIn:    0,
		ToadySignOut:   0,
	}
	reg := regexp.MustCompile(`(\d{2}):(\d{2})`)
	if reg.MatchString(from) {
		ma := reg.FindStringSubmatch(from)
		workTime.DayStartHour = StrToInt(ma[1])
		workTime.DayStartMinute = StrToInt(ma[2])
	} else {
		return
	}
	if reg.MatchString(to) {
		ma := reg.FindStringSubmatch(to)
		workTime.DayEndMinute = StrToInt(ma[2])
		h := StrToInt(ma[1])
		// 结束时间早于开始时间，则转24小时制
		if h*3600+workTime.DayEndMinute*60 < workTime.DayStartHour*3600+workTime.DayStartMinute*60 {
			workTime.DayEndHour = 12 + h
			return
		} else {
			workTime.DayEndHour = h
		}
	} else {
		return
	}
	if !(workTime.DayStartHour < 24 && workTime.DayStartHour >= 0) || !(workTime.DayStartMinute >= 0 && workTime.DayStartMinute < 60) || !(workTime.DayEndHour < 24 && workTime.DayEndHour >= 0) || !(workTime.DayEndMinute >= 0 && workTime.DayEndMinute < 60) {
		return
	}
	todayBegin := GetTodayBegin()
	workTime.ToadySignIn = todayBegin + int64(workTime.DayStartHour)*3600 + int64(workTime.DayStartMinute)*60
	workTime.ToadySignOut = todayBegin + int64(workTime.DayEndHour)*3600 + int64(workTime.DayEndMinute)*60
	return workTime, nil
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
