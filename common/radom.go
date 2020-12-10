package common

import (
	"math/rand"
	"time"
)

type Range struct {
	Begin int
	End   int
}

// 取区间随机数
func RandIntFromRange(r Range) int {
	rand.Seed(time.Now().UnixNano())
	if r.End-r.Begin <= 0 {
		return r.Begin
	}
	return rand.Intn((r.End-r.Begin)+1) + r.Begin
}

// 获取区间随机整数
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// 生成指定出的随机字符串
func RandSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
