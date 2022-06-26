package services

import (
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"github.com/saratchandra13/sampleProject/third_party/platlogger"
)

type AppInterface interface {
	ListVegetable() (VegetableList, error)
	AddVegetable(*VegetableInfo) (string, error)
}

type appLogic struct {
	VegetableRepo entity.VegetableRepo
	userRepo      entity.UserRepo
	logger        *platlogger.Client
}

func NewAppLogic(vegetable entity.VegetableRepo, user entity.UserRepo, logger *platlogger.Client) AppInterface {
	return &appLogic{
		VegetableRepo: vegetable,
		userRepo:      user,
		logger:        logger,
	}
}
