package services

import (
	"errors"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
	"regexp"
)

var (
	errInvalidName = errors.New("validationError: invalid input name")
)

type VegetableInfo struct {
	Name   string  `json:"name"`
	Seller string  `json:"seller"`
	Price  float64 `json:"price"`
}

func (b *VegetableInfo) validate() error {
	result, err := regexp.Match("[a-z]+", []byte(b.Name))
	if err != nil || !result {
		return errInvalidName
	}
	return nil
}

func (al *appLogic) AddVegetable(b *VegetableInfo) (string, error) {
	err := b.validate()
	if err != nil {
		return "", err
	}

	// assuming we get calorie content from third party service
	var calorieAmt = 5.5
	VegetableMeta := &entity.Vegetable{
		Name:         b.Name,
		Seller:       b.Seller,
		Price:        b.Price,
		CalorieCount: calorieAmt,
	}

	VegetableId, err := al.VegetableRepo.AddVegetable(VegetableMeta)
	if err != nil {
		return "", err
	}

	return VegetableId, nil
}
