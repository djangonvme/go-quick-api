package helper

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
	"math/rand"
)

func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func GetRandNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}

func GetRandString() string  {
	n:= GetRandNum()
	return strconv.Itoa(n)
}

func GetMd5RandString() string {
	return Md5(GetRandString())
}

