package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type listReviewReq struct {
	BeerId string `json:"beerId"`
}

type reviewMeta struct {
	ReviewId     string `json:"reviewId"`
	BeerId string `json:"beerId"`
	UserId string `json:"userId"`
	Meta   string `json:"meta"`
	Rating int8 `json:"rating"`
}

type listReviewRes struct {
	ReviewList []*reviewMeta `json:"reviewList"`
}

func listReview(c *gin.Context) {
	var req listReviewReq
	err := c.Bind(&req)
	if err != nil {
		appInteractor.logger.Error("invalid request payload", err, &req)
		c.JSON(http.StatusBadRequest, "Failed")
		return
	}
	reviewList, err := appInteractor.appLogic.ListReview(req.BeerId)
	if err != nil {
		appInteractor.logger.Error("failed to list review", err, &req)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	var res = listReviewRes{ReviewList: []*reviewMeta{}}
	for _, r := range reviewList {
		rm := reviewMeta{
			ReviewId:     r.Id,
			BeerId: r.BeerId,
			UserId: r.UserId,
			Meta:   r.Meta,
			Rating: r.Rating,
		}
		res.ReviewList = append(res.ReviewList, &rm)
	}
	c.JSON(http.StatusOK, res)
}
