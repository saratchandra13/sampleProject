package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VegetableInfo struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Seller       string  `json:"Seller"`
	Price        float64 `json:"price"`
	CalorieCount float64 `json:"CalorieCount"`
}

type listVegetableRes struct {
	VegetableList []*VegetableInfo `json:"VegetableList"`
}

// handler for list vegetable functionality
func listVegetable(c *gin.Context) {
	VegetableList, err := appInteractor.appLogic.ListVegetable()
	if err != nil {
		appInteractor.logger.Error("failed to list vegetable", err, nil)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	// making response in the desired format
	var res = listVegetableRes{VegetableList: []*VegetableInfo{}}
	for _, vegetable := range VegetableList {
		bi := &VegetableInfo{
			Id:           vegetable.Id,
			Name:         vegetable.Name,
			Seller:       vegetable.Seller,
			Price:        vegetable.Price,
			CalorieCount: vegetable.CalorieCount,
		}
		res.VegetableList = append(res.VegetableList, bi)
	}
	c.JSON(http.StatusOK, res)
}
