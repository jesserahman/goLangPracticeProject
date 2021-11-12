package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler *CustomerHandler) handleCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		_ = json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(customers)
	}

}

func (handler *CustomerHandler) handleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := handler.service.GetCustomer(customerId)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		_ = json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(customer)
	}
}

type TimeStruct struct {
	CurrentTime time.Time `json:"current_time"`
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	currentTime := TimeStruct{CurrentTime: time.Now()}
	_ = json.NewEncoder(w).Encode(currentTime)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

func handleCreateCustomer(w http.ResponseWriter, r *http.Request) {

}
