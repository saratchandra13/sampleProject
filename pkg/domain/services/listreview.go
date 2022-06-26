package services

import (
	"github.com/ShareChat/service-template/pkg/domain/entity"
)

func (al *appLogic) ListReview(beerId string) ([]*entity.Review, error) {
	resp, err := al.reviewRepo.ListReview(beerId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
