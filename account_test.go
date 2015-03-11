package paysimple

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var creditCardExample = `{
   "CreditCardNumber": "************1111",
   "ExpirationDate": "12/2015",
   "Issuer": "Visa",
   "BillingZipCode": "80202",
   "CustomerId": 255939,
   "IsDefault": true,
   "Id": 395560,
   "LastModified": "2013-10-11T17:41:36.6661404Z",
   "CreatedOn": "2013-10-11T17:41:36.6661404Z"
}`

func TestAccount(t *testing.T) {
	var card CreditCard
	require.Nil(t, json.Unmarshal([]byte(creditCardExample), &card))

	assert.Equal(t, "************1111", card.CreditCardNumber)
	assert.Equal(t, "Visa", card.Issuer)
	assert.Equal(t, 255939, card.CustomerID)
	assert.Equal(t, true, card.IsDefault)
	assert.Equal(t,
		time.Date(2013, 10, 11, 17, 41, 36, 666140400, time.UTC),
		card.CreatedOn,
	)
}
