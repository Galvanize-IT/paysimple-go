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

// An actual customer response
var customerResponseExample = `{
  "Meta": {
    "Errors": null,
    "HttpStatus": "Created",
    "HttpStatusCode": 201,
    "PagingDetails": null
  },
  "Response": {
    "Website": null,
    "ShippingAddress": null,
    "Fax": null,
    "Id": 292152,
    "MobilePhone": null,
    "FirstName": "Test B.",
    "Phone": null,
    "MiddleName": null,
    "LastName": "Customer",
    "Company": null,
    "AltPhone": null,
    "CreatedOn": "2015-03-11T05:31:53.2209562Z",
    "Email": null,
    "ShippingSameAsBilling": false,
    "CustomerAccount": null,
    "BillingAddress": null,
    "LastModified": "2015-03-11T05:31:53.2209562Z",
    "AltEmail": null,
    "Notes": null
  }
}`

func TestCustomer(t *testing.T) {
	var customer Customer
	require.Nil(t, json.Unmarshal([]byte(customerExample), &customer))

	assert.Equal(t, 255939, customer.ID)
	assert.Equal(t, "Test", customer.FirstName)
	assert.Equal(t, "Person", customer.LastName)
	assert.Equal(t, "", customer.MiddleName)

	var response CustomerResponse
	require.Nil(t, json.Unmarshal([]byte(customerResponseExample), &response))
	customer = response.Response

	assert.Equal(t, 292152, customer.ID)
	assert.Equal(t, "Test B.", customer.FirstName)
	assert.Equal(t, "Customer", customer.LastName)
	assert.Equal(t, "", customer.MiddleName)

}
