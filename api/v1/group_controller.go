package v1

import (
	"chat_rooms/internal/service"
	"chat_rooms/pkg/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	groups, err := service.GroupService.GetGroups(uuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(groups))
}

func SaveGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	groups, err := service.GroupService.GetGroups(uuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(groups))
}

func JoinGroup(c *gin.Context) {
	groupUuid := c.Param("groupUuid")
	userUuid := c.Param("userUuid")
	err := service.GroupService.JoinGroup(groupUuid, userUuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

func GetGroupUsers(c *gin.Context) {
	groupUuid := c.Param("uuid")
	users := service.GroupService.GetUserIdByGroupUuid(groupUuid)
	c.JSON(http.StatusOK, response.SuccessMsg(users))
}
