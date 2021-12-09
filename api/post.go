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

var postUser model.Post

func Post(c *gin.Context) {
	iUsername, _ := c.Get("username")
	postUser.Username = iUsername.(string)
	postUser.Txt = c.PostForm("userPost")

	err := service.AddPost(postUser.Username, postUser.Txt)
	if err != nil {
		fmt.Println("insert post failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithUsernameAndDate(c, postUser.Username, "留言成功！", time.Now())
}

func GetOnesPost(c *gin.Context) {
	postUser.Username = c.PostForm("wantGetPostUsername")
	postUser.Txt = c.PostForm("postTxt")

	PostID, err := service.SelectByPostID(postUser.Username, postUser.Txt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时没有留言")
			return
		}
		fmt.Println("get ones post failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	err, posts, comments := service.GetPost(PostID)
	if err != nil {
		fmt.Println("get post failed at slice,err:", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range posts {
		tool.RespPostAndComment(c, posts[i].Username, posts[i].Txt, comments[i].Username, comments[i].Txt)
	}
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

	err = service.ChangePost(newPost, PostID)
	if err != nil {
		fmt.Println("changePost failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithUsernameAndDate(c, postUser.Username, "更改留言成功", time.Now())
}

func getAllPost(c *gin.Context) {
	err, userPosts, Time := service.GetAllPost()
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "竟然没留言")
			return
		}
		fmt.Println("getAllPost failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range userPosts {
		tool.RespSuccessfulWithUsernameAndDate(c, userPosts[i].Username, userPosts[i].Txt, Time[i])
	}
}
