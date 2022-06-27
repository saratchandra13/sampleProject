package memory

import (
	_ "context"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"time"
)

type InMemory struct {
	ChatRoomStore []entity.ChatRoom
	ChatRoomUsers []entity.ChatRoomUsers
	Users         []entity.User
	Messages      []entity.Message
}

func NewMemoryStore() *InMemory {
	return &InMemory{ChatRoomStore: make([]entity.ChatRoom, 0), ChatRoomUsers: make([]entity.ChatRoomUsers, 0),
		Users: make([]entity.User, 0), Messages: make([]entity.Message, 0)}
}

func (Mem *InMemory) CreateChatRoom(chatroom *entity.ChatRoom) (err error) {
	if chatroom != nil {
		Mem.ChatRoomStore = append(Mem.ChatRoomStore, *chatroom)
	}
	return nil
}

func (Mem *InMemory) GetChatRoom(id string) (*entity.ChatRoom, error) {
	if len(Mem.ChatRoomStore) != 0 {
		for _, room := range Mem.ChatRoomStore {
			if room.RoomId == id {
				return &room, nil
			}
		}
	}
	return nil, nil
}

func (Mem *InMemory) GetAllChatRooms() ([]entity.ChatRoom, error) {
	if len(Mem.ChatRoomStore) != 0 {
		return Mem.ChatRoomStore, nil
	}
	return nil, nil
}

func (Mem *InMemory) AddUserToChatroom(userId string, chatroomId string) (err error) {
	chatRoomUser := entity.ChatRoomUsers{RoomId: chatroomId, UserId: userId}
	Mem.ChatRoomUsers = append(Mem.ChatRoomUsers, chatRoomUser)
	return nil
}

func (Mem *InMemory) IsUserPartOfChatRoom(userId string, chatroomId string) (bool, error) {
	if len(Mem.ChatRoomUsers) > 0 {
		for _, user := range Mem.ChatRoomUsers {
			if user.UserId == userId && user.RoomId == chatroomId {
				return true, nil
			}
		}
	}
	return false, nil
}

func (Mem *InMemory) SendMessageToChatRoom(userID string, chatRoomId string, Message string) error {
	if Mem.Messages != nil {
		Message := entity.Message{
			Message:         Message,
			TimeStamp:       time.Now(),
			MessageMetaData: nil,
			SenderId:        userID,
			ChatroomId:      chatRoomId,
		}

		Mem.Messages = append(Mem.Messages, Message)
	}
	return nil
}

func (Mem *InMemory) RetrieveAllMessagesFromChatRoom(chatRoomId string) (response []entity.Message, err error) {
	response = make([]entity.Message, 0)
	if Mem.Messages != nil {
		for _, message := range Mem.Messages {
			if message.ChatroomId == chatRoomId {
				response = append(response, message)
			}
		}
	}
	return response, nil
}
