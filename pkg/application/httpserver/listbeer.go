package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type beerInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Manufacturer   string `json:"manufacturer"`
	Price          float64 `json:"price"`
	AlcoholContent float64 `json:"alcoholContent"`
}

type listBeerRes struct {
	BeerList []*beerInfo `json:"beerList"`
}


// handler for list beer functionality
func listBeer(c *gin.Context) {
	beerList, err := appInteractor.appLogic.ListBeer()
	if err != nil {
		appInteractor.logger.Error("failed to list beer", err, nil)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	// making response in the desired format
	var res = listBeerRes{BeerList: []*beerInfo{}}
	for _, beer := range beerList {
		bi := &beerInfo{
			Id:             beer.Id,
			Name:           beer.Name,
			Manufacturer:   beer.Manufacturer,
			Price:          beer.Price,
			AlcoholContent: beer.AlcoholContent,
		}
		res.BeerList = append(res.BeerList, bi)
	}
	c.JSON(http.StatusOK, res)
}
