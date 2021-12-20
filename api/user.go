package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	res, err := service.CheckPassword(username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此账号")
			return
		}
		fmt.Println(err)
		tool.RespInternetError(c)
		return
	}
	if res {
		c.SetCookie("user_login", username, 600, "/", "", false, true)
		tool.RespSuccessful(c)
	} else {
		tool.RespErrorWithDate(c, "密码错误")
		return
	}

}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := service.CheckSensitiveWords(username)
	if !flag {
		tool.RespErrorWithDate(c, "用户名含敏感词")
	}
	flag = service.CheckUsernameLength(username)
	if !flag {
		tool.RespErrorWithDate(c, "用户名长度不符合要求")
	}

	err, res := service.CheckUsername(username)
	if res == false && err == nil {
		tool.RespErrorWithDate(c, "用户名已存在")
		return
	} else if err != nil && res == false {
		fmt.Println("CheckUsername failed , err : ", err)
		tool.RespInternetError(c)
		return
	}

	flag = service.CheckPasswordLength(password)
	if !flag {
		tool.RespErrorWithDate(c, "密码长度不符合要求")
		return
	}

	err, password = service.Encryption(password)
	if err != nil {
		fmt.Println("register failed , err :", err)
		tool.RespInternetError(c)
		return
	}
	err = service.WriteIn(username, password)
	if err != nil {
		fmt.Println("register failed , err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func changePassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")

	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	res, err := service.CheckPassword(username, oldPassword)
	if err != nil {
		fmt.Println(err)
		tool.RespInternetError(c)
		return
	}
	if !res {
		tool.RespErrorWithDate(c, "旧密码错误")
		return
	}

	flag := service.CheckPasswordLength(newPassword)
	if !flag {
		tool.RespErrorWithDate(c, "密码长度不符合要求")
		return
	}

	err, newPassword = service.Encryption(newPassword)
	if err != nil {
		fmt.Println("change password failed, err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password failed, err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessful(c)
}

func writeInfo(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	var userInfo model.UserInfo
	userInfo.Name = c.PostForm("Name")
	userInfo.Professional = c.PostForm("Professional")
	userInfo.Specialty = c.PostForm("Specialty")
	userInfo.School = c.PostForm("University")

	flag := service.CheckInfoLength(userInfo)
	if !flag {
		tool.RespErrorWithDate(c, "提交了过长信息")
		return
	}
	flag = service.CheckInfoBySensitiveWord(userInfo)
	if !flag {
		tool.RespErrorWithDate(c, "提交信息含敏感词汇")
		return
	}

	err := service.WriteInfo(userInfo, username)
	if err != nil {
		fmt.Println("write info failed , err : ", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessful(c)
}

func getInfo(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	userInfo, err := service.GetInfo(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "还没填写个人信息")
			return
		}
		fmt.Println("get info failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespErrorWithDate(c, userInfo)
}

func changeInfo(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	var newUserInfo model.UserInfo

	newUserInfo.Name = c.PostForm("newName")
	newUserInfo.Professional = c.PostForm("newProfessional")
	newUserInfo.Specialty = c.PostForm("newSpecialty")
	newUserInfo.School = c.PostForm("newUniversity")
	err, newUserInfo = service.CheckInputInfo(username, newUserInfo)
	if err != nil {
		fmt.Println("check input info failed,err", err)
		tool.RespInternetError(c)
		return
	}

	flag := service.CheckInfoLength(newUserInfo)
	if !flag {
		tool.RespErrorWithDate(c, "提交了过长信息")
		return
	}
	flag = service.CheckInfoBySensitiveWord(newUserInfo)
	if !flag {
		tool.RespErrorWithDate(c, "提交信息含有敏感词")
		return
	}

	flag, err = service.ChangeInfo(newUserInfo)
	if err != nil {
		fmt.Println("change userinfo failed , err :", err)
		tool.RespInternetError(c)
		return
	}

	if !flag {
		tool.RespErrorWithDate(c, "修改失败，查询不到该用户信息")
		return
	}
	tool.RespSuccessful(c)
}
