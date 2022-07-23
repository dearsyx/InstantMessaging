package util

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateMD5 生成md5值
func GenerateMD5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// 比较md5值
