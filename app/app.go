package app

import (
	"log"
	"net/http"
)

func Run(){
	http.HandleFunc("/greet", handleGreet)
	http.HandleFunc("/customers", handleCustomers)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}