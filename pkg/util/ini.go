package util

import (
	"encoding/json"
	"regexp"
	"strconv"

	"gopkg.in/ini.v1"
)

// 将ini配置内容解析到指定的结构体中
func ParseIni(file string, receiveObj interface{}) error {
	iniFile, err := ini.Load(file)
	if err != nil {
		return err
	}
	return parseIniToObj(iniFile, receiveObj)
}

func parseIniToObj(iniFile *ini.File, receiveObj interface{}) error {
	contentMap := make(map[string]map[string]string)
	for _, s := range iniFile.Sections() {
		subMap := make(map[string]string)
		for _, s2 := range s.Keys() {
			subMap[s2.Name()] = s2.Value()
			contentMap[s.Name()] = subMap
		}
	}
	by, err := json.Marshal(receiveObj)
	if err != nil {
		return err
	}
	cfgMap := make(map[string]map[string]interface{})
	if err := json.Unmarshal(by, &cfgMap); err != nil {
		return nil
	}
	var mapNumDefault float64
	for k, vMap := range cfgMap {
		for k2, v2 := range vMap {
			if value, ok := getValueFromCfgMap(k, k2, contentMap); ok {
				if v2 == mapNumDefault {
					cfgMap[k][k2] = getInt64FromString(value)
				} else {
					cfgMap[k][k2] = value
				}
			}
		}
	}
	by, err = json.Marshal(cfgMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(by, receiveObj)
}

func getValueFromCfgMap(k1 string, k2 string, m map[string]map[string]string) (string, bool) {
	if v, ok := m[k1]; ok {
		if v, ok := v[k2]; ok {
			return v, true
		}
	}
	return "", false
}

func getInt64FromString(str string) int64 {
	reg, err := regexp.Compile(`(\d+)`)
	if err != nil {
		return 0
	}
	if reg.MatchString(str) {
		match := reg.FindStringSubmatch(str)
		if len(match) >= 1 {
			num, _ := strconv.ParseInt(match[1], 10, 64)
			return num
		}
	}
	return 0
}
