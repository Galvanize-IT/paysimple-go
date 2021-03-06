package paysimple

import (
	"encoding/json"
	"net/url"
	"time"
)

type Status int64

const (
	Pending Status = iota
	Posted
	Settled
	Failed
	_
	Voided
	Reversed
	_
	_
	ReversePosted
	Chargeback
	_
	Authorized
	_
	Returned
	_
	ReverseNSF
	_
	RefundSettled
)

type Failure struct {
	Code               string
	Description        string
	MerchantActionText string
	IsDecline          bool
}

type Payment struct {
	ID                  int64 `json:"Id"`
	CustomerID          int64 `json:"CustomerId"`
	CustomerFirstName   string
	CustomerLastName    string
	CustomerCompany     string
	ReferenceID         int64 `json:"ReferenceId"`
	Status              string
	RecurringScheduleID int64 `json:"RecurringScheduleId"`
	PaymentType         string
	PaymentSubType      string
	ProviderAuthCode    string
	TraceNumber         string
	PaymentDate         time.Time
	ReturnDate          *time.Time
	EstimatedSettleDate time.Time
	ActualSettledDate   *time.Time
	CanVoidUntil        time.Time
	Amount              float64
	FailureData         Failure
	AccountID           int64 `json:"AccountId"`
	InvoiceID           int64 `json:"InvoiceId"`
	IsDebit             bool
	InvoiceNumber       string // TODO or int64?
	PurchaseOrderNumber string // TODO or int64?
	OrderID             int64  `json:"OrderId"`
	Description         string
	Latitude            *float64
	Longitude           *float64
	LastModified        time.Time
	CreateOn            time.Time
	// SuccessReceiptOptions
	// FailureReceiptOptions
}

// PaymentsResponse is the response returned from the payments endpoint
type PaymentsResponse struct {
	Meta     Meta
	Response []Payment
}

// PaymentFilters create GET parameters for the payments endpoint
type PaymentFilters struct {
	PaginationFilters
	StartDate, EndDate Date
	Lite               bool
}

// Values returns GET parameters for payments
func (f PaymentFilters) Values() url.Values {
	values := f.PaginationFilters.Values()
	if !f.StartDate.IsZero() {
		values.Set("startdate", f.StartDate.String())
	}
	if !f.EndDate.IsZero() {
		values.Set("enddate", f.EndDate.String())
	}
	if f.Lite {
		values.Set("lite", "true")
	}
	return values
}

// Payments handles requests to the Payment API endpoint
type Payments struct {
	api *api
}

// List returns multiple payments
// TODO pass filters
func (c *Payments) List() ([]Payment, error) {
	// Create a new request
	req, err := c.api.Get(c.api.URL("v4", "payment"))
	if err != nil {
		return nil, err
	}

	// Perform the request using the given backend
	resp, err := c.api.backend.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		// Convert the response body to the error format
		return nil, c.api.decodeError(resp)
	}

	var response PaymentsResponse
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if response.Meta.Errors != nil {
		return nil, response.Meta.Errors
	}

	// Return the actual payments
	return response.Response, nil
}
