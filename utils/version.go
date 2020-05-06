package utils

var Build buildInfo

// 代码构建信息
type buildInfo struct {
	Time    string
	Version string
}

func SetBuildInfo(version string, time string) {
	Build = buildInfo{
		Time:    time,
		Version: version,
	}
}
