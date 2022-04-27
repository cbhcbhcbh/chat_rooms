package v1

import (
	"chat_rooms/internal/service"
	"chat_rooms/pkg/common/request"
	"chat_rooms/pkg/common/response"
	"chat_rooms/pkg/global/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(c *gin.Context) {
	log.Logger.Info(c.Query("key"))
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if err != nil {
		log.Logger.Error("bindQueryError", log.Any("bindQueryError", err))
	}
	log.Logger.Info("messageRequest params: ", log.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}
