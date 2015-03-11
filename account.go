package paysimple

import (
	"encoding/json"
	"net/http"
	"time"
)

type Issuer string

const (
	Visa       Issuer = "12"
	MasterCard        = "13"
	Amex              = "14"
	Discover          = "15"
)

type CreditCard struct {
	CreditCardNumber string
	ExpirationDate   string
	Issuer           Issuer
	BillingZipCode   string
	CustomerID       int64 `json:"CustomerId"`
	IsDefault        bool
	ID               int64 `json:"Id"`
	LastModified     time.Time
	CreatedOn        time.Time
}

// CreditCardResponse is returned by Accounts.CreateCreditCard
type CreditCardResponse struct {
	Meta     Meta
	Response CreditCard
}

// Accounts handles requests to the Account API endpoint
type Accounts struct {
	api *api
}

// CreateCreditCard creates a new credit card account
func (c *Accounts) CreateCreditCard(create CreditCard) (CreditCard, error) {
	// Blank card that will be returned on error
	var created CreditCard

	// Create a new request
	uri := c.api.URL("v4", "account", "creditcard")
	req, err := c.api.Post(uri, create)
	if err != nil {
		return created, err
	}

	// Perform the request using the given backend
	resp, err := c.api.backend.Do(req)
	if err != nil {
		return created, err
	}

	if resp.StatusCode != http.StatusCreated {
		// Convert the response body to the error format
		return created, c.api.decodeError(resp)
	}

	var response CreditCardResponse

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Response, err
}
