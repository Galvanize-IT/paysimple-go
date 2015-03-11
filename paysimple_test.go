package paysimple

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/codegangsta/envy/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaysimple_test(t *testing.T) {
	// Bootstrap the environment
	envy.Bootstrap()

	// Start an http test server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Attempt to unmarshal the request body
		var customer Customer
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			panic(err)
		}

		// And set it back down with an ID set
		customer.ID = 99
		content, err := json.Marshal(customer)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(content)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	api := test(ts.URL)

	// Create a new customer
	customer := Customer{
		FirstName:             "Test B.",
		LastName:              "Customer",
		ShippingSameAsBilling: true,
	}
	created, err := api.Customers.Create(customer)
	require.Nil(t, err, "Customers Create should not error")
	assert.Equal(t, 99, created.ID)
	assert.Equal(t, customer.FirstName, created.FirstName)
	assert.Equal(t, customer.LastName, created.LastName)
	assert.Equal(t,
		customer.ShippingSameAsBilling,
		created.ShippingSameAsBilling,
	)
}

func TestPaysimple_URL(t *testing.T) {
	assert := assert.New(t)

	// Create a new API
	api := &api{baseURL: url.URL{Scheme: "https", Host: "api.paysimple.com"}}

	// Test customers
	assert.Equal(
		"https://api.paysimple.com/v4/customer",
		api.URL("v4", "customer").String(),
	)

	// Test Payments
	assert.Equal(
		"https://api.paysimple.com/v4/payment",
		api.URL("v4", "payment").String(),
	)
}

func TestPaysimple_Payments(t *testing.T) {
	// Bootstrap the environment
	envy.Bootstrap()

	// Create a sanbox API
	api := Sandbox()

	_, err := api.Payments.List()
	require.Nil(t, err, "Payments List should not error")
}

func TestPaysimple_Customers(t *testing.T) {
	// Bootstrap the environment
	envy.Bootstrap()

	// Create a sanbox API
	api := Sandbox()

	var err error

	_, err = api.Customers.List()
	require.Nil(t, err, "Customers List should not error")

	// Create a new customer
	customer := Customer{
		FirstName:             "Test B.",
		LastName:              "Customer",
		ShippingSameAsBilling: true,
	}
	_, err = api.Customers.Create(customer)
	require.Nil(t, err, "Customers Create should not error")
}
