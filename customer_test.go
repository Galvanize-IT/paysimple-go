package paysimple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var customerExample = `{
    "MiddleName": null,
    "AltEmail": null,
    "AltPhone": null,
    "MobilePhone": null,
    "Fax": null,
    "Website": null,
    "BillingAddress": {
        "StreetAddress1": "123 Some St",
        "StreetAddress2": null,
        "City": "Denver",
        "StateCode": "CO",
        "ZipCode": "80202",
        "Country": null
    },
    "ShippingSameAsBilling": true,
    "ShippingAddress": {
        "StreetAddress1": "123 Some St",
        "StreetAddress2": null,
        "City": "Denver",
        "StateCode": "CO",
        "ZipCode": "80202",
        "Country": null
    },
    "Company": "ABC Company",
    "Notes": "This is a note about ABC Company.",
    "CustomerAccount": "TP-117",
    "FirstName": "Test",
    "LastName": "Person",
    "Email": "testperson@abcco.com",
    "Phone": "8005551212",
    "Id": 255939,
    "LastModified": "2013-10-01T19:16:38.5061103Z",
    "CreatedOn": "2013-10-01T19:16:38.5061103Z"
}`

func TestCustomer(t *testing.T) {
	assert := assert.New(t)

	var customer Customer
	require.Nil(t, json.Unmarshal([]byte(customerExample), &customer))

	assert.Equal(255939, customer.ID)
	assert.Equal("Test", customer.FirstName)
	assert.Equal("Person", customer.LastName)
	assert.Equal("", customer.MiddleName)
}
