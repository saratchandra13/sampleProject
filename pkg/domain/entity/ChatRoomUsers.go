package entity

type ChatRoomUsers struct {
	UserId string
	RoomId string
}

func NewChatRoomUsers() *ChatRoomUsers {
	return &ChatRoomUsers{}
}

type ChatRoomUsersRepo interface {
	AddUserToChatroom(userId string, chatroomId string) (err error)
	IsUserPartOfChatRoom(userId string, chatroomId string) (bool, error)
}
