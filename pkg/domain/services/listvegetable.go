package services

import (
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
)

type VegetableList []*entity.Vegetable

func (al *appLogic) ListVegetable() (VegetableList, error) {
	VegetableList, err := al.VegetableRepo.GetAllVegetable()
	if err != nil {
		return nil, err
	}
	return VegetableList, nil
}
