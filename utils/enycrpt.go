package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Sha256(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	enPwd := h.Sum(nil)
	//16进制转换为字符串
	return fmt.Sprintf("%x", enPwd)
}

/*
func Pwd(pwd string)  string {
	return Sha256(pwd)
}*/

func ToJson(v interface{}) string {
	j, err := json.Marshal(v)
	if err != nil {
		log.Fatal("json Marshal err")
		return ""
	}
	return string(j)
}

func RandNum(length int) []byte {
	letters := []byte("0123456789")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rd.Intn(len(letters))]
	}
	return b
}

func RandToken() string {
	return Sha256(string(RandNum(100)))
}
