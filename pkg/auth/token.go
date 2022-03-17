package auth

import (
	"crypto/md5"
	"fmt"
)

func GenerateToken(timeStamp string) string {
	preHashTok := "R300_VozniRed_2015" + timeStamp
	postHashTok := fmt.Sprintf("%x", md5.Sum([]byte(preHashTok)))
	return postHashTok
}
