package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ShareChat/service-template/config"
	"github.com/ShareChat/service-template/pkg/domain/entity"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	errFailedToGetUser = "failed to get user meta"
	errFailedToCreateRequest = "failed to create request"
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

func (us *UserSvc) GetUser(userId string) (*entity.User, error) {
	getUrl := fmt.Sprintf("%s/user-service/v1/users/%s", us.config.DataSources.UserSvc.HttpEndpoint.Url, userId)
	client := http.Client{
		Timeout: time.Duration(us.config.DataSources.UserSvc.HttpEndpoint.Timeout) * time.Second,
	}
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, errFailedToCreateRequest)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, errFailedToGetUser)
	}
	defer resp.Body.Close()
	byteData, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp:", string(byteData))
	if err != nil {
		return nil, errors.Wrap(err,errFailedToParseBody)
	}
	var respBody = userSvcRes{}

	if err := json.Unmarshal(byteData, &respBody); err != nil {
		return nil, errors.Wrap(err, errFailedToParseBody)
	}

	// creating user entity from the response of user-service

	// one point to note here is that instead of parsing the response to `entity.user` there is a conversion layer.
	// So response is parsed in the struct `userSvcRes` and then used to create `entity.user`. This decouples the entity
	// structure from the actual response received from the service layer and both these structs can be modified independently
	user := entity.NewUser()
	user.Name = respBody.Name
	user.Handle = respBody.Handle
	user.Id = respBody.UserId
	user.Verified = false
	if respBody.Verified == 1 {
		user.Verified = true
	}

	user.Gender = entity.MaleTypeGender
	if respBody.Gender == "F" {
		user.Gender = entity.FemaleTypeGender
	}
	user.Language = respBody.Language

	return user, nil
}
