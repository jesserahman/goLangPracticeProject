package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	mux2 "github.com/gorilla/mux"
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Env variables not defined...")
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sanityCheck()

	// create instance of the handler
	handler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDbConnection())}

	router := mux2.NewRouter()
	router.HandleFunc("/greet", handleGreet).Methods(http.MethodGet)
	router.HandleFunc("/customers", handler.handleCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", handler.handleCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", handleCreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/time", handleTime).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
