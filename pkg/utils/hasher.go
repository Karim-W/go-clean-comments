package utils

import (
	"crypto"
	"encoding/hex"
)

func CreateMd5Hash(str string) string {
	hasher := crypto.MD5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
