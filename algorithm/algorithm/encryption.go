package algorithm

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"
)

// 加密方式
func Encryption(password string) string {
	// 随机字符
	words := "qwersd,yuio.tpa;vbnm6789QW~1234'=HJKL[zxc5`!@ERTYU]?fghFG&*-ZNM#ASD%^$jklIOP_XCVB+"
	rand.Seed(time.Now().UnixNano())
	tempStr := ""
	// 加盐
	for i := range password {
		tempStr += string(words[int(password[i])%len(words)]) + string(password[i])
	}

	passwd := make([]byte, 0)
	// sha512 加密
	code := sha512.Sum512([]byte(tempStr))
	passwd = append(passwd, code[:]...)

	waitTo := hex.EncodeToString(passwd)
	tempStr = ""
	// 加盐
	for i := range waitTo {
		tempStr += string(words[int(waitTo[i])%len(words)]) + string(waitTo[i])
	}
	// 再加密
	code = sha512.Sum512([]byte(tempStr))
	passwd = []byte{}
	passwd = append(passwd, code[:]...)

	return hex.EncodeToString(passwd)
}
