package entity

import "encoding/json"

const (
	MaleTypeGender   = "male"
	FemaleTypeGender = "female"
)

type User struct {
	Id       string
	Name     string
	Handle   string
	Gender   string
	Language string
	Verified bool
}

func NewUser() *User {
	return &User{}
}

func (u *User) PrettyPrint() []byte {
	marshalVal, _ := json.MarshalIndent(u, "", " ")
	return marshalVal
}

type UserRepo interface {
	GetUser(string) (*User, error)
}
