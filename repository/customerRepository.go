package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"

	"crudsales/entity"
)

// DB set up
func SetupDB()  *sql.DB {
	cfg :=  mysql.Config{
		User: "test",
		Passwd: "test",
		Net: "tcp",
		Addr: "localhost:3306",
		DBName: "classicmodels",
		AllowNativePasswords: true,
	}

    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        panic(err.Error())
    }
	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
    fmt.Println("Is connect to the DB!")
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(2*time.Minute)
	return db
}

func GetCustomers(db *sql.DB) *[]entity.Customer{
    rows, err := db.Query("SELECT * FROM customers")
    checkErr(err)
    var customers []entity.Customer
    for rows.Next() {
		var cus entity.Customer
		err = rows.Scan(&cus.CustomerNumber, &cus.CustomerName, &cus.ContactLastName, &cus.ContactFirstName, &cus.Phone, &cus.AddressLine1, &cus.AddressLine2,
			&cus.City, &cus.State, &cus.PostalCode, &cus.Country, &cus.SalesRepEmployeeNumber, &cus.CreditLimit)
        checkErr(err)
        customers = append(customers, cus)
    }
	fmt.Println("@OnboardingConnectionPool MYSQL MAX Open Connections: ",db.Stats().MaxOpenConnections)
	fmt.Println("@OnboardingConnectionPool MYSQL Open Connections: ", db.Stats().OpenConnections)
	fmt.Println("@OnboardingConnectionPool MYSQL InUse Connections: ", db.Stats().InUse)
	fmt.Println("@OnboardingConnectionPool MYSQL Idle Connections: ", db.Stats().Idle)
	return &customers
}

func GetCustomer(db *sql.DB, id string) (entity.Customer, error){
	var cus entity.Customer
    row := db.QueryRow("SELECT * FROM customers WHERE customerNumber= ?", id)
	if err := row.Scan(&cus.CustomerNumber, &cus.CustomerName, &cus.ContactLastName, &cus.ContactFirstName, &cus.Phone, &cus.AddressLine1, &cus.AddressLine2,
		&cus.City, &cus.State, &cus.PostalCode, &cus.Country, &cus.SalesRepEmployeeNumber, &cus.CreditLimit); err != nil {
        if err == sql.ErrNoRows {
            return cus, fmt.Errorf("GetCustomer %s: no such customer", id)
        }
        return cus, fmt.Errorf("GetCustomer %s: %v", id, err)
    }
	return cus, nil
}

func AddCustomer(db *sql.DB, cus entity.Customer) (int, error) {
    result, err := db.Exec("INSERT INTO `customers` (`customerNumber`, `customerName`, `contactLastName`, `contactFirstName`, `phone`, `addressLine1`, `addressLine2`, `city`, `state`, `postalCode`, `country`, `salesRepEmployeeNumber`, `creditLimit`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
	cus.CustomerNumber, cus.CustomerName, cus.ContactLastName,cus.ContactFirstName, cus.Phone, cus.AddressLine1, cus.AddressLine2, cus.City, cus.State, cus.PostalCode,
	cus.Country, cus.SalesRepEmployeeNumber, cus.CreditLimit)
    if err != nil {
        return 0, fmt.Errorf("AddCustomer: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddCustomer: %v,%d", err, id)
    }
    return cus.CustomerNumber, nil
}



func DeleteCustomer(db *sql.DB, id string) (string, error){
    _, err := db.Exec("DELETE FROM customers WHERE customerNumber= ?", id)
	if err != nil {
        return "No se pudo borrar el registro", err
    }
	return "Se borro el registro satisfactorimente", nil
}

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}