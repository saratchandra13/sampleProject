package services

import (
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"github.com/saratchandra13/sampleProject/pkg/infrastructure/memory"
)

type AppInterface interface {
	AddUserToChatroom(users entity.ChatRoomUsers) error
}

type appLogic struct {
	memoryStore memory.InMemory
}

func NewAppLogic(memStore *memory.InMemory) AppInterface {
	return &appLogic{
		memoryStore: *memStore,
	}
}
