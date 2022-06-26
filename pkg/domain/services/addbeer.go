package services

import (
	"errors"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"regexp"
)

var (
	errInvalidName = errors.New("validationError: invalid input name")
)

type BeerInfo struct {
	Name  string  `json:"name"`
	Manuf string  `json:"manuf"`
	Price float64 `json:"price"`
}

func (b *BeerInfo) validate() error {
	result, err := regexp.Match("[a-z]+", []byte(b.Name))
	if err != nil || !result {
		return errInvalidName
	}
	return nil
}

func (al *appLogic) AddBeer(b *BeerInfo) (string, error) {
	err := b.validate()
	if err != nil {
		return "", err
	}

	// assuming we get alcohol content from third party service
	var alcoholAmt = 5.5
	beerMeta := &entity.Beer{
		Name:           b.Name,
		Manufacturer:   b.Manuf,
		Price:          b.Price,
		AlcoholContent: alcoholAmt,
	}

	beerId, err := al.beerRepo.AddBeer(beerMeta)
	if err != nil {
		return "", err
	}

	return beerId, nil
}
