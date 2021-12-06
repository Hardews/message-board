package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
	"message-board/tool"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	res,err :=service.CheckPassword(username,password)
	if err != nil{
		switch err {
		case sql.ErrNoRows: c.JSON(http.StatusOK,gin.H{
			"info":"无此账号",
		})
		default:fmt.Println(err)
		        tool.RespInternetError(c)
		}
	}
	if res{
		tool.RespSuccessful(c)
	}else {
		c.JSON(http.StatusOK,gin.H{
			"info":"密码错误！",
		})
	}


}

func Register(c *gin.Context) {

}

func replacePassword(c *gin.Context) {

}

func getInfo(c *gin.Context) {

}

func replaceInfo(c *gin.Context) {

}
