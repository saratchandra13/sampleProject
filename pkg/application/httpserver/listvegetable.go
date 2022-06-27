package httpserver

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
