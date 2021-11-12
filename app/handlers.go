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
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (handler *CustomerHandler) handleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := handler.service.GetCustomer(customerId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
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
