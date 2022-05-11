package router

import (
	v1 "chat_rooms/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	socket := RunSocket

	group := server.Group("")
	{
		group.GET("/user", v1.GetUserList)
		group.GET("/user/:uuid", v1.GetUserDetails)
		group.GET("/user/name", v1.GetUserOrGroupByName)
		group.POST("/user/register", v1.Register)
		group.POST("/user/login", v1.Login)
		group.PUT("/user", v1.ModifyUserInfo)

		group.POST("/friend", v1.AddFriend)

		group.GET("/message", v1.GetMessage)

		group.GET("/file/:fileName", v1.GetFile)
		group.POST("/file", v1.SaveFile)

		group.GET("/group/:uuid", v1.GetGroup)
		group.POST("/group/:uuid", v1.SaveGroup)
		group.POST("/group/join/:userUuid/:groupUuid", v1.JoinGroup)
		group.GET("/group/user/:uuid", v1.GetGroupUsers)

		group.GET("/socket.io", socket)
	}
	return server
}

func Recovery(context *gin.Context) {

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin !=
	}
}
