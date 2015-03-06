package paysimple

import (
	"net/url"
	"testing"

	"github.com/codegangsta/envy/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaysimple_URL(t *testing.T) {
	assert := assert.New(t)

	// Create a new API
	api := &api{baseURL: url.URL{Scheme: "https", Host: "api.paysimple.com"}}

	// Test customers
	assert.Equal(
		"https://api.paysimple.com/v4/customers",
		api.URL("v4", "customers").String(),
	)

	// Test Payments
	assert.Equal(
		"https://api.paysimple.com/v4/payments",
		api.URL("v4", "payments").String(),
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
		FirstName:             "Test A.",
		LastName:              "Customer",
		ShippingSameAsBilling: true,
	}
	_, err = api.Customers.Create(customer)
	require.Nil(t, err, "Customers Create should not error")
}
