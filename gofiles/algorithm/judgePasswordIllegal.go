package algorithm

import (
	"errors"
	"strconv"
)

// 判断密码合法性（复杂性）
func JudgePasswordIllegal(password string) error {
	// 合法长度限制
	passedLength := 12
	// 密码长度
	length := len(password)
	// 字母数量
	alphaNum := 0
	digitNum := 0
	// 字母及特殊字符表
	alpha := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM-_+=.?!@#$%^&*"
	// 数字表
	digit := "0123456789"
	// 字母及特殊字符记录
	recAlpha := map[rune]bool{}
	// 数字字符记录
	recDigit := map[rune]bool{}

	// 记录
	for _, data := range alpha {
		recAlpha[data] = true
	}
	for _, data := range digit {
		recDigit[data] = true
	}

	// 判断字母及特殊字符数量、数字数量
	for _, data := range password {
		if _, ok := recAlpha[data]; ok {
			alphaNum++
		} else if _, ok := recDigit[data]; ok {
			digitNum++
		}
	}

	// 字母及特殊字符数量至少为总长度一半
	if digitNum+alphaNum != length {
		return errors.New("密码格式错误，只能存在数字字符、字母或特殊字符！")
	} else if length < passedLength {
		return errors.New("密码长度至少为 " + strconv.Itoa(passedLength) + " 位")
	} else if alphaNum < passedLength/3 {
		return errors.New("字母至少需要 " + strconv.Itoa(passedLength/3) + " 位")
	}

	// 返回
	return nil
}
