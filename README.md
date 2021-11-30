# goLangPracticeProject

GoLangPracticeProject is a sample application based on this [udemy course](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/)
This project has been modified from the original project make use of additional endpoints


## Usage
There are currently 7 working endpoints in this project 

- GET `/customers` returns a list of all the customers in the customers db
- GET `/accounts` returns a list of all the accounts in the accounts db
- POST `/customer` creates a new customer
- GET `/customer/#{customer_id}` returns all the info for a specific customer (name, city, zip, status)
- PATCH `/customer/#{customer_id}` updates the info for a specific customer (name, city, zip, status)
- POST `/customer/#{customer?id}/account` creates an account for a specific customer
- GET `/customer/#{customer?id}/accounts` returns all accounts for a specific customer
- DELETE `/customer/#{customer?id}/account/#{account_id}` deletes a customer's account and all transactions associated with that account
- POST `/customer/#{customer?id}/account/#{account_id}/transaction` creates a new transaction for a specific customer's account
- GET `/customer/#{customer?id}/account/#{account_id}/transactions` returns all transactions for a specific customer's account

Sample JSON examples

``` 
POST /customer
{
    "name" : "John Doe",
    "date_of_birth" : "1990-01-30",
    "city" : "Los Angeles, California",
    "zip_code" : 90210,
    "status" : 1
}
```


``` 
PATCH /customer/{customer_id}
{
    "name" : "Updated John Doe",
    "date_of_birth" : "1990-01-30",
    "city" : "Los Angeles, California",
    "zip_code" : 90210,
    "status" : 1
}
```

``` 
POST /customer/{customer_id}/account
{
    "account_type" : "checking",
    "amount" : 5000
}
```

``` 
POST /customer/{customer_id}/account/{account_id}/transactions
{
    "amount" : 100.00,
    "transaction_type" : "deposit"
}
```