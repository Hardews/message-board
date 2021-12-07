package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/tool"
)

func postComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	postUsername := c.PostForm("postUsername")
	post := c.PostForm("post")
	_, err := dao.SelectPost(postUsername, post)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select userPost failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	comment := c.PostForm("comment")

	err = dao.PostComment(username, post, comment)
	if err != nil {
		fmt.Println("insert comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}
