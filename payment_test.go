package paysimple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var paymentExample = `{
    "CustomerId": 255939,
    "CustomerFirstName": "Test",
    "CustomerLastName": "Person",
    "CustomerCompany": "ABC Company",
    "ReferenceId": 0,
    "Status": "Posted",
    "RecurringScheduleId": 0,
    "PaymentType": "ACH",
    "PaymentSubType": "Tel",
    "ProviderAuthCode": "Approved",
    "TraceNumber": "93e36116-6f29-4a2c-9ab0-b824fcb4f735",
    "PaymentDate": "2013-10-15T06:00:00Z",
    "ReturnDate": null,
    "EstimatedSettleDate": "2013-10-18T06:00:00Z",
    "ActualSettledDate": null,
    "CanVoidUntil": "2013-10-15T23:00:00Z",
    "FailureData": {
        "Code": null,
        "Description": null,
        "MerchantActionText": null,
        "IsDecline": false
    },
    "AccountId": 395564,
    "InvoiceId": null,
    "Amount": 9.44,
    "IsDebit": false,
    "InvoiceNumber": null,
    "PurchaseOrderNumber": null,
    "OrderId": null,
    "Description": null,
    "Latitude": null,
    "Longitude": null,
    "SuccessReceiptOptions": null,
    "FailureReceiptOptions": null,
    "Id": 1610692,
    "LastModified": "2013-10-15T20:37:29Z",
    "CreatedOn": "2013-10-15T20:37:29Z"
}`

func TestPayment(t *testing.T) {
	assert := assert.New(t)

	var payment Payment
	require.Nil(t, json.Unmarshal([]byte(paymentExample), &payment))

	assert.Equal(1610692, payment.ID)
	assert.Equal(255939, payment.CustomerID)
	assert.Equal("Test", payment.CustomerFirstName)
	assert.Equal("Person", payment.CustomerLastName)

	// TODO Test dates
}
