package entity

import "encoding/json"

const (
	MaleTypeGender   = "male"
	FemaleTypeGender = "female"
)

type User struct {
	userId   string
	userName string
	Gender   string
}

func NewUser() *User {
	return &User{}
}

func (u *User) PrettyPrint() []byte {
	marshalVal, _ := json.MarshalIndent(u, "", " ")
	return marshalVal
}

type UserRepo interface {
	GetUser(userID string) (*User, error)
	AddUser(userId string) (*User, error)
}
