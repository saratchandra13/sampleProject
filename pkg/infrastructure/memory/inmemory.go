package memory

import (
	_ "context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/saratchandra13/sampleProject/pkg/domain/entity"
)

var (
	errNoDataForVegetableId = "no data present for vegetable id"
)

type VegetableSchema struct {
	id            string
	name          string
	company       string
	price         float64
	calorieAmount float64
}

type inMemory struct {
	store map[string]VegetableSchema
}

func NewMemoryStore() entity.VegetableRepo {
	return &inMemory{store: make(map[string]VegetableSchema)}
}

func generateVegetableId() uuid.UUID {
	return uuid.New()
}

func parseVegetableMeta(VegetableMeta VegetableSchema) *entity.Vegetable {
	var parsedVegetableMeta = &entity.Vegetable{
		Id:           VegetableMeta.id,
		Name:         VegetableMeta.name,
		Seller:       VegetableMeta.company,
		Price:        VegetableMeta.price,
		CalorieCount: VegetableMeta.calorieAmount,
	}
	return parsedVegetableMeta
}

func (mem *inMemory) AddVegetable(vegetable *entity.Vegetable) (string, error) {
	VegetableId := generateVegetableId().String()

	var VegetableMeta = VegetableSchema{
		id:            VegetableId,
		name:          vegetable.Name,
		company:       vegetable.Seller,
		price:         vegetable.Price,
		calorieAmount: vegetable.CalorieCount,
	}

	mem.store[VegetableId] = VegetableMeta
	return VegetableId, nil
}

func (mem *inMemory) GetVegetable(id string) (*entity.Vegetable, error) {
	VegetableMeta, ok := mem.store[id]
	if !ok {
		return nil, errors.New(errNoDataForVegetableId)
	}

	var parsedVegetableMeta = parseVegetableMeta(VegetableMeta)
	return parsedVegetableMeta, nil
}

func (mem *inMemory) GetAllVegetable() ([]*entity.Vegetable, error) {
	var VegetableList = make([]*entity.Vegetable, 0)
	for _, value := range mem.store {
		parsedVegetableMeta := parseVegetableMeta(value)
		VegetableList = append(VegetableList, parsedVegetableMeta)
	}
	return VegetableList, nil
}
