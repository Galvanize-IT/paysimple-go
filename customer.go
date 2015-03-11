package paysimple

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	ID                    int64 `json:"Id,omitempty"`
	FirstName             string
	MiddleName            string `json:",omitempty"`
	LastName              string
	BillingAddress        *Address `json:",omitempty"`
	ShippingSameAsBilling bool
	ShippingAddress       *Address   `json:",omitempty"`
	Company               string     `json:",omitempty"`
	CustomerAccount       string     `json:",omitempty"`
	Phone                 string     `json:",omitempty"`
	AltPhone              string     `json:",omitempty"`
	MobilePhone           string     `json:",omitempty"`
	Fax                   string     `json:",omitempty"`
	Email                 string     `json:",omitempty"`
	AltEmail              string     `json:",omitempty"`
	Website               string     `json:",omitempty"`
	Notes                 string     `json:",omitempty"`
	LastModified          *time.Time `json:",omitempty"`
	CreatedOn             *time.Time `json:",omitempty"`
}

type CustomersResponse struct {
	Meta     Meta
	Response []Customer
}

// Customers handles requests to the Customers API endpoint
type Customers struct {
	api *api
}

// Create creates and returns a single customer
func (c *Customers) Create(create Customer) (created Customer, err error) {
	// Create a new request
	var req *http.Request
	if req, err = c.api.Post(c.api.URL("v4", "customer"), create); err != nil {
		return
	}

	// Perform the request using the given backend
	var resp *http.Response
	if resp, err = c.api.backend.Do(req); err != nil {
		return
	}

	if resp.StatusCode != http.StatusCreated {
		// Convert the response body to the error format
		err = c.api.decodeError(resp)
		// TODO error might still be nil!
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&created)
	return
}

// Get returns a single customer
func (c *Customers) Get(id int64) (customer Customer, err error) {
	// Create a new request
	uri := c.api.URL("v4", "customer", fmt.Sprintf("%d", id))
	var req *http.Request
	if req, err = c.api.Get(uri); err != nil {
		return
	}

	// Perform the request using the given backend
	var resp *http.Response
	if resp, err = c.api.backend.Do(req); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		// Convert the response body to the error format
		err = c.api.decodeError(resp)
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&customer)
	return
}

// List returns a list of customers
// TODO filters
func (c *Customers) List() ([]Customer, error) {
	// Create a new request
	req, err := c.api.Get(c.api.URL("v4", "customer"))
	if err != nil {
		return nil, err
	}

	// Perform the request using the given backend
	resp, err := c.api.backend.Do(req)
	if err != nil {
		return nil, err
	}

	var response CustomersResponse
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// Convert the response body to the error format
		return nil, c.api.decodeError(resp)
	}

	if response.Meta.Errors != nil {
		return nil, response.Meta.Errors
	}

	// Return the actual customers
	return response.Response, nil
}
