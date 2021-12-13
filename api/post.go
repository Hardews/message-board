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
var err error

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

	tool.RespSuccessfulWithDate(c, "发布成功")
}

func DeletePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	postUser.Username = iUsername.(string)
	postWantDelete := c.PostForm("post")
	postUser, err = service.SelectAllByPostID(postUser.Username, postWantDelete)
	if err != nil {
		tool.RespErrorWithDate(c, "未查询到该留言")
		return
	}

	err = service.DeletePost(postUser.PostID, postWantDelete)
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

	postUser, err = service.SelectAllByPostID(postUser.Username, postUser.Txt)
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

	err = service.ChangePost(newPost, postUser.PostID)
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

	postUser, err = service.SelectAllByPostID(postUser.Username, postUser.Txt)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时没有这个评论区")
			return
		}
		fmt.Println("get ones post failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	//获取下属评论
	err, comments := service.GetCommentsSection(postUser.PostID)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("get comments section failed,err :", err)
		return
	}

	//输出第一层，留言
	root := &model.TreeNode{
		ID:  postUser.PostID,
		Txt: postUser.Txt,
	}
	tool.RespSuccessfulWithInfo(c, root.Txt, postUser.Username, postUser.PostTime, 1, postUser.LikeNum)

	//遍历，输出评论
	for i, _ := range comments {
		branch := &model.TreeNode{
			ID:      comments[i].CommentId,
			Name:    comments[i].Username,
			Txt:     comments[i].Txt,
			Time:    comments[i].Time,
			LikeNum: comments[i].LikeNum,
		}
		root.Branch = branch
		tool.RespSuccessfulWithInfo(c, root.Branch.Txt, root.Branch.Name, root.Branch.Time, i+2, root.Branch.LikeNum)
	}
}

func GetAllPosts(c *gin.Context) {
	err, userPost := service.GetAllPosts()
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "暂时没有留言")
			return
		}
		fmt.Println("get all posts failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	for i, _ := range userPost {
		tool.RespSuccessfulWithInfo(c, userPost[i].Txt, userPost[i].Username, userPost[i].PostTime, i, userPost[i].LikeNum)
	}
}
