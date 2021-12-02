package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zip         int    `json:"zip_code"`
	Status      string `json:"status"`
	DateOfBirth string `json:"date_of_birth"`
}
