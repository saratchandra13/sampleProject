package entity

type Vegetable struct {
	Id           string
	Name         string
	Seller       string
	Price        float64
	CalorieCount float64
}

func NewVegetable() *Vegetable {
	return &Vegetable{}
}

type VegetableRepo interface {
	AddVegetable(vegetable *Vegetable) (string, error)
	GetVegetable(id string) (*Vegetable, error)
	GetAllVegetable() ([]*Vegetable, error)
}
