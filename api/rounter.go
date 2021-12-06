package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()
	/*
		userGroup := engine.Group("/user")

		PostGroup := engine.Group("/post")

		comment   := engine.Group("/comment")
	*/
	engine.Run()
}
