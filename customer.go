package paysimple

import (
	"time"
)

type Address struct {
	StreetAddress1 string
	StreetAddress2 string
	City           string
	StateCode      string
	ZipCode        string
	Country        string
}

type Customer struct {
	ID                    int64 `json:"Id"`
	FirstName             string
	MiddleName            string
	LastName              string
	BillingAddress        *Address
	ShippingSameAsBoolean bool
	ShippingAddress       *Address
	Company               string
	CustomerAccount       string
	Phone                 string
	AltPhone              string
	MobilePhone           string
	Fax                   string
	Email                 string
	AltEmail              string
	Website               string
	Notes                 string
	LastModified          time.Time
	CreatedOn             time.Time
}

type CustomersResponse struct {
	Meta     Meta
	Response []Customer
}
