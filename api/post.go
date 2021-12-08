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

var user model.Post

func Post(c *gin.Context) {
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	user.Txt = c.PostForm("userPost")
	err := dao.Post(user.Username, user.Txt)
	if err != nil {
		fmt.Println("post insert failed, err : ", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithUsernameAndDate(c, user.Username, "留言成功！", time.Now())
}

func GetOnesPost(c *gin.Context) {
	user.Username = c.PostForm("wantGetPostUsername")
	user.Txt = c.PostForm("postTxt")
	PostID, err := dao.SelectByPostId(user.Username, user.Txt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "此账号无留言")
			return
		}
		fmt.Println("get PostID failed, err :", err)
		tool.RespInternetError(c)
		return
	}

	err, posts, comments := dao.SelectPostAndCommentByPostID(PostID)
	if err != nil {
		fmt.Println("get PostID failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range posts {
		tool.RespPostAndComment(c, posts[i].Username, posts[i].Txt, comments[i].Username, comments[i].Txt)
	}
}

func DeletePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	postWantDelete := c.PostForm("post")
	_, err := dao.SelectPost(user.Username, postWantDelete)
	if err != nil {
		tool.RespErrorWithDate(c, "删除失败，未查询到该留言")
		return
	}
	err = dao.DeletePost(user.Username, postWantDelete)
	if err != nil {
		fmt.Println("delete post failed , err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessful(c)
}

func changePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	user.Username = iUsername.(string)
	err := dao.GetPost(user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "无留言")
			return
		}
		fmt.Println("select post failed , err :", err)
		tool.RespInternetError(c)
		return
	}

	newPost := c.PostForm("newPost")
	oldPost := c.PostForm("oldPost")
	err = dao.ChangePost(user.Username, newPost, oldPost)
	if err != nil {
		fmt.Println("changePost failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithUsernameAndDate(c, user.Username, "更改留言成功", time.Now())
}

func getAllPost(c *gin.Context) {
	err, username, txt, Time := dao.GetAllPost()
	if err != nil {
		fmt.Println("getAllPost failed, err :", err)
		tool.RespInternetError(c)
		return
	}
	for i, _ := range username {
		tool.RespSuccessfulWithUsernameAndDate(c, username[i], txt[i], Time[i])
	}
}
