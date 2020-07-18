package util

import "strconv"

type StringNumber string

func (s *StringNumber) Int64() int64 {
	str := string(*s)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func (s *StringNumber) Int() int {
	str := string(*s)
	i, _ := strconv.Atoi(str)
	return i
}

func (s *StringNumber) Uint() uint {
	return uint(s.Int())
}

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

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func UniqueSliceInt64(src []int64) []int64 {
	res := make([]int64, 0)
	for _, sv := range src {
		exists := false
		for _, rv := range res {
			if rv == sv {
				exists = true
				break
			}
		}
		if exists == false {
			res = append(res, sv)
		}
	}
	return res
}

func UniqueSliceString(src []string) []string {
	res := make([]string, 0)
	for _, sv := range src {
		exists := false
		for _, rv := range res {
			if rv == sv {
				exists = true
				break
			}
		}
		if exists == false {
			res = append(res, sv)
		}
	}
	return res
}

func InStringSlice(key string, src []string) bool {
	for _, v := range src {
		if v == key {
			return true
		}
	}
	return false
}

func MapStringKeys(m map[string]interface{}) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
