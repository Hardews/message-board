package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", Register)
		userGroup.POST("/login", Login)
		userGroup.POST("/password", auth, changePassword)

		userGroup.GET("/info", getInfo)      //获取个人信息
		userGroup.POST("/info", replaceInfo) //更改个人信息
	}

	PostGroup := engine.Group("/post")
	{
		PostGroup.POST("/", auth, Post)         //发留言
		PostGroup.GET("/", auth, GetPost)       //获取自己的留言
		PostGroup.DELETE("/", auth, DeletePost) //删除留言
		PostGroup.GET("/all", getAllPost)
		PostGroup.POST("change", auth, changePost)
	}

	comment := engine.Group("/comment")
	{
		comment.POST("/")   //发表评论
		comment.GET("/:id") //获取评论
		comment.DELETE("/") //删除评论
	}

	engine.Run(":8090")
}
