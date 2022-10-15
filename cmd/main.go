package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"crudsales/entity"
	"crudsales/repository"

	"github.com/gorilla/mux"
)

type JsonResponse struct {
    Type    string `json:"type"`
    Data[]  entity.Customer `json:"data"`
    Message string `json:"message"`
}

var db *sql.DB

func main() {
	db = repository.SetupDB()
	router := mux.NewRouter()
    router.HandleFunc("/customers/", GetCustomers).Methods("GET")
    router.HandleFunc("/customer/{id}", GetCustomer).Methods("GET")
    router.HandleFunc("/addCustomer/", AddCustomer).Methods("POST")
    router.HandleFunc("/deleteCustomer/{id}", DeleteCustomer).Methods("DELETE")
	fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8000", router))

}

// response and request handlers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers := repository.GetCustomers(db)
	var response = JsonResponse{Type: "success", Data: *customers}
    json.NewEncoder(w).Encode(response)
}

// response and request handlers
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")	
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := repository.GetCustomer(db, id)
	if err != nil {
        json.NewEncoder(w).Encode(err.Error())
    }else{
    	json.NewEncoder(w).Encode(customer)
	}
}

// response and request handlers
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer entity.Customer
	json.Unmarshal(reqBody, &customer)
	key, err := repository.AddCustomer(db, customer)
	if err != nil{
		json.NewEncoder(w).Encode(err.Error())
	}else{
		json.NewEncoder(w).Encode(key)
	}
	
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")	
	vars := mux.Vars(r)
	id := vars["id"]
	response, err := repository.DeleteCustomer(db, id)
	if err != nil {
        json.NewEncoder(w).Encode(err.Error())
    }else{
    	json.NewEncoder(w).Encode(response)
	}
}