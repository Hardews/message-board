package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithNoRows(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "账号不存在",
	})
}

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
