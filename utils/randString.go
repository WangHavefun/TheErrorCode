package utils

import (
	"encoding/hex"
	"log"
	"math/rand"
)

var n = 12

func GenerateSubId() string {
	randomBytes := make([]byte, 16) // 16 字节将生成 32 位的随机哈希值
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(randomBytes)
}
