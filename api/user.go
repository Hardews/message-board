package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
	"message-board/tool"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	res, err := service.CheckPassword(username, password)
	if err != nil {
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
	err, res := service.CheckUsername(username)
	if res == false && err == nil {
		tool.RespErrorWithDate(c, "用户名已存在")
		return
	} else if err != nil && res == false {
		fmt.Println("CheckUsername failed , err : ", err)
		tool.RespInternetError(c)
		return
	}

	if len(password) < 6 {
		tool.RespErrorWithDate(c, "密码过短")
		return
	}
	err = dao.WriteIn(username, password)
	if err != nil {
		fmt.Println("insert failed, err :", err)
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
		fmt.Println("checkPassword failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	if !res {
		tool.RespErrorWithDate(c, "旧密码错误")
		return
	}
	if len(newPassword) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"info": "密码过短",
		})
		return
	}
	err = dao.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("replace password failed, err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func getInfo(c *gin.Context) {

}

func replaceInfo(c *gin.Context) {

}
