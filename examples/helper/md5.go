package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 ...
func MD5(sign string) string {
	h := md5.New()
	h.Write([]byte(sign))
	return hex.EncodeToString(h.Sum(nil))
}
