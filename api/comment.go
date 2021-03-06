package api

import (
	"fmt"
	"time"

	"message-board/model"
	"message-board/service"
	"message-board/tool"

	"database/sql"
	"github.com/gin-gonic/gin"
)

var commentUser model.Comment

func addComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	commentUser.Username = iUsername.(string)

	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")
	commentUser.Txt = c.PostForm("comment")

	flag := service.CheckTxtLength(commentUser.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "评论过长")
		return
	}
	flag = service.CheckSensitiveWords(commentUser.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "评论含有敏感词")
		return
	}

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

	err = service.DeleteComment(commentUser)
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

	flag := service.CheckTxtLength(newComment)
	if !flag {
		tool.RespErrorWithDate(c, "评论过长")
		return
	}
	flag = service.CheckSensitiveWords(newComment)
	if !flag {
		tool.RespErrorWithDate(c, "评论含有敏感词")
		return
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

	err = service.ChangeComment(newComment, commentUser)
	if err != nil {
		fmt.Println("changeComment failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithUsernameAndDate(c, commentUser.Username, "修改评论成功", time.Now(), commentUser.LikeNum)
}
