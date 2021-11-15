package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"

	"github.com/jesserahman/goLangPracticeProject/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	dbClient *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	customersQuery := fmt.Sprintf("select customer_id, name, city, zipcode, status from customers where customer_id = '%s'", id)
	err := d.dbClient.Get(&c, customersQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error searching DB for customer_id " + err.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error searching DB for customer_id " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) ByStatus(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	customerStatus, _ := strconv.Atoi(status)
	customersQuery := fmt.Sprintf("select customer_id, name, city, zipcode, status from customers where status = %d", customerStatus)

	// query the DB, and store the result in ${customers}
	err := d.dbClient.Select(&customers, customersQuery)
	if err != nil {
		logger.Error("Error searching by status in the customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	customersQuery := "select customer_id, name, city, zipcode, status from customers"
	// query the DB, and store the result in ${customers}
	err := d.dbClient.Select(&customers, customersQuery)
	if err != nil {
		logger.Error("Error querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func NewCustomerRepositoryDbConnection(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
