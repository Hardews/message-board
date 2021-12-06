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

func RespSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"info": "成功！",
	})
}

func RespSuccessfulWithDate(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"date": date,
	})
}
