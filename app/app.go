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
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Env variables not defined...")
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sanityCheck()

	dbClient := getDbClient()

	// create instance of the handler
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDbConnection(dbClient))}
	accountHandler := AccountHandler{service.NewAccountService(domain.NewAccountRepositoryDbConnection(dbClient))}
	transactionHandler := TransactionHandler{service.NewTransactionService(domain.NewTransactionRepositoryDbConnection(dbClient))}

	router := mux2.NewRouter()
	router.HandleFunc("/customers", customerHandler.handleCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.HandleFunc("/accounts", accountHandler.handleAccounts).
		Methods(http.MethodGet).
		Name("GetAllAccounts")
	router.HandleFunc("/customer", customerHandler.handleCreateCustomer).
		Methods(http.MethodPost).
		Name("CreateCustomer")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleCustomer).
		Methods(http.MethodGet).
		Name("GetCustomerById")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleUpdateCustomer).
		Methods(http.MethodPatch).
		Name("UpdateCustomerByID")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.handleDeleteCustomer).
		Methods(http.MethodDelete).
		Name("DeleteCustomerById")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", accountHandler.handleCreateAccount).
		Methods(http.MethodPost).
		Name("CreateAccount")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/accounts", accountHandler.handleGetAccountsByCustomerId).
		Methods(http.MethodGet).
		Name("GetAccountsByCustomerId")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleGetAccountById).
		Methods(http.MethodGet).
		Name("GetAccountsByAccountId")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleUpdateAccount).
		Methods(http.MethodPatch).
		Name("UpdateAccountByAccountId")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.handleDeleteAccount).
		Methods(http.MethodDelete).
		Name("DeleteAccountById")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", transactionHandler.handleCreateNewTransaction).
		Methods(http.MethodPost).
		Name("CreateTransaction")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transactions", transactionHandler.handleGetAllTransactionsByAccountId).
		Methods(http.MethodGet).
		Name("GetAllTransactionsByAccountId")

	// *********  TEMPORARILY COMMENTING OUT TO TEST WITH DOCKER *********
	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	fmt.Println("WITH auth")
	port := os.Getenv("SERVER_PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DOCKER_DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	fmt.Println("Datasource: ", datasource)
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
