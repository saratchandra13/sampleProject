package entity

type Review struct {
	Id     string
	BeerId string
	UserId string
	Meta   string
	Rating int8
}

func NewReview() *Review {
	return &Review{
		BeerId: "",
		UserId: "",
		Id:     "",
		Meta:   "",
		Rating: 0,
	}
}

type ReviewRepo interface {
	AddReview(*Review) (string, error)
	ListReview(string) ([]*Review, error)
}
