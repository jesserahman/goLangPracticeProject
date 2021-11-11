package domain

type Customer struct {
	Id string `json:"Id"`
	Name string `json:"name"`
	City string `json:"city"`
	Zip  int    `json:"zip"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

