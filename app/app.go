package app

import (
	"log"
	"net/http"

	mux2 "github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/service"
)

func Run() {
	// create instance of the handler

	handler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDbConnection())}

	router := mux2.NewRouter()
	router.HandleFunc("/greet", handleGreet).Methods(http.MethodGet)
	router.HandleFunc("/customers", handler.handleCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id}", handler.handleCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", handleCreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/time", handleTime).Methods(http.MethodGet)
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
