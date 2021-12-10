package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"time"
)

var commentUser model.Comment

func addComment(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	commentUser.Username = iUsername.(string)

	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")
	commentUser.Txt = c.PostForm("comment")

	commentUser.PostID, err = service.SelectByPostID(postName, postTxt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select postID failed , err :", err)
		tool.RespInternetError(c)
	}

	err = service.AddComment(commentUser)
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
	commentUser.Username = iUsername.(string)

	commentUser.Txt = c.PostForm("commentWantDelete")
	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")

	commentUser.PostID, err = service.SelectByPostID(postName, postTxt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select postID failed , err :", err)
		tool.RespInternetError(c)
	}

	commentUser.CommentId, err = service.SelectCommentID(commentUser)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "没有这个评论")
			return
		}
		fmt.Println("select commentId failed, err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.DeleteComment(commentUser.CommentId, commentUser.PostID)
	if err != nil {
		fmt.Println("select commentId failed, err:", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func changeComment(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	commentUser.Username = iUsername.(string)

	commentUser.Txt = c.PostForm("oldComment")
	newComment := c.PostForm("newComment")

	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")

	commentUser.PostID, err = service.SelectByPostID(postName, postTxt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无此用户或留言")
			return
		}
		fmt.Println("select postID failed , err :", err)
		tool.RespInternetError(c)
	}

	commentUser.CommentId, err = service.SelectCommentID(commentUser)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "没有这个评论")
			return
		}
		fmt.Println("select commentId failed, err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.ChangeComment(newComment, commentUser.CommentId)
	if err != nil {
		fmt.Println("changeComment failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithUsernameAndDate(c, commentUser.Username, "修改评论成功", time.Now(), commentUser.LikeNum)
}
