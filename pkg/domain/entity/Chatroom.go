package entity

type ChatRoom struct {
	RoomId   string
	RoomName string
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

type ChatRoomRepo interface {
	CreateChatRoom(chatroom *ChatRoom) (err error)
	GetChatRoom(id string) (*ChatRoom, error)
	GetAllChatRooms() ([]ChatRoom, error)
}
