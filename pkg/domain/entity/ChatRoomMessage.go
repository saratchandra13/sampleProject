package entity

import "time"

type Message struct {
	SenderId        string
	ChatroomId      string
	Message         string
	TimeStamp       time.Time
	MessageMetaData interface{}
}

type ChatRoomMessageRepo interface {
	SendMessageToChatRoom(userID string, chatRoomId string, Message string) error
	RetrieveAllMessagesFromChatRoom(chatRoomId string) ([]Message, error)
}
