package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jesserahman/goLangPracticeProject/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	dbClient *sql.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customersQuery := fmt.Sprintf("select customer_id, name, city, zipcode, status from customers where customer_id = '%s'", id)
	row := d.dbClient.QueryRow(customersQuery)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) ByStatus(status string) ([]Customer, *errs.AppError) {
	fmt.Println("Status passed in ", status)
	customerStatus := 0
	if status == "active" {
		customerStatus = 1
	}

	fmt.Println("Customer status value: ", customerStatus)
	customersQuery := fmt.Sprintf("select customer_id, name, city, zipcode, status from customers where status = %d", customerStatus)
	rows, err := d.dbClient.Query(customersQuery)
	if err != nil {
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.Status)
		if err != nil {
			log.Println("Error scanning Customers " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	customersQuery := "select customer_id, name, city, zipcode, status from customers"
	rows, err := d.dbClient.Query(customersQuery)
	if err != nil {
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.Status)
		if err != nil {
			log.Println("Error scanning Customers " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDbConnection() CustomerRepositoryDb {
	dbClient, err := sql.Open("mysql", "root:jesse jesse@tcp(localhost:3306)/banking")
	if err != nil {
		fmt.Errorf("Unable to connect to DB")
		panic(err)
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return CustomerRepositoryDb{dbClient}
}
