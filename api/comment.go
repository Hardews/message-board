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

func deleteComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	commentWantDelete := c.PostForm("comment")
	err := dao.DeleteComment(username, commentWantDelete)
	if err != nil {
		fmt.Println("delete comment failed , err :", err)
		return
	}
	tool.RespSuccessful(c)
}

func updateComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	oldComment := c.PostForm("oldComment")
	newComment := c.PostForm("newComment")

	err, flag := dao.SelectComment(username, oldComment)
	if err != nil {
		fmt.Println("select comment failed , err :", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "无此评论")
		return
	}

	err = dao.ChangeComment(username, newComment, oldComment)
	if err != nil {
		fmt.Println("changeComment failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithUsernameAndDate(c, username, "更改评论成功")
}
