package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub)FindAll() ([]Customer, error){
	return c.customers, nil
}

func GenerateNewCustomers() CustomerRepositoryStub {
	customers := CustomerRepositoryStub{[]Customer{
		{
			Id:   "001",
			Name: "Jesse",
			City: "LIC",
			Zip:  11101,
		}, {
			Id:   "002",
			Name: "Test",
			City: "Brooklyn",
			Zip:  11468,
		},
	}}
	return customers
}
