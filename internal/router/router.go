package router

import (
	v1 "chat_rooms/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	server := gin.Default()
	group := server.Group("")
	{
		group.GET("/group/:uuid", v1.GetGroup)
	}
	return server
}
