package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"message-board/tool"
)

func LikeComment(c *gin.Context) {
	var err error
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	postName := c.PostForm("postName")
	postTxt := c.PostForm("postTxt")
	commentUser.Username = c.PostForm("commentName")
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

	err, userLikeNum := service.SelectCommentLikeNum(commentUser.CommentId)
	if err != nil {
		fmt.Println("select comment like num failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	err = service.LikeComment(userLikeNum, commentUser, username)
	if err != nil {
		fmt.Println("like failed , err :", err)
		tool.RespInternetError(c)
		return
	}

	tool.RespSuccessfulWithDate(c, "点赞成功")
}

func LikePost(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

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

	err, likeNum := service.SelectPostNum(PostID)
	if err != nil {
		fmt.Println("get post like num failed ,err", err)
		tool.RespInternetError(c)
		return
	}

	err = service.LikePost(likeNum, PostID, username)
	if err != nil {
		fmt.Println("insert post like failed ,err :", err)
		tool.RespInternetError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, "点赞成功")

}
