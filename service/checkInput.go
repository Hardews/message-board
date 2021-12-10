package service

import (
	"message-board/model"
	"strings"
)

func CheckPasswordLength(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func CheckUsernameLength(username string) bool {
	if len(username) < 6 {
		return false
	}
	return true
}

func CheckTxtLength(TXT string) bool {
	if len(TXT) > 20 {
		return false
	}
	return true
}

func CheckSensitiveWords(txt string) bool {
	var SensitiveWords = make([]string, 0)
	SensitiveWords = append(SensitiveWords, "fuck")
	SensitiveWords = append(SensitiveWords, "傻逼")

	for i, _ := range SensitiveWords {
		flag := strings.Contains(txt, SensitiveWords[i])
		if !flag {
			return false
		}
	}
	return true
}

func CheckInfoBySensitiveWord(userInfo model.UserInfo) bool {
	// 判断是否含有敏感词汇
	flag := CheckSensitiveWords(userInfo.Name)
	if !flag {
		return false
	}
	flag = CheckSensitiveWords(userInfo.Professional)
	if !flag {
		return false
	}
	flag = CheckSensitiveWords(userInfo.Specialty)
	if !flag {
		return false
	}
	flag = CheckSensitiveWords(userInfo.Professional)
	if !flag {
		return false
	}
	return true
}

func CheckInfoLength(userInfo model.UserInfo) bool {
	flag := CheckTxtLength(userInfo.Name)
	if !flag {
		return false
	}
	flag = CheckTxtLength(userInfo.Professional)
	if !flag {
		return false
	}
	flag = CheckTxtLength(userInfo.Specialty)
	if !flag {
		return false
	}
	flag = CheckTxtLength(userInfo.Professional)
	if !flag {
		return false
	}
	return true
}
