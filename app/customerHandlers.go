package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jesserahman/goLangPracticeProject/domain"

	"github.com/jesserahman/goLangPracticeProject/dto"

	"github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler *CustomerHandler) handleCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	status = strings.ToLower(status)

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}

	if status == "0" || status == "1" {
		customers, err := handler.service.GetCustomersByStatus(status)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}
	} else {
		customers, err := handler.service.GetAllCustomers()
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}
	}
}

func (handler *CustomerHandler) handleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := handler.service.GetCustomerById(customerId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func (handler *CustomerHandler) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var request dto.NewCustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		customer, customerErr := handler.service.CreateNewCustomer(request)
		if customerErr != nil {
			writeResponse(w, customerErr.Code, customerErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customer)
		}
	}
}

func (handler *CustomerHandler) handleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var customer domain.Customer
	customer.Id = customerId

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		customer, customerErr := handler.service.UpdateCustomer(customer)
		if customerErr != nil {
			writeResponse(w, customerErr.Code, customerErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customer)
		}
	}
}

func (handler *CustomerHandler) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	err := handler.service.DeleteCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, nil)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

type TimeStruct struct {
	CurrentTime time.Time `json:"current_time"`
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	writeResponse(w, http.StatusOK, "hello")
	currentTime := TimeStruct{CurrentTime: time.Now()}
	_ = json.NewEncoder(w).Encode(currentTime)
}
