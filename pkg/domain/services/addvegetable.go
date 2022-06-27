package services

import (
	"errors"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
)

func (al *appLogic) AddUserToChatroom(users entity.ChatRoomUsers) error {
	if len(users.RoomId) == 0 || len(users.UserId) == 0 {
		return errors.New("not valid data")
	}

	if room, err := al.memoryStore.GetChatRoom(users.RoomId); err != nil {
		return err
	} else {
		if room == nil {
			return errors.New("no chatRoom Id found with this Id")
		}
	}

	return al.memoryStore.AddUserToChatroom(users.UserId, users.RoomId)
}
