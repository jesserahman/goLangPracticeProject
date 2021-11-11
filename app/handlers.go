package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/service"
	"net/http"
	"time"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler *CustomerHandler)handleCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		fmt.Errorf("Error retrieving customers")
	}
	json.NewEncoder(w).Encode(customers)
}


type TimeStruct struct {
	CurrentTime time.Time `json:"current_time"`
}

func getCustomer(id string, customers []domain.Customer) *domain.Customer {
	for _, customer := range customers {
		if customer.Id == id {
			fmt.Println("Customer found", customer)
			return &customer
		}
	}

	return nil
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	currentTime := TimeStruct{CurrentTime: time.Now()}
	json.NewEncoder(w).Encode(currentTime)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func (handler *CustomerHandler)handleCustomer(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	w.Header().Add("Content-Type", "application/json")

	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		fmt.Errorf("Error retrieving customers")
	}

	json.NewEncoder(w).Encode(getCustomer(customerId, customers))
}

func handleCreateCustomer(w http.ResponseWriter, r *http.Request){

}
