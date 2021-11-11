package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
	Zip  int    `json:"zip"`
}

func getAllCustomers() []Customer {
	customers := []Customer{{
		Name: "Jesse",
		City: "LIC",
		Zip:  11101,
	}, {
		Name: "Test",
		City: "Brooklyn",
		Zip:  11468,
	},
	}
	return customers
}


func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllCustomers())
}
