package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// S 字符串类型转换
type S string

func (s S) String() string {
	return string(s)
}

// Bytes 转换为[]byte
func (s S) Bytes() []byte {
	return []byte(s)
}

// Bool 转换为bool
func (s S) Bool() (bool, error) {
	b, err := strconv.ParseBool(s.String())
	if err != nil {
		return false, err
	}
	return b, nil
}

// DefaultBool 转换为bool，如果出现错误则使用默认值
func (s S) DefaultBool(defaultVal bool) bool {
	b, err := s.Bool()
	if err != nil {
		return defaultVal
	}
	return b
}

// Int64 转换为int64
func (s S) Int64() (int64, error) {
	i, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultInt64 转换为int64，如果出现错误则使用默认值
func (s S) DefaultInt64(defaultVal int64) int64 {
	i, err := s.Int64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Int 转换为int
func (s S) Int() (int, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// DefaultInt 转换为int，如果出现错误则使用默认值
func (s S) DefaultInt(defaultVal int) int {
	i, err := s.Int()
	if err != nil {
		return defaultVal
	}
	return i
}

// Uint64 转换为uint64
func (s S) Uint64() (uint64, error) {
	i, err := strconv.ParseUint(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultUint64 转换为uint64，如果出现错误则使用默认值
func (s S) DefaultUint64(defaultVal uint64) uint64 {
	i, err := s.Uint64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Uint 转换为uint
func (s S) Uint() (uint, error) {
	i, err := s.Uint64()
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

// DefaultUint 转换为uint，如果出现错误则使用默认值
func (s S) DefaultUint(defaultVal uint) uint {
	i, err := s.Uint()
	if err != nil {
		return defaultVal
	}
	return uint(i)
}

// Float64 转换为float64
func (s S) Float64() (float64, error) {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// DefaultFloat64 转换为float64，如果出现错误则使用默认值
func (s S) DefaultFloat64(defaultVal float64) float64 {
	f, err := s.Float64()
	if err != nil {
		return defaultVal
	}
	return f
}

// Float32 转换为float32
func (s S) Float32() (float32, error) {
	f, err := s.Float64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// DefaultFloat32 转换为float32，如果出现错误则使用默认值
func (s S) DefaultFloat32(defaultVal float32) float32 {
	f, err := s.Float32()
	if err != nil {
		return defaultVal
	}
	return f
}

// ToJSON 转换为JSON

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				// fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

type Location struct {
	Country      string
	Province     string
	City         string
	Area         string
	ProvinceCode int
	CityCode     int
	AreaCode     int
}

func (l *Location) MakeLocation() string {
	return fmt.Sprintf("%s,%s,%s,%s", l.Country, l.Province, l.City, l.Area)
}

// 解析 location 里有省市区地址： "中国,北京市,北京市,东城区" "中国,浙江省,杭州市,萧山区"
func ParseLocation(addr string) (Location, error) {
	addr = strings.TrimSpace(addr)
	loc := Location{}
	matches := regexp.MustCompile(`中国,(\S+),(\S+),(\S+)`).FindStringSubmatch(addr)

	if len(matches) != 4 {
		return loc, errors.New(addr + ": 腾讯地图定位解析省市区失败!")
	}
	loc.Province = matches[1]
	loc.City = matches[2]
	loc.Area = matches[3]
	return loc, nil
}
