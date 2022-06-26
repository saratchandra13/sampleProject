package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/saratchandra13/sampleProject/pkg/domain/services"
	"net/http"
)

type AddVegetableReq struct {
	Name   string  `json:"name"`
	Seller string  `json:"Seller"`
	Price  float64 `json:"price"`
}

type AddVegetableRes struct {
	VegetableId string `json:"VegetableId"`
}

func addVegetable(c *gin.Context) {
	var req AddVegetableReq
	err := c.Bind(&req)
	if err != nil {
		appInteractor.logger.Error("invalid request payload", err, &req)
		c.JSON(http.StatusBadRequest, "Failed")
		return
	}

	bi := services.VegetableInfo{
		Name:   req.Name,
		Seller: req.Seller,
		Price:  req.Price,
	}
	VegetableId, err := appInteractor.appLogic.AddVegetable(&bi)
	if err != nil {
		appInteractor.logger.Error("failed to create vegetable entry", err, &req)
		c.JSON(http.StatusInternalServerError, "Failed")
		return
	}
	c.JSON(http.StatusOK, &AddVegetableRes{VegetableId: VegetableId})
}
