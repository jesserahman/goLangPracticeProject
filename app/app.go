package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Adding auth middleware
	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	fmt.Println("WITH auth")
	port := os.Getenv("SERVER_PORT")

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
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

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPassword, dbAddress, dbPort, dbName)
	fmt.Println("Datasource: ", datasource)
	dbClient, err := sqlx.Open("mysql", datasource)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatalf("could not connect to the MySQL database... %v", err)
	}

	// Pinging DB until it's ready
	fmt.Println("Starting to Ping DB")
	for {

		err = db.Ping()
		if err == nil {
			fmt.Println("DB ready for Migrations")
			break
		}

		fmt.Println("DB not ready.. Pinging again!")
		time.Sleep(2 * time.Second)
		continue

	}

	// Run migrations
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://db/migrations"), // file://path/to/directory
		"mysql", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
	return dbClient
}
