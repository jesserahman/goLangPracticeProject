package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"

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
	log.Println("Address: " , os.Getenv("SERVER_ADDRESS"))
	log.Println("Port: " , os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sanityCheck()
	log.Println("sanity check passed")

	dbClient := getDbClient()

	log.Println("got db client")
	log.Println("test")

	// create instance of the handler
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDbConnection(dbClient))}
	accountHandler := AccountHandler{service.NewAccountService(domain.NewAccountRepositoryDbConnection(dbClient))}
	transactionHandler := TransactionHandler{service.NewTransactionService(domain.NewTransactionRepositoryDbConnection(dbClient))}

	log.Println("created handlers")
	router := mux2.NewRouter()
	router.HandleFunc("/customers", customerHandler.handleCustomers).Methods(http.MethodGet)
	router.HandleFunc("/accounts", accountHandler.handleAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customer", customerHandler.handleCreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleUpdateCustomer).Methods(http.MethodPatch)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleDeleteCustomer).Methods(http.MethodDelete)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", accountHandler.handleCreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/accounts", accountHandler.handleGetAccountsByCustomerId).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleGetAccountById).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleUpdateAccount).Methods(http.MethodPatch)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleDeleteAccount).Methods(http.MethodDelete)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", transactionHandler.handleCreateNewTransaction).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transactions", transactionHandler.handleGetAllTransactionsByAccountId).Methods(http.MethodGet)
	router.HandleFunc("/api/time", handleTime).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	fmt.Println("DataSource: ", datasource)
	dbClient, err := sqlx.Open("mysql", datasource)
	if err != nil {
		logger.Error("Error connecting to the DB " + err.Error())
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
