package services

import (
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"github.com/saratchandra13/sampleProject/third_party/platlogger"
)

type AppInterface interface {
	ListBeer() (beerList, error)
	AddBeer(*BeerInfo) (string, error)
	ReviewBeer(string, *ReviewInfo) (string, error)
	ListReview(string) ([]*entity.Review, error)
}

type appLogic struct {
	beerRepo   entity.BeerRepo
	userRepo   entity.UserRepo
	reviewRepo entity.ReviewRepo
	logger     *platlogger.Client
}

func NewAppLogic(beer entity.BeerRepo, user entity.UserRepo, review entity.ReviewRepo, logger *platlogger.Client) AppInterface {
	return &appLogic{
		beerRepo:   beer,
		userRepo:   user,
		reviewRepo: review,
		logger:     logger,
	}
}
