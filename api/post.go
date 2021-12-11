package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"net/http"
	"time"
)

var postUser model.Post

func Post(c *gin.Context) {
	iUsername, _ := c.Get("username")
	postUser.Username = iUsername.(string)
	postUser.Txt = c.PostForm("userPost")

	flag := service.CheckSensitiveWords(postUser.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "留言含有敏感词")
		return
	}
	flag = service.CheckTxtLength(postUser.Txt)
	if !flag {
		tool.RespErrorWithDate(c, "留言过长")
		return
	}

	err := service.AddPost(postUser.Username, postUser.Txt)
	if err != nil {
		fmt.Println("insert post failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	ID, _ := service.SelectByPostID(postUser.Username, postUser.Txt)
	err = service.CreateCommentsSection(ID, postUser)
	tool.RespSuccessfulWithDate(c, "创建留言区成功")
}

func DeletePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	postUser.Username = iUsername.(string)
	postWantDelete := c.PostForm("post")
	PostID, err := service.SelectByPostID(postUser.Username, postWantDelete)
	if err != nil {
		tool.RespErrorWithDate(c, "未查询到该留言")
		return
	}

	err = service.DeletePost(PostID, postWantDelete)
	if err != nil {
		fmt.Println("delete post failed , err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessful(c)
}

func changePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	postUser.Username = iUsername.(string)
	postUser.Txt = c.PostForm("oldPost")
	newPost := c.PostForm("newPost")

	PostID, err := service.SelectByPostID(postUser.Username, postUser.Txt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无留言")
			return
		}
		fmt.Println("select post failed , err :", err)
		tool.RespInternetError(c)
		return
	}

	flag := service.CheckTxtLength(newPost)
	if !flag {
		tool.RespErrorWithDate(c, "留言过长(大于20字)")
		return
	}
	flag = service.CheckSensitiveWords(newPost)
	if !flag {
		tool.RespErrorWithDate(c, "留言含有敏感词")
		return
	}

	err = service.ChangePost(newPost, PostID)
	if err != nil {
		fmt.Println("changePost failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithUsernameAndDate(c, postUser.Username, "更改留言成功", time.Now(), postUser.LikeNum)
}

func GetOnesPost(c *gin.Context) {
	postUser.Username = c.PostForm("wantGetPostUsername")
	postUser.Txt = c.PostForm("postTxt")

	PostID, err := service.SelectByPostID(postUser.Username, postUser.Txt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时没有这个评论区")
			return
		}
		fmt.Println("get ones post failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	err, commentsSection := service.GetPost(PostID)
	if err != nil {
		fmt.Println("get post failed at slice,err:", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range commentsSection {
		c.JSON(http.StatusOK, gin.H{
			"level":    i,
			"username": commentsSection[i].Username,
			"txt":      commentsSection[i].Txt,
			"likeNum":  commentsSection[i].LikeNum,
		})
	}
}
