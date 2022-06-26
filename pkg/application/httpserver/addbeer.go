package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/saratchandra13/sampleProject/pkg/domain/services"
	"net/http"
)

type AddBeerReq struct {
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}

type AddBeerRes struct {
	BeerId string `json:"beerId"`
}

func addBeer(c *gin.Context) {
	var req AddBeerReq
	err := c.Bind(&req)
	if err != nil {
		appInteractor.logger.Error("invalid request payload", err, &req)
		c.JSON(http.StatusBadRequest, "Failed")
		return
	}

	bi := services.BeerInfo{
		Name:  req.Name,
		Manuf: req.Manufacturer,
		Price: req.Price,
	}
	beerId, err := appInteractor.appLogic.AddBeer(&bi)
	if err != nil {
		appInteractor.logger.Error("failed to create beer entry", err, &req)
		c.JSON(http.StatusInternalServerError, "Failed")
		return
	}
	c.JSON(http.StatusOK, &AddBeerRes{BeerId: beerId})
}
