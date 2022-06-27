package rest

import (
	"github.com/saratchandra13/sampleProject/config"
)

const (
	errFailedToParseBody = "failed to parse response body"
)

type UserSvc struct {
	config *config.Store
}

func NewUserSvc(config *config.Store) *UserSvc {
	return &UserSvc{config: config}
}

type userSvcRes struct {
	Name     string `json:"name"`
	UserId   string `json:"userId"`
	Handle   string `json:"handle"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	Verified int8   `json:"verified"`
}
