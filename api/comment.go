package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/model"
	"message-board/tool"
	"time"
)

var users model.Comment

func addComment(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	users.Username = iUsername.(string)

	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")

	users.PostID, err = dao.SelectByPostId(postName, postTxt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select postID failed , err :", err)
		tool.RespInternetError(c)
	}

	users.Txt = c.PostForm("comment")

	err = dao.PostComment(users.Username, users.Txt, users.PostID)
	if err != nil {
		fmt.Println("insert comment failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func deleteComment(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	users.Username = iUsername.(string)

	commentWantDelete := c.PostForm("comment")
	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")

	users.PostID, err = dao.SelectByPostId(postName, postTxt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select postID failed , err :", err)
		tool.RespInternetError(c)
	}

	err, users.CommentId = dao.SelectByCommentId(user.Username, commentWantDelete, user.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "没有这个评论")
			return
		}
		fmt.Println("select commentId failed, err:", err)
		tool.RespInternetError(c)
		return
	}
	err = dao.DeleteComment(users.CommentId)
	if err != nil {
		fmt.Println("select commentId failed, err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func changeComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	users.Username = iUsername.(string)

	oldComment := c.PostForm("oldComment")
	newComment := c.PostForm("newComment")

	err, flag := dao.SelectComment(users.Username, oldComment)
	if err != nil {
		fmt.Println("select comment failed , err :", err)
		tool.RespInternetError(c)
		return
	}
	if !flag {
		tool.RespErrorWithDate(c, "无此评论")
		return
	}

	err = dao.ChangeComment(oldComment, newComment)
	if err != nil {
		fmt.Println("changeComment failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithUsernameAndDate(c, users.Username, "修改评论成功", time.Now())
}
