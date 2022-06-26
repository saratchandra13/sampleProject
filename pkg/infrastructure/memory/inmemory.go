package memory

import (
	_ "context"
	"github.com/ShareChat/service-template/pkg/domain/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	errNoDataForBeerId = "no data present for beer id"
)

type BeerSchema struct {
	id            string
	name          string
	company       string
	price         float64
	alcoholAmount float64
}

type inMemory struct {
	store map[string]BeerSchema
}

func NewMemoryStore() entity.BeerRepo {
	return &inMemory{store: make(map[string]BeerSchema)}
}

func generateBeerId() uuid.UUID {
	return uuid.New()
}

func parseBeerMeta(beerMeta BeerSchema) *entity.Beer {
	var parsedBeerMeta = &entity.Beer{
		Id:             beerMeta.id,
		Name:           beerMeta.name,
		Manufacturer:   beerMeta.company,
		Price:          beerMeta.price,
		AlcoholContent: beerMeta.alcoholAmount,
	}
	return parsedBeerMeta
}

func (mem *inMemory) AddBeer(beer *entity.Beer) (string, error) {
	beerId := generateBeerId().String()

	var beerMeta = BeerSchema{
		id:            beerId,
		name:          beer.Name,
		company:       beer.Manufacturer,
		price:         beer.Price,
		alcoholAmount: beer.AlcoholContent,
	}

	mem.store[beerId] = beerMeta
	return beerId, nil
}

func (mem *inMemory) GetBeer(id string) (*entity.Beer, error) {
	beerMeta, ok := mem.store[id]
	if !ok {
		return nil, errors.New(errNoDataForBeerId)
	}

	var parsedBeerMeta = parseBeerMeta(beerMeta)
	return parsedBeerMeta, nil
}

func (mem *inMemory) GetAllBeer() ([]*entity.Beer, error) {
	var beerList = make([]*entity.Beer, 0)
	for _, value := range mem.store {
		parsedBeerMeta := parseBeerMeta(value)
		beerList = append(beerList, parsedBeerMeta)
	}
	return beerList, nil
}
