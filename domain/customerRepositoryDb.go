package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	dbClient *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	customersQuery := "select customer_id, name, city, zipcode from customers"
	rows, err := d.dbClient.Query(customersQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println("ROWS FOUND")
	fmt.Println(&rows)
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zip)
		if err != nil {
			log.Println("Error scanning Customers " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
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
