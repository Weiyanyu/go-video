package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//GenHashWithMD5 generate hash by param with MD5 encode
func GenHashWithMD5(str string) string {
	md5Bytes := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Bytes[:])
}
