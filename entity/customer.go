package entity

import (
	"database/sql"
)


type Customer struct {
	CustomerNumber         int             `json:"CustomerNumber"`
	CustomerName           string          `json:"CustomerName"`
	ContactLastName        string          `json:"ContactLastName"`
	ContactFirstName       string          `json:"ContactFirstName"`
	Phone                  string          `json:"Phone"`
	AddressLine1           string          `json:"AddressLine1"`
	AddressLine2           sql.NullString  `json:"AddressLine2"`
	City                   string          `json:"City"`
	State                  sql.NullString  `json:"State"`
	PostalCode             sql.NullString  `json:"PostalCode"`
	Country                string          `json:"Country"`
	SalesRepEmployeeNumber sql.NullInt64   `json:"SalesRepEmployeeNumber"`
	CreditLimit            sql.NullFloat64 `json:"CreditLimit"`
}