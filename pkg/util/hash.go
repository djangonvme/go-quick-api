package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}


// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	enPwd := h.Sum(nil)
	// 16进制转换为字符串
	return fmt.Sprintf("%x", enPwd)
}

func ToJson(v interface{}) string {
	if sv, ok := v.(string); ok {
		return sv
	}
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

