package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/tool"
)

func Post(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	userPost := c.PostForm("userPost")
	err := dao.Post(username, userPost)
	if err != nil {
		fmt.Println("post insert failed, err : ", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "留言成功！")
}

func GetPost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	err, user := dao.GetPost(username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "此账号无留言")
			return
		}
		fmt.Println("get output failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range user {
		tool.RespSuccessfulWithDate(c, user[i])
	}

}

func DeletePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	postWantDelete := c.PostForm("post")
	err := dao.DeletePost(username, postWantDelete)
	if err != nil {
		fmt.Println("delete post failed , err :", err)
		return
	}
	tool.RespSuccessful(c)
}

func getAllPost(c *gin.Context) {
	err, output := dao.GetAllPost()
	if err != nil {
		fmt.Println("getAllPost failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range output {
		tool.RespSuccessfulWithDate(c, output[i])
	}
}
