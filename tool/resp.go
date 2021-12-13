package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithDate(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"info": date,
	})
}

func RespInternetError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespSuccessfulWithDate(c *gin.Context, Date interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"info": Date,
	})
}

func RespSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "成功！",
	})
}

func RespSuccessfulWithUsernameAndDate(c *gin.Context, username string, date, time interface{}, likeNum int) {
	c.JSON(http.StatusOK, gin.H{
		"time":     time,
		"username": username,
		"post":     date,
		"likeNum":  likeNum,
	})
}

func RespSuccessfulWithInfo(c *gin.Context, txt, username, time string, branch, likeNum int) {
	c.JSON(http.StatusOK, gin.H{
		"楼层":  branch,
		"用户名": username,
		"内容":  txt,
		"时间":  time,
		"点赞数": likeNum,
	})
}
