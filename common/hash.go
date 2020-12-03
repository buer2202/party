package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func HashSha256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HashSha1(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HashMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
