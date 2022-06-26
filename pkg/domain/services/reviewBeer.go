package services

import (
	"fmt"
	"github.com/ShareChat/service-template/pkg/domain/entity"
	"github.com/ShareChat/service-template/pkg/domain/entity/enterr"
	"github.com/pkg/errors"
)

const (
	errUserNotVerified = "user not allowed to add review"
	errFailedToAddReview = "failed to add review"
	fallbackReviewId = "1234"
)

type ReviewInfo struct {
	BeerId        string `json:"beerId"`
	ReviewComment string `json:"reviewComment"`
	Rating        int8   `json:"rating"`
}

func (al *appLogic) ReviewBeer(userId string, ri *ReviewInfo) (string, error) {
	user, err := al.userRepo.GetUser(userId)
	if err != nil {
		return "", err
	}
	if !user.Verified {
		return "", errors.New(errUserNotVerified)
	}

	reviewMeta := entity.NewReview()
	reviewMeta.BeerId = ri.BeerId
	reviewMeta.Meta = ri.ReviewComment
	reviewMeta.Rating = ri.Rating
	reviewMeta.UserId = user.Id

	reviewId, err := al.reviewRepo.AddReview(reviewMeta)
	if err != nil {
		switch err := errors.Cause(err).(type) {
		case *enterr.CustomError:
			fmt.Println(err.Code, err.Msg, err.Err)
			if err.Code == enterr.AddReviewFailed {
				// do something
				// error is supposed to be handled once. Since we have handled it at this point we have to figure out
				// a logic of what to do next. We can log the original error here like this `logger.error(err.Msg, err.Err, nil)`
				// then continue with the logic after logging
				al.logger.Info(err.Msg, err.Err, nil)
			}
			return fallbackReviewId, nil
		default:
			return "", errors.New(errFailedToAddReview)
		}
	}

	return reviewId, nil
}
