package entity

type Beer struct {
	Id             string
	Name           string
	Manufacturer   string
	Price          float64
	AlcoholContent float64
}

func NewBeer() *Beer {
	return &Beer{}
}

type BeerRepo interface {
	AddBeer(beer *Beer) (string, error)
	GetBeer(id string) (*Beer, error)
	GetAllBeer() ([]*Beer, error)
}
