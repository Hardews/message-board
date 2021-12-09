package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", Register)
	engine.POST("/login", Login)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)

		userGroup.POST("/password", changePassword)
		userGroup.POST("/writeInfo", writeInfo)     //填写个人信息
		userGroup.GET("/info", getInfo)             //获取个人信息
		userGroup.POST("/info", changeInfo)         //更改个人信息
		userGroup.POST("/likeComment", LikeComment) //给某个人的评论点赞
		userGroup.POST("/likePost", LikePost)
	}

	PostGroup := engine.Group("/post")
	{
		PostGroup.Use(auth)

		PostGroup.POST("/", Post)         //发留言
		PostGroup.GET("/", GetOnesPost)   //获取某个人的留言及其下属评论
		PostGroup.DELETE("/", DeletePost) //删除留言
		PostGroup.GET("/all", getAllPost)
		PostGroup.POST("change", changePost)
	}

	comment := engine.Group("/comment")
	{
		comment.Use(auth)

		comment.POST("/", addComment)      //发表评论
		comment.DELETE("/", deleteComment) //删除评论
		comment.POST("/update", changeComment)
	}

	engine.Run(":8090")
}
