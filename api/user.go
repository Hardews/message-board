package api

import (
	"database/sql"
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
	res,err :=service.CheckPassword(username,password)
	if err != nil{
		switch  {
		case res==false && err==sql.ErrNoRows:
			c.JSON(http.StatusOK,gin.H{
			"info":"无此账号",
		})
			return
		default:fmt.Println(err)
		        tool.RespInternetError(c)
		}
		return
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
	username := c.PostForm("username")
	password := c.PostForm("password")
	err,res := service.CheckUsername(username)
	switch  {
	case err == nil && res == false:
		c.JSON(http.StatusOK,gin.H{
		"info":"用户名已存在",
	})
		return
	case err !=nil && res == false:
		fmt.Println("CheckUsername failed , err : ",err)
		return
	}
	if len(password)<6 {
		c.JSON(http.StatusOK,gin.H{
			"info":"密码过短",
		})
		return
	}
	err = dao.WriteIn(username,password)
	if err!=nil{
		fmt.Println("insert failed, err :",err)
		return
	}
	tool.RespSuccessful(c)
}

func replacePassword(c *gin.Context) {

}

func getInfo(c *gin.Context) {

}

func replaceInfo(c *gin.Context) {

}
