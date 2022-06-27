package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"net/http"
)

type AddUserToChatRoom struct {
	UserId     string `json:"userId"`
	ChatroomId string `json:"chatroomId"`
}

func JoinChatRoom(c *gin.Context) {
	var req AddUserToChatRoom
	err := c.Bind(&req)
	if err != nil {
		appInteractor.logger.Error("invalid request payload", err, &req)
		c.JSON(http.StatusBadRequest, "Failed")
		return
	}

	request := entity.ChatRoomUsers{
		UserId: req.UserId,
		RoomId: req.ChatroomId,
	}

	err = appInteractor.appLogic.AddUserToChatroom(request)
	if err != nil {
		appInteractor.logger.Error("failed to add user to chatroom", err, &req)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, "ok")
}
