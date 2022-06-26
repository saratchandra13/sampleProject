package httpserver

import (
	"github.com/ShareChat/service-template/pkg/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddReviewReq struct {
	BeerId        string `json:"beerId"`
	ReviewComment string `json:"reviewComment"`
	Rating        int8   `json:"rating"`
}

type AddReviewRes struct {
	ReviewId string `json:"reviewId"`
}


func addReview(c *gin.Context) {
	userId := c.GetHeader("x-user-id")

	var req AddReviewReq
	err := c.Bind(&req)
	if err != nil {
		appInteractor.logger.Error("invalid request payload", err, &req)
		c.JSON(http.StatusBadRequest, "Failed")
		return
	}
	
	var ri = services.ReviewInfo{
		BeerId:        req.BeerId,
		ReviewComment: req.ReviewComment,
		Rating:        req.Rating,
	}
	reviewId, err := appInteractor.appLogic.ReviewBeer(userId, &ri)
	if err != nil {
		appInteractor.logger.Error("failed to add review", err, &req)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, &AddReviewRes{ReviewId:reviewId})
}
